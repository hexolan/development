package messagerelay

import (
	"time"
	"errors"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgproto3"
)

// https://github.com/obsidiandynamics/goharvest/wiki/Comparison-of-harvesting-methods


// https://microservices.io/patterns/data/transaction-log-tailing.html
//
// https://www.postgresql.org/docs/current/logicaldecoding-walsender.html
// https://www.postgresql.org/docs/current/logical-replication-architecture.html
func PostgresTailing(ctx context.Context, pCl *pgxpool.Pool) {
	// implemented using pg_output (logical replication connection)

	// ESTABLISH LOGICAL REPLICATION CONNECTION
	// https://pkg.go.dev/github.com/jackc/pglogrepl#section-readme

	// current exper highly derived from demo: (MIT LICENSE)
	// https://github.com/jackc/pglogrepl/blob/master/example/pglogrepl_demo/main.go

	//
	// acquire low level conn from pgconn
	//
	conn, err := pgconn.Connect(ctx, "postgresql://postgres:postgres@test-service-postgres:5432/postgres?sslmode=disable&replication=database")
	// "?replication=database" connection flag is required (for calling IDENTIFY_SYSTEM)
	if err != nil {
		log.Fatal().Err(err).Msg("postgres: failed to open db connection")
	}

	//
	// create publication (if it doesn't exist)
	//
	const slotName string = "messagerelay_testsvc"
	
	result := conn.Exec(ctx, "CREATE PUBLICATION " + slotName + " FOR TABLE event_outbox;")
	_, err = result.ReadAll()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42710" {
			log.Warn().Err(err).Msg("postgres: publication already exists")
		} else {
			log.Panic().Err(err).Msg("postgres: failed to create publication")
		}
	}
	
	//
	// create logical repl w/ pgoutput plugin
	//
	const outputPlugin string = "pgoutput"

	sysident, err := pglogrepl.IdentifySystem(ctx, conn)
	if err != nil {
		log.Panic().Err(err).Msg("postgres: failed to identify database system. could not acquire WAL flush location")
	}

	// create replication slot
	//
	// https://www.postgresql.org/docs/current/protocol-replication.html
	// "Temporary slots are not saved to disk and are automatically dropped on error or when the session has finished."
	_, err = pglogrepl.CreateReplicationSlot(ctx, conn, slotName, outputPlugin, pglogrepl.CreateReplicationSlotOptions{})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42710" {
			log.Warn().Err(err).Msg("postgres: replication slot already exists")
		} else {
			log.Panic().Err(err).Msg("postgres: failed to create replication slot")
		}
	}

	// logcl repl slot
	log.Info().Any("sysident.XLogPos", sysident.XLogPos).Msg("postgres: x log position (used for startLSN on replication)")
	err = pglogrepl.StartReplication(
		ctx,
		conn,
		slotName,
		sysident.XLogPos,
		pglogrepl.StartReplicationOptions{
			PluginArgs: []string{
				"proto_version '2'",
				"publication_names '" + slotName + "'",
				"messages 'true'",
				"streaming 'true'",
			},
		},
	)
	if err != nil {
		log.Fatal().Err(err).Msg("postgres: failed to start logical replication for WAL")
	}

	//
	// begin consuming logcl repl slot
	//

	// do not understand this stuff?
	const standbyTimeout time.Duration = time.Second * 10
	standbyDeadline := time.Now().Add(standbyTimeout)

	clLSN := sysident.XLogPos
	log.Info().Any("value", clLSN).Msg("client LSN: start at sysident.XLogPos")
	relations := map[uint32]*pglogrepl.RelationMessageV2{}
	typeMap := pgtype.NewMap()

	// whenever we get StreamStartMessage we set inStream to true and then pass it to DecodeV2 function
	// on StreamStopMessage we set it back to false
	inStream := false

	for {
		// stop processing main context is cancelled
		select {
			case <-ctx.Done():
				return
			default:
				//
				// processing
				//

				// check if deadline has been reached 
				// to send another standby status message?
				if time.Now().After(standbyDeadline) {
					log.Info().Any("value", clLSN).Msg("client LSN: standby update")
					err = pglogrepl.SendStandbyStatusUpdate(ctx, conn, pglogrepl.StandbyStatusUpdate{WALWritePosition: clLSN})
					if err != nil {
						log.Panic().Err(err).Msg("postgres: failed to send standby status update")
					}

					standbyDeadline = time.Now().Add(standbyTimeout)
				}

				// todo: check my interpretation is correct
				//
				// attempt to recieve another wire message from postgres
				// unless the deadline is reached to send another standby status update
				// before another message is recieved
				tmpCtx, cancel := context.WithDeadline(ctx, standbyDeadline)
				wireMsg, err := conn.ReceiveMessage(tmpCtx)
				cancel()
				if err != nil {
					if pgconn.Timeout(err) {
						continue
					}
					log.Panic().Err(err).Msg("failed to recieve message from postgres (not result of timeout)")
				}

				// check that the wireMsg is not an error response
				if errMsg, ok := wireMsg.(*pgproto3.ErrorResponse); ok {
					log.Panic().Any("errMsg", errMsg).Msg("received Postgres WAL error")
				}

				// check the msg is valid
				msg, ok := wireMsg.(*pgproto3.CopyData)
				if !ok {
					log.Info().Any("msg", wireMsg).Msg("recieved unexpected msg")
					continue
				}

				// determine msg type
				switch msg.Data[0] {
					case pglogrepl.PrimaryKeepaliveMessageByteID:
						pkm, err := pglogrepl.ParsePrimaryKeepaliveMessage(msg.Data[1:])
						if err != nil {
							log.Panic().Err(err).Msg("failed to parse keepalive message")
						}
						
						log.Info().Any("val", pkm.ServerWALEnd).Msg("pkm.ServerWalEnd (before greater than check)")
						if pkm.ServerWALEnd > clLSN {
							clLSN = pkm.ServerWALEnd
							log.Info().Any("value", clLSN).Msg("client LSN: updated to pkm.ServerWALEnd")
						}
						
						if pkm.ReplyRequested {
							standbyDeadline = time.Time{}
						}

					case pglogrepl.XLogDataByteID:
						xld, err := pglogrepl.ParseXLogData(msg.Data[1:])
						if err != nil {
							log.Panic().Err(err).Msg("failed to parse an XLogDataMessage")
						}

						processV2(xld.WALData, relations, typeMap, &inStream)

						log.Info().Any("val", xld.WALStart).Msg("xld.WALStart (before greater than check)")
						if xld.WALStart > clLSN {
							clLSN = xld.WALStart
							log.Info().Any("value", clLSN).Msg("client LSN: updated to xld.WALStart")
						}
				}
		}
	}
}

func processV2(walData []byte, relations map[uint32]*pglogrepl.RelationMessageV2, typeMap *pgtype.Map, inStream *bool) {
	// parse the data to a WAL log data msg
	logicalMsg, err := pglogrepl.ParseV2(walData, *inStream)
	if err != nil {
		log.Panic().Err(err).Msg("failed to parse logical replication msg")
	}
	
	switch logicalMsg := logicalMsg.(type) {
		case *pglogrepl.RelationMessageV2:
			relations[logicalMsg.RelationID] = logicalMsg
		case *pglogrepl.InsertMessageV2:
			rel, ok := relations[logicalMsg.RelationID]
			if !ok {
				log.Panic().Any("id", logicalMsg.RelationID).Msg("unknown relation ID")
			}

			values := map[string]interface{}{}
			for idx, col := range logicalMsg.Tuple.Columns {
				colName := rel.Columns[idx].Name
				switch col.DataType {
				case 'n': // null
					values[colName] = nil
				case 't': //text
					val, err := decodeTextColumnData(typeMap, col.Data, rel.Columns[idx].DataType)
					if err != nil {
						log.Panic().Err(err).Msg("error decoding column data")
					}
					values[colName] = val
				}
			}
			
			log.Printf("INSERT INTO %s.%s: %v", rel.Namespace, rel.RelationName, values)

		case *pglogrepl.StreamStartMessageV2:
			*inStream = true
		case *pglogrepl.StreamStopMessageV2:
			*inStream = false
		default:
			log.Info().Type("msgType", logicalMsg).Msg("unprocessed msg type in pgoutput stream")
	}
}

func decodeTextColumnData(mi *pgtype.Map, data []byte, dataType uint32) (interface{}, error) {
	if dt, ok := mi.TypeForOID(dataType); ok {
		return dt.Codec.DecodeValue(mi, dataType, pgtype.TextFormatCode, data)
	}
	return string(data), nil
}

// https://microservices.io/patterns/data/polling-publisher.html
func PostgresScraping() {

}