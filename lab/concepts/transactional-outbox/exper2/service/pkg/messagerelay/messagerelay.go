package messagerelay

import (
	"time"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgproto3"
)

// https://microservices.io/patterns/data/transaction-log-tailing.html
//
// https://debezium.io/documentation/reference/stable/connectors/postgresql.html#debezium-connector-for-postgresql
// https://www.postgresql.org/docs/current/logicaldecoding-walsender.html
// https://www.postgresql.org/docs/current/logical-replication-architecture.html
//
// statements used by debezium:
// CREATE_REPLICATION_SLOT "debezium"  LOGICAL decoderbufs
// CREATE_REPLICATION_SLOT "debezium"  LOGICAL pgoutput
/*
2024-02-19 15:08:46.610 UTC [207] STATEMENT:  CREATE_REPLICATION_SLOT "debezium"  LOGICAL pgoutput
2024-02-19 15:08:46.610 UTC [207] LOG:  exported logical decoding snapshot: "00000006-00000003-1" with 0 transaction IDs
2024-02-19 15:08:46.610 UTC [207] STATEMENT:  CREATE_REPLICATION_SLOT "debezium"  LOGICAL pgoutput
2024-02-19 15:08:47.547 UTC [207] LOG:  starting logical decoding for slot "debezium"
2024-02-19 15:08:47.547 UTC [207] DETAIL:  Streaming transactions committing after 0/1555570, reading WAL from 0/1555538.
2024-02-19 15:08:47.547 UTC [207] STATEMENT:  START_REPLICATION SLOT "debezium" LOGICAL 0/1555570 ("proto_version" '1', "publication_names" 'dbz_publication', "messages" 'true')
2024-02-19 15:08:47.547 UTC [207] LOG:  logical decoding found consistent point at 0/1555538
2024-02-19 15:08:47.547 UTC [207] DETAIL:  There are no running transactions.
2024-02-19 15:08:47.547 UTC [207] STATEMENT:  START_REPLICATION SLOT "debezium" LOGICAL 0/1555570 ("proto_version" '1', "publication_names" 'dbz_publication', "messages" 'true')
2024-02-19 15:09:22.065 UTC [226] LOG:  starting logical decoding for slot "debezium"
2024-02-19 15:09:22.065 UTC [226] DETAIL:  Streaming transactions committing after 0/1555570, reading WAL from 0/1555538.
2024-02-19 15:09:22.065 UTC [226] STATEMENT:  START_REPLICATION SLOT "debezium" LOGICAL 0/1555570 ("proto_version" '1', "publication_names" 'dbz_publication', "messages" 'true')
2024-02-19 15:09:22.066 UTC [226] LOG:  logical decoding found consistent point at 0/1555538
2024-02-19 15:09:22.066 UTC [226] DETAIL:  There are no running transactions.
2024-02-19 15:09:22.066 UTC [226] STATEMENT:  START_REPLICATION SLOT "debezium" LOGICAL 0/1555570 ("proto_version" '1', "publication_names" 'dbz_publication', "messages" 'true')
*/
func PostgresLogTailing(ctx context.Context, pCl *pgxpool.Pool) {
	// implemented using pg_output (logical replication connection)

	// ESTABLISH LOGICAL REPLICATION CONNECTION
	// https://pkg.go.dev/github.com/jackc/pglogrepl#section-readme

	// current exper highly derived from demo: (MIT LICENSE)
	// https://github.com/jackc/pglogrepl/blob/master/example/pglogrepl_demo/main.go

	//
	// acquire low level conn from pgxpool
	//
	/*
	poolConn, err := pCl.Acquire(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to acquire db conn from pool")
	}

	pgxConn := poolConn.Conn()
	conn := pgxConn.PgConn()
	// ^ todo: instead of all this - just open low level conn direct?
	*/
	conn, err := pgconn.Connect(ctx, "postgresql://postgres:postgres@test-service-postgres:5432/postgres?sslmode=disable&replication=database")
	// https://github.com/jackc/pglogrepl/issues/6
	// > "I was missing the ?replication=database option" results in error when calling IDENTIFY_SYSTEM
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db connection")
	}

	//
	// create publication?
	//
	result := conn.Exec(ctx, "DROP PUBLICATION IF EXISTS messagerelay_testsvc;")
	_, err = result.ReadAll()
	if err != nil {
		log.Fatal().Err(err).Msg("failed drop publication")
	}

	result = conn.Exec(ctx, "CREATE PUBLICATION messagerelay_testsvc FOR TABLE event_outbox;")
	_, err = result.ReadAll()
	if err != nil {
		log.Fatal().Err(err).Msg("failed create publication")
	}

	//
	// create logical repl w/ pgoutput plugin
	//
	sysident, err := pglogrepl.IdentifySystem(ctx, conn)
	if err != nil {
		log.Fatal().Err(err).Msg("IdentifySystem failed.")
	}

	// slot name is 'messagerelay_testsvc' as defined in publications above
	const slotName string = "messagerelay_testsvc"
	const outputPlugin string = "pgoutput"

	// temporary repl slot
	_, err = pglogrepl.CreateReplicationSlot(ctx, conn, slotName, outputPlugin, pglogrepl.CreateReplicationSlotOptions{Temporary: true})
	if err != nil {
		log.Fatal().Err(err).Msg("CreateReplicationSlot failed")
	}

	// logcl repl slot
	pluginArguments := []string{
		"proto_version '2'",
		"publication_names 'messagerelay_testsvc'",
		"messages 'true'",
		"streaming 'true'",
	}
	err = pglogrepl.StartReplication(ctx, conn, slotName, sysident.XLogPos, pglogrepl.StartReplicationOptions{PluginArgs: pluginArguments})
	if err != nil {
		log.Fatal().Err(err).Msg("StartReplication failed")
	}

	//
	// begin consuming logcl repl slot
	//

	// do not understand this stuff?
	clientXLogPos := sysident.XLogPos
	standbyMessageTimeout := time.Second * 10
	nextStandbyMessageDeadline := time.Now().Add(standbyMessageTimeout)
	relationsV2 := map[uint32]*pglogrepl.RelationMessageV2{}
	typeMap := pgtype.NewMap()

	// whenever we get StreamStartMessage we set inStream to true and then pass it to DecodeV2 function
	// on StreamStopMessage we set it back to false
	inStream := false

	for {
		// stop if context is cancelled
		select {
			case <-ctx.Done():
				return
			default:
				//
				// processing
				//
				if time.Now().After(nextStandbyMessageDeadline) {
					err = pglogrepl.SendStandbyStatusUpdate(ctx, conn, pglogrepl.StandbyStatusUpdate{WALWritePosition: clientXLogPos})
					if err != nil {
						log.Fatal().Err(err).Msg("SendStandbyStatusUpdate failed")
					}
					log.Printf("Sent Standby status message at %s\n", clientXLogPos.String())
					nextStandbyMessageDeadline = time.Now().Add(standbyMessageTimeout)
				}

				ctx2, cancel := context.WithDeadline(ctx, nextStandbyMessageDeadline)
				rawMsg, err := conn.ReceiveMessage(ctx2)
				cancel()
				if err != nil {
					if pgconn.Timeout(err) {
						continue
					}
					log.Fatal().Err(err).Msg("ReceiveMessage failed")
				}

				if errMsg, ok := rawMsg.(*pgproto3.ErrorResponse); ok {
					log.Fatal().Any("errMsg", errMsg).Msg("received Postgres WAL error")
				}

				msg, ok := rawMsg.(*pgproto3.CopyData)
				if !ok {
					log.Printf("Received unexpected message: %T\n", rawMsg)
					continue
				}

				switch msg.Data[0] {
				case pglogrepl.PrimaryKeepaliveMessageByteID:
					pkm, err := pglogrepl.ParsePrimaryKeepaliveMessage(msg.Data[1:])
					if err != nil {
						log.Fatal().Err(err).Msg("ParsePrimaryKeepaliveMessage failed")
					}
					// log.Println("Primary Keepalive Message =>", "ServerWALEnd:", pkm.ServerWALEnd, "ServerTime:", pkm.ServerTime, "ReplyRequested:", pkm.ReplyRequested)
					if pkm.ServerWALEnd > clientXLogPos {
						clientXLogPos = pkm.ServerWALEnd
					}
					if pkm.ReplyRequested {
						nextStandbyMessageDeadline = time.Time{}
					}

				case pglogrepl.XLogDataByteID:
					xld, err := pglogrepl.ParseXLogData(msg.Data[1:])
					if err != nil {
						log.Fatal().Err(err).Msg("ParseXLogData failed")
					}

					processV2(xld.WALData, relationsV2, typeMap, &inStream)

					if xld.WALStart > clientXLogPos {
						clientXLogPos = xld.WALStart
					}
				}
		}
	}
}

/*
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:48Z","message":"Receive a logical replication message: Begin"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:48Z","message":"Receive a logical replication message: Relation"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:48Z","message":"Receive a logical replication message: Insert"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:48Z","message":"insert for xid 0\n"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:48Z","message":"INSERT INTO public.event_outbox: map[id:1 msg:[8 1 18 17 10 1 49 18 6 97 98 99 100 101 102 24 167 134 206 174 6] topic:items.created]"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:48Z","message":"Receive a logical replication message: Commit"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:34:53Z","message":"Sent Standby status message at 0/156F988\n"}
test-messagerelay-1  | {"level":"debug","time":"2024-02-19T16:35:03Z","message":"Sent Standby status message at 0/156F988\n"}
*/

// abc
func processV2(walData []byte, relations map[uint32]*pglogrepl.RelationMessageV2, typeMap *pgtype.Map, inStream *bool) {
	logicalMsg, err := pglogrepl.ParseV2(walData, *inStream)
	if err != nil {
		log.Fatal().Err(err).Msg("Parse logical replication message")
	}
	log.Printf("Receive a logical replication message: %s", logicalMsg.Type())
	switch logicalMsg := logicalMsg.(type) {
	case *pglogrepl.RelationMessageV2:
		relations[logicalMsg.RelationID] = logicalMsg

	case *pglogrepl.BeginMessage:
		// Indicates the beginning of a group of changes in a transaction. This is only sent for committed transactions. You won't get any events from rolled back transactions.

	case *pglogrepl.CommitMessage:

	case *pglogrepl.InsertMessageV2:
		rel, ok := relations[logicalMsg.RelationID]
		if !ok {
			log.Fatal().Any("id", logicalMsg.RelationID).Msg("unknown relation ID")
		}
		values := map[string]interface{}{}
		for idx, col := range logicalMsg.Tuple.Columns {
			colName := rel.Columns[idx].Name
			switch col.DataType {
			case 'n': // null
				values[colName] = nil
			case 'u': // unchanged toast
				// This TOAST value was not changed. TOAST values are not stored in the tuple, and logical replication doesn't want to spend a disk read to fetch its value for you.
			case 't': //text
				val, err := decodeTextColumnData(typeMap, col.Data, rel.Columns[idx].DataType)
				if err != nil {
					log.Fatal().Err(err).Msg("error decoding column data")
				}
				values[colName] = val
			}
		}
		log.Printf("insert for xid %d\n", logicalMsg.Xid)
		log.Printf("INSERT INTO %s.%s: %v", rel.Namespace, rel.RelationName, values)

	case *pglogrepl.UpdateMessageV2:
		log.Printf("update for xid %d\n", logicalMsg.Xid)
		// ...
	case *pglogrepl.DeleteMessageV2:
		log.Printf("delete for xid %d\n", logicalMsg.Xid)
		// ...
	case *pglogrepl.TruncateMessageV2:
		log.Printf("truncate for xid %d\n", logicalMsg.Xid)
		// ...

	case *pglogrepl.TypeMessageV2:
	case *pglogrepl.OriginMessage:

	case *pglogrepl.LogicalDecodingMessageV2:
		log.Printf("Logical decoding message: %q, %q, %d", logicalMsg.Prefix, logicalMsg.Content, logicalMsg.Xid)

	case *pglogrepl.StreamStartMessageV2:
		*inStream = true
		log.Printf("Stream start message: xid %d, first segment? %d", logicalMsg.Xid, logicalMsg.FirstSegment)
	case *pglogrepl.StreamStopMessageV2:
		*inStream = false
		log.Printf("Stream stop message")
	case *pglogrepl.StreamCommitMessageV2:
		log.Printf("Stream commit message: xid %d", logicalMsg.Xid)
	case *pglogrepl.StreamAbortMessageV2:
		log.Printf("Stream abort message: xid %d", logicalMsg.Xid)
	default:
		log.Printf("Unknown message type in pgoutput stream: %T", logicalMsg)
	}
}

func decodeTextColumnData(mi *pgtype.Map, data []byte, dataType uint32) (interface{}, error) {
	if dt, ok := mi.TypeForOID(dataType); ok {
		return dt.Codec.DecodeValue(mi, dataType, pgtype.TextFormatCode, data)
	}
	return string(data), nil
}

// https://microservices.io/patterns/data/polling-publisher.html
func PostgresPollingPublisher() {

}