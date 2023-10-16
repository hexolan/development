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