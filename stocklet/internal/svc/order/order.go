package order

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

// Interface for database methods
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type DataController interface {
	GetOrderById(ctx context.Context, id string) (*pb.Order, error)
	GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error)

	UpdateOrderById(ctx context.Context, id string) ([]*pb.Order, error) // todo: update args
	DeleteOrderById(ctx context.Context, id string) error
}

// Interface for event related methods
// Allows implementing seperate controllers for different messaging systems (e.g. Kafka, NATS, etc)
type EventController interface {
	DispatchCreatedEvent(order *pb.Order)
	DispatchUpdatedEvent(order *pb.Order)
	DispatchDeletedEvent(req *pb.CancelOrderRequest)
}

// The implemented order service served as a gRPC service
type OrderService struct {
	DataCtrl DataController
	EvtCtrl EventController

	pb.UnimplementedOrderServiceServer
}

func NewOrderService(dataCtrl DataController, evtCtrl EventController) OrderService {
	return OrderService{
		DataCtrl: dataCtrl,
		EvtCtrl: evtCtrl,
	}
}

func (svc OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Validate args - todo:
	if req.GetOrderId() == "" {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid order id")
	}
	
	// Get the order
	order, err := svc.DataCtrl.GetOrderById(ctx, req.GetOrderId())
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
	orders, err := svc.DataCtrl.GetOrdersByCustomerId(ctx, req.GetCustomerId())
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

func (svc OrderService) ProcessPlaceOrderEvent(ctx context.Context, req *pb.PlaceOrderEvent) (*emptypb.Empty, error) {
	// todo: improve

	// Ignore events dispatched by the order service
	if req.Type == pb.PlaceOrderEvent_TYPE_UNSPECIFIED || req.Status == pb.PlaceOrderEvent_STATUS_UNSPECIFIED {
		return &emptypb.Empty{}, nil
	}

	// Mark the order as rejected if a failure status was reported at any stage.
	if req.Status == pb.PlaceOrderEvent_STATUS_FAILURE {
		// todo
		return &emptypb.Empty{}, nil
	}
	
	// Otherwise, if the event is from the last stage of the saga (shipping svc)
	// ... then mark the order as succesful.
	if req.Type == pb.PlaceOrderEvent_TYPE_SHIPPING {
		// todo: update order status and details
		// (append transaction id, etc to stored order)
	}

	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}