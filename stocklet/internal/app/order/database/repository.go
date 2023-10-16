package database

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/app/order"
)

type dbRepository struct {
	db *pgxpool.Pool
}

func NewDBRepository(db *pgxpool.Pool) order.OrderRepository {
	return dbRepository{
		db: db,
	}
}