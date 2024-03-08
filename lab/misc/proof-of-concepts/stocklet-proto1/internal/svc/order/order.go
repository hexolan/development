package order

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

// Interface for the service
type OrderService struct {
	pb.UnimplementedOrderServiceServer
	
	pbVal *protovalidate.Validator

	StrCtrl StorageController
	EvtCtrl EventController
}

// Interface for database methods
//
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {
	GetOrderById(ctx context.Context, orderId string) (*pb.Order, error)
	GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error)

	CreateOrder(ctx context.Context, orderObj *pb.Order) (*pb.Order, error)
	UpdateOrder(ctx context.Context, orderId string, orderObj *pb.Order, mask *fieldmaskpb.FieldMask) error
	DeleteOrderById(ctx context.Context, id string) error
}

// Interface for the messaging methods
//
// Allows flexibility to have seperate controllers for different messaging systems (e.g. Kafka, NATS, etc)
type EventController interface {
	PrepareConsumer(svc *OrderService) messaging.EventConsumerController
}