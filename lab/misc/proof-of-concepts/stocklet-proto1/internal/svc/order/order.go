package order

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/mennanov/fmutils"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

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
//
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {
	GetOrderById(ctx context.Context, id string) (*pb.Order, error)
	GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error)

	CreateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error)
	UpdateOrder(ctx context.Context, order *pb.Order, mask *fieldmaskpb.FieldMask) error
	DeleteOrderById(ctx context.Context, id string) error
}

// Interface for event methods
//
// Methods for assembling this service's events.
type EventBuilder interface {
	PrepareCreatedEvent(order *pb.Order) (string, []byte)
	PrepareUpdatedEvent(order *pb.Order) (string, []byte)
	PrepareDeletedEvent(req *pb.CancelOrderRequest) (string, []byte)
}

// Event controller is responsible for dispatching.
//
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
	// Validate the request args
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
	// Validate the request args
	// todo: unwrapping error (if validation error) to expose the argument issues in the response
	if err := svc.pbVal.Validate(req); err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeInvalidArgument, "invalid orders request", err)
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
		log.Error().Err(err).Msg("order err")
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

// Update an order.
//
// Uses the provided field mask as a selector
// for which fields to update.
//
// Effectively operates as a PATCH function.
func (svc OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	// Validating the core request
	if req.Order == nil || req.Mask == nil {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "malformed request")
	}

	// todo: ensuring req.Mask is real and req.Order is real

	// TODO:
	// validating all the inputs specified in the UpdateMask
	// ensuring that they are VALID and ALLOWED fields
	//
	// protovalidate has no support for FieldMask
	// think of solution to pull validation for individual fields
	// 	defined in the proto file for the Order type

	// Take a copy of orderId
	// The `req.Order` object will be overriden when filtering by field mask
	orderId := req.Order.Id

	// Filter to only field mask elements 
	// TODO: ensuring that Id, Items, CreatedAt and UpdatedAt cannot be specified
	// items will be handled in a seperate str ctrl call
	maskPaths := req.Mask.GetPaths()
	log.Info().Any("paths", maskPaths).Msg("abc")
	fmutils.Filter(req.Order, maskPaths)

	// protovalidate has no support for FieldMask
	//
	// think of solution to pull validation for individual fields
	// so I only validate the fields specified in the field mask
	//
	// this will currently fail unless an entire PUT request is taking place
	//
	// although I still need to add checks to ensure that ID, CreatedAt and UpdatedAt are not being updated
	// maliciously in that case
	//
	// since this is a user-facing operation, it cannot be willy-nilly allowed to update
	// order.Status and etc...
	// users will have to call a seperate UserOrders / OrderHandler service or something requesting they cancel their order
	//
	// need to think of how authentication is going to play out
	// already have methods of checking if a request has come through gRPC gateway (so they can skip any user validation)
	// - but for handling internally (aside events) - maybe there should be an OrderHandler service which users use to place their order
	// 		> that can act as an orchestrator or whatever - this can still be called directly to view order details
	//
	// maybe also split services into:
	// Order Service (business and validation logic)
	//  > handles inbound gRPC calls
	//  > performs validation before making data service calls
	//
	// ((( AS USING EVENT DRIVEN DESIGN )))
	// 		Order service doesn't even need to call the Order Data Service via gRPC
	// 		- the order service performs validation and dispatches created, updated and deleted events
	// 		- the order consumer will then TELL the order data service to enact that event (however stipulations for generating unique ids)
	// 		- for updates this could be entirely fine
	//
	// Order Data Service - called by Order Service thru gRPC - stores/retrieves stuff in databases // can also split into order-postgres-data-svc, order-mongo-data-svc
	// 	> also acts as an outbox for Created, Updated, Deleted events
	//  > as it is acting as an outbox (functionality also required by main Order Service)
    //  	could splinter another service for Order Producer // split into order-kafka-producer, order-nats-producer
	//
	// Order Consumer (Service) - calls the Order Service upon reciept of events // can also split into order-kafka-consumer, order-nats-consumer
	//
	// This does mean that there is tight coupling (runtime-wise) between these services though.
	//
	// However there are benefits to having a separate data service // load balancing or sending repeat requests to different instances of data services (1 req -> Postgres and Mongo instances)
	// In addition, requests could be routed by order ids - can coalesce requests
	if err := svc.pbVal.Validate(req.Order); err != nil {
		log.Error().Err(err).Msg("invalid order update err")
		return nil, errors.WrapServiceError(errors.ErrCodeInvalidArgument, "invalid order update", err)
	}

	// Update the order (database level)
	// TODO: remove after (temp)
	req.Order.Id = orderId
	err := svc.StrCtrl.UpdateOrder(ctx, req.Order, req.Mask)
	if err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to update order")
	}
	
	// todo: reliably sending created/updated/deleted events
	// https://microservices.io/patterns/data/transactional-outbox.html
	// https://batalin.dev/posts/transaction-outbox-pattern/
	//
	// https://docs.aws.amazon.com/prescriptive-guidance/latest/cloud-design-patterns/transactional-outbox.html
	//
	// - -
	// https://www.baeldung.com/transactions-across-microservices
	//
	// CREATING SEPARATE EXPER

	// get the updated order
	order, err := svc.StrCtrl.GetOrderById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	// todo: dispatching created,updated,deleted events at DB level on succesful updates
	// on the storage controller level

	return &pb.UpdateOrderResponse{Data: order}, nil
}

func (svc OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) DeleteUserData(ctx context.Context, req *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc OrderService) ProcessPlaceOrderEvent(ctx context.Context, req *pb.PlaceOrderEvent) (*emptypb.Empty, error) {
	// todo: properly validating the request?

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
			&fieldmaskpb.FieldMask{Paths: []string{"status"}},
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
			&fieldmaskpb.FieldMask{Paths: []string{"transaction_id", "status"}},
		)
		if err != nil {
			log.Panic().Any("evt", req).Err(err).Msg("failed to mark order as successful in response to PlaceOrderEvent")
		}

		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, nil
}