package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseRepo struct {
	db *pgxpool.Pool
}

func NewDatabaseRepo(db *pgxpool.Pool) DatabaseRepo {
	return DatabaseRepo{db: db}
}
