package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexolan/stocklet/internal/pkg/errors"
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
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {
	GetOrderById(ctx context.Context, id string) (*pb.Order, error)
	GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error)

	CreateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error)
	UpdateOrder(ctx context.Context, order *pb.Order) error
	DeleteOrderById(ctx context.Context, id string) error
}

// Interface for event methods
// Allows flexibility to have seperate controllers for different messaging systems (e.g. Kafka, NATS, etc)
type EventController interface {
	PrepareConsumer(svc *OrderService) messaging.EventConsumerController

	DispatchCreatedEvent(order *pb.Order)
	DispatchUpdatedEvent(order *pb.Order)
	DispatchDeletedEvent(req *pb.CancelOrderRequest)
}

// Create the order service
func NewOrderService(cfg *ServiceConfig, strCtrl StorageController, evtCtrl EventController) *OrderService {
	// Initialise the protobuf validator
	pbVal, err := protovalidate.New()
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialise protobuf validator")
	}

	// Initialise the service
	return &OrderService{
		StrCtrl: strCtrl,
		EvtCtrl: evtCtrl,
		pbVal: pbVal,
	}
}

func (svc OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Validate the request
	// todo: unwrapping error (if validation error) to expose the argument issues in the response
	if err := svc.pbVal.Validate(req); err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeInvalidArgument, "invalid order request", err)
	}
	
	// Get the order from the storage controller
	order, err := svc.StrCtrl.GetOrderById(ctx, req.GetId())
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
	orders, err := svc.StrCtrl.GetOrdersByCustomerId(ctx, req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetOrdersResponse{Data: orders}, nil
}

func (svc OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// todo: perform validation
	// orders require at least one item
	// ensure items are valid
	// get prices of items


	// create pending order
	// todo: properly implement CreateOrder storage func
	// 	(order items are not added at the moment)
	req.Order.Status = pb.OrderStatus_ORDER_STATUS_PENDING
	order, err := svc.StrCtrl.CreateOrder(ctx, req.Order)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeUnknown, "failed to create order", err)
	}
	
	// dispatch order created event
	svc.EvtCtrl.DispatchCreatedEvent(order)

	// dispatch placed order SAGA event
	// ^^^^ maybe not - Placed Order SAGA will start from the OrderCreated event
	// actually... unless I want to validate the order items before?
	// but that can be done syncronously - at the validation stage
	//
	// or this service could not be publicly exposed - and that is done in a seperate checkout service
	// ... this service (on the public face) would just be for checking order statuses and attempting
	//		order updates
	//
	// checkout service would/could be an orchestrator?
	// checkout service ->
	// 	-> warehouse svc
	//  -> order svc
	//  -> payment svc
	// -> checkout service?

	// return the pending order
	// todo: return a message (/ 'detail') in create order response?
	return &pb.CreateOrderResponse{Data: order}, nil
}

func (svc OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	// idea:
	// on change of thought, have separate API for each editable property?
	

	// alternatively keep as is, patch:
	// look at: https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/patch_feature/

	// Validate the inputs
	// req.Order

	// Merge the order patch in the request with current state
	
	// todo: delete after
	// for testing
	// todo: check field_mask and order object
	b, _ := json.Marshal(req)
	fmt.Print(b)

	// todo: validate merged
	// ... patchedOrder 
	patchedOrder := req.Order

	// Update the order (db level)
	err := svc.StrCtrl.UpdateOrder(ctx, patchedOrder)
	if err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to update order")
	}

	// todo: dispatching created,updated,deleted events at DB level on succesful updates
	// on the storage controller level
	
	/*
	// todo: return updated order
	order, err := svc.DataCtrl.GetOrderById(ctx, req.Order.Id)
	if err != nil {
		return nil, err
	}

	// order items are immutable currently
	// something to change in the future

	// todo: get order at start
	// merge to form patch with existing body
	// check permissions
	// return updated patch on success (instead of double query)
	*/

	return &pb.UpdateOrderResponse{}, nil
}

func (svc OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) DeleteUserData(ctx context.Context, req *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) ProcessPlaceOrderEvent(ctx context.Context, req *pb.PlaceOrderEvent) (*emptypb.Empty, error) {
	// Ignore events dispatched by the order service
	if req.Type == pb.PlaceOrderEvent_TYPE_UNSPECIFIED || req.Status == pb.PlaceOrderEvent_STATUS_UNSPECIFIED {
		return &emptypb.Empty{}, nil
	}
	
	// Mark the order as rejected if a failure status was reported at any stage.
	if req.Status == pb.PlaceOrderEvent_STATUS_FAILURE {
		err := svc.StrCtrl.UpdateOrder(
			context.Background(),
			&pb.Order{
				Id: req.GetPayload().GetOrderId(),
				Status: pb.OrderStatus_ORDER_STATUS_REJECTED,
			},
		)
		if err != nil {
			log.Panic().Any("evt", req).Err(err).Msg("failed to mark order as failed in response to PlaceOrderEvent")
		}

		return &emptypb.Empty{}, nil
	}
	
	// Otherwise, if the event is from the last stage of the saga (shipping svc)
	// ... then mark the order as succesful.
	if req.Type == pb.PlaceOrderEvent_TYPE_SHIPPING {
		err := svc.StrCtrl.UpdateOrder(
			context.Background(),
			&pb.Order{
				Id: req.GetPayload().GetOrderId(),

				// todo: ensuring typesafe - changed proto to field optional 
				// check if checks needed to ensure exists
				TransactionId: req.GetPayload().PaymentId,
				Status: pb.OrderStatus_ORDER_STATUS_APPROVED,
			},
		)
		if err != nil {
			log.Panic().Any("evt", req).Err(err).Msg("failed to mark order as successful in response to PlaceOrderEvent")
		}

		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, nil
}