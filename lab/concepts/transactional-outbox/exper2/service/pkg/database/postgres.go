package database

import (
	"time"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConn(dsn string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	pCl, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return pCl, nil
}