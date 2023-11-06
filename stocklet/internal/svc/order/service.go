package order

import (
	"context"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type orderService struct {
	dbC DataController
	evtC EventController

	pb.UnimplementedOrderServiceServer
}

func NewOrderService(evtC EventController, dbC DataController) orderService {
	return orderService{
		dbC: dbC,
		evtC: evtC,
	}
}

func (svc orderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Validate args - todo:
	if req.GetOrderId() == "" {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid order id")
	}
	
	// Get the order
	order, err := svc.dbC.GetOrderById(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	
	return &pb.GetOrderResponse{Data: order}, nil
}

func (svc orderService) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	// Validate args - todo:
	if req.GetCustomerId() == "" {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid customer id")
	}
	
	// Get the orders
	orders, err := svc.dbC.GetOrdersByCustomerId(ctx, req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetOrdersResponse{Data: orders}, nil
}

func (svc orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// create pending order - dispatch created event

	// dispatch placed order SAGA event

	// return pending order

	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc orderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	// update order (db level)

	// if succesful dispatch updated event
	
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc orderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc orderService) DeleteUserData(ctx context.Context, req *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}