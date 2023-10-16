package order

import (
	order_v1 "github.com/hexolan/stocklet/internal/pkg/protobuf/order/v1"
)

type OrderRepository interface {
	CreateOrder(*order_v1.CreateOrderRequest) (*order_v1.Order, error)
	GetOrder(*order_v1.GetOrderRequest) (*order_v1.Order, error)
	UpdateOrder(*order_v1.UpdateOrderRequest) (*order_v1.Order, error)
	DeleteOrder(*order_v1.DeleteOrderRequest) error
}

type svcRepository struct {
	next OrderRepository
}

func NewOrderService(next OrderRepository) OrderRepository {
	return svcRepository{
		next: next,
	}
}

func (svc svcRepository) GetOrder(req *order_v1.GetOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}

func (svc svcRepository) UpdateOrder(req *order_v1.UpdateOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}

func (svc svcRepository) DeleteOrder(req *order_v1.DeleteOrderRequest) error {
	return nil // todo: 
}

func (svc svcRepository) CreateOrder(req *order_v1.CreateOrderRequest) (*order_v1.Order, error) {
	return new(order_v1.Order), nil // todo:
}