package order

import (
	"github.com/jackc/pgx/v5/pgxpool"

	order_v1 "github.com/hexolan/stocklet/internal/pkg/protobuf/order/v1"
)

type dbRepository struct {
	db *pgxpool.Pool
}

func NewDBRepository(db *pgxpool.Pool) OrderRepository {
	return dbRepository{
		db: db,
	}
}

func (repo dbRepository) GetOrder(req *order_v1.GetOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}

func (repo dbRepository) UpdateOrder(req *order_v1.UpdateOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}

func (repo dbRepository) DeleteOrder(req *order_v1.DeleteOrderRequest) error {
	return nil  // tood
}

func (repo dbRepository) CreateOrder(req *order_v1.CreateOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}