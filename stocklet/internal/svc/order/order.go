package order

import (
	"context"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

// Generic database access methods
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type DataController interface {
	GetOrderById(ctx context.Context, id string) (*pb.Order, error)
	GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error)

	UpdateOrderById(ctx context.Context, id string) ([]*pb.Order, error) // todo: update args
	DeleteOrderById(ctx context.Context, id string) error
}

// Generic event methods
// Allows implementing seperate event controllers for different messaging systems (e.g. Kafka, NATS, etc)
type EventController interface {
	DispatchCreatedEvent(req *pb.CreateOrderRequest, item *pb.Order)
	DispatchUpdatedEvent(req *pb.UpdateOrderRequest, item *pb.Order)
	DispatchDeletedEvent(req *pb.CancelOrderRequest)

	ProcessPlaceOrderEvent(evt *pb.PlaceOrderEvent) // todo: move into service? - dispatch place order origin event
}