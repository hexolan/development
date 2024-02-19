package main

import (
	"time"
	"errors"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/proto"

	pb "null.hexolan.dev/dev/protogen/testpb"
)

const pgItemBaseQuery string = "SELECT id, name, created_at, updated_at FROM items"

type DatabaseInterface interface {
	GetById(ctx context.Context, id string) (*pb.TestItem, error)
	CreateItem(ctx context.Context, item *pb.TestItem) (*pb.TestItem, error)
	StartProducer()
}
type databaseInterface struct {
	pCl *pgxpool.Pool
}

func newPostgresConn() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	pCl, err := pgxpool.New(ctx, "postgresql://postgres:postgres@test-service-postgres:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return pCl, nil
}

func NewDatabaseInterface() DatabaseInterface {
	pCl, err := newPostgresConn()
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	return &databaseInterface{pCl: pCl}
}

func scanRowToItem(row pgx.Row) (*pb.TestItem, error) {
	var item pb.TestItem

	// convert from postgres timestamp format to int64 unix timestamp
	var tmpCreatedAt pgtype.Timestamp
	var tmpUpdatedAt pgtype.Timestamp

	err := row.Scan(
		&item.Id,
		&item.Name,
		&tmpCreatedAt,
		&tmpUpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// convert timestamp to unix
	if tmpCreatedAt.Valid {
		item.CreatedAt = tmpCreatedAt.Time.Unix()
	} else {
		return nil, errors.New("failed to convert created_at timestamp")
	}

	if tmpUpdatedAt.Valid {
		item.UpdatedAt = tmpUpdatedAt.Time.Unix()
	}

	return &item, nil
}

func (db *databaseInterface) StartProducer() {
	ep := NewEventProducer(db.pCl)
	ep.Start()
}

func (db *databaseInterface) GetById(ctx context.Context, id string) (*pb.TestItem, error) {
	row := db.pCl.QueryRow(ctx, pgItemBaseQuery + " WHERE id=$1", id)

	item, err := scanRowToItem(row)
	return item, err
}

func (db *databaseInterface) CreateItem(ctx context.Context, item *pb.TestItem) (*pb.TestItem, error) {
	// Begin a DB transaction
	tx, err := db.pCl.Begin(ctx)
	if err != nil {
		return nil, errors.Join(errors.New("failed to begin transaction"), err)
	}
	defer tx.Rollback(ctx)

	// Insert order query
	var id string
	err = tx.QueryRow(
		ctx,
		"INSERT INTO items (name) VALUES ($1) RETURNING id",
		item.Name,
	).Scan(&id)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create item"), err)
	}

	// add missing attrs to item
	item.Id = id
	item.CreatedAt = time.Now().Unix()

	// event outbox table
	// add order created event to the outbox
	wireEvent, err := proto.Marshal(&pb.ItemStateEvent{
		Type: pb.ItemStateEvent_TYPE_CREATED,
		Payload: item,
	})
	if err != nil {
		// todo: proper error handling
		log.Panic().Err(err).Msg("error marshaling protobuf event")
	}

	_, err = tx.Exec(
		ctx,
		"INSERT INTO event_outbox (topic, msg) VALUES ($1, $2)",
		"items.created",
		wireEvent,
	)
	if err != nil {
		return nil, errors.Join(errors.New("failed to add item event"), err)
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.Join(errors.New("failed to commit transaction"), err)
	}

	// Return the created item
	return item, nil
}