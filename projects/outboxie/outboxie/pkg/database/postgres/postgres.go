// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package postgres

import (
	"errors"
	"context"

	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/hexolan/outboxie/pkg/database"
)

type DriverConfig struct {
	Url string

	OutboxTableName string

	EnsurePublication bool
	EnsureSlot bool

	SlotName string
}

type PostgresDriver struct {
	conn *pgconn.PgConn
}

// postgres driver for push-based watching using logical replication
func NewPostgresDriver() database.Driver {
	return &PostgresDriver{}
}

func (dvr *PostgresDriver) Setup(ctx context.Context, conf DriverConfig) (database.Driver, error) {
	// todo: ensuring "?replication=database" is there

	// Acquire a low level connection for postgres
	conn, err := pgconn.Connect(ctx, conf.Url)
	if err != nil {
		return nil, errors.Join(errors.New("postgres: failed to open connection"), err)
	}
	defer conn.Close(ctx)

	// todo: commenting & testing & expanding upon + improving config opts
	if (conf.EnsurePublication) {
		createPublication(ctx, conn, conf.SlotName, conf.OutboxTableName)
	}
	
	// todo: commenting & testing & expanding upon
	if (conf.EnsureSlot) {
		createReplicationSlot(ctx, conn, conf.SlotName)
	}

	// Identify system check
	_, err = pglogrepl.IdentifySystem(ctx, conn)
	if err != nil {
		return nil, errors.Join(errors.New("failed to identify postgres system (ensure '?replication=database' flag on conn string)"), err)
	}

	return &PostgresDriver{conn: conn}, nil
}

func (dvr *PostgresDriver) Start(callback *database.ReceiverCallback) error {
	return nil
}

func (dvr *PostgresDriver) Stop() error {
	return nil
}


////

func createPublication(ctx context.Context, conn *pgconn.PgConn, name string, table string) error {
	result := conn.Exec(ctx, "CREATE PUBLICATION " + name + " FOR TABLE " + table + ";")
	_, err := result.ReadAll()
	if ignorePgAlreadyExists(err) != nil {
		return errors.Join(errors.New("postgres: failed to create publication"), err)
	}
}

func createReplicationSlot(ctx context.Context, conn *pgconn.PgConn, name string) error {
	// todo: taking temporary slot option
	_, err := pglogrepl.CreateReplicationSlot(ctx, conn, name, "pgoutput", pglogrepl.CreateReplicationSlotOptions{})
	if ignorePgAlreadyExists(err) != nil {
		return errors.Join(errors.New("postgres: failed to create replication slot"), err)
	}
}

func ignorePgAlreadyExists(err error) error {
	if err == nil {
		return err
	}
	
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "42710" {
		return nil
	}

	return err
}