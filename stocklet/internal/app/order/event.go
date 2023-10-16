package order

import (
	order_v1 "github.com/hexolan/stocklet/internal/pkg/protobuf/order/v1"
)

type evtRepository struct {
	next OrderRepository
}

func NewEventRepository(next OrderRepository) OrderRepository {
	return evtRepository{
		next: next,
	}
}

func (svc evtRepository) GetOrder(req *order_v1.GetOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}

func (svc evtRepository) UpdateOrder(req *order_v1.UpdateOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}

func (svc evtRepository) DeleteOrder(req *order_v1.DeleteOrderRequest) error {
	return nil // todo:
}

func (svc evtRepository) CreateOrder(req *order_v1.CreateOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}