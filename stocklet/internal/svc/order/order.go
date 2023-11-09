package order

import (
	"context"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

// todo: tidy up
// main focus is on implementing core functionality then clean up the mess and do documentation, etc

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

// The implemented service
type OrderService struct {
	DbC DataController
	EvtC EventController

	pb.UnimplementedOrderServiceServer
}

func NewOrderService(dbC DataController, evtC EventController) OrderService {
	return OrderService{
		DbC: dbC,
		EvtC: evtC,
	}
}

func (svc OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Validate args - todo:
	if req.GetOrderId() == "" {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid order id")
	}
	
	// Get the order
	order, err := svc.DbC.GetOrderById(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	
	return &pb.GetOrderResponse{Data: order}, nil
}

func (svc OrderService) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	// Validate args - todo:
	if req.GetCustomerId() == "" {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid customer id")
	}
	
	// Get the orders
	orders, err := svc.DbC.GetOrdersByCustomerId(ctx, req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetOrdersResponse{Data: orders}, nil
}

func (svc OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// create pending order - dispatch created event

	// dispatch placed order SAGA event

	// return pending order

	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	// update order (db level)

	// if succesful dispatch updated event
	
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) DeleteUserData(ctx context.Context, req *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}