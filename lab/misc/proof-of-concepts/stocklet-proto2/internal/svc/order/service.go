// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package order

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/mennanov/fmutils"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

// Create the order service
func NewOrderService(cfg *ServiceConfig, strCtrl StorageController, evtCtrl EventController) *OrderService {
	// Initialise the protobuf validator
	pbVal, err := protovalidate.New()
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialise protobuf validator")
	}

	// Initialise the service
	return &OrderService{
		pbVal: pbVal,
		StrCtrl: strCtrl,
		EvtCtrl: evtCtrl,
	}
}

func (svc OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Validate the request args
	//
	// Provide the validation error to the user.
	if err := svc.pbVal.Validate(req); err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid request: " + err.Error())
	}
	
	// Get the order from the storage controller
	order, err := svc.StrCtrl.GetOrder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	
	return &pb.GetOrderResponse{Data: order}, nil
}

func (svc OrderService) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	// Validate the request args
	if err := svc.pbVal.Validate(req); err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "invalid request: " + err.Error())
	}
	
	// Get the orders from the storage controller
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


	// Create a pending order
	//
	// This will initiate a SAGA process involving
	// the services required to create the order.
	req.Order.Status = pb.OrderStatus_ORDER_STATUS_PENDING
	order, err := svc.StrCtrl.CreateOrder(ctx, req.Order)
	if err != nil {
		log.Error().Err(err).Msg("failed to create order")
		return nil, errors.WrapServiceError(errors.ErrCodeUnknown, "failed to create order", err)
	}

	// Return the pending order
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
	//
	// Ensure that order and mask objects have been provided.
	if req.Order == nil || req.Mask == nil {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "malformed request")
	}

	// Take a copy of orderId
	// The `req.Order` object will be overriden when filtering by field mask
	orderId := req.Order.Id

	// Validating all the inputs specified in the mask
	// ensuring that they are VALID and ALLOWED fields
	//
	// protovalidate does not ship with support for skipping 
	// validation of fields using a field mask.
	// GH Issue: https://github.com/bufbuild/protovalidate/issues/112
	// think of solution to pull validation for individual fields
	// defined in the proto file for the Order type

	// Validate the mask paths
	maskPaths := req.Mask.GetPaths()
	if len(maskPaths) == 0 {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "no fields specified")
	}
	
	// TODO: ensuring that Id, Items, CreatedAt and UpdatedAt cannot be specified
	// items will be handled in a seperate str ctrl call

	// Filter to only field mask elements 
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

	// since this is a user-facing operation, it cannot be willy-nilly allowed to update
	// order.Status and etc...

	if err := svc.pbVal.Validate(req.Order); err != nil {
		log.Error().Err(err).Msg("invalid order update err")
		return nil, errors.WrapServiceError(errors.ErrCodeInvalidArgument, "invalid order update", err)
	}

	// Update the order
	order, err := svc.StrCtrl.UpdateOrder(ctx, orderId, req.Order, req.Mask)
	if err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to update order")
	}

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
		_, err := svc.StrCtrl.UpdateOrder(
			context.Background(),
			req.GetPayload().GetOrderId(),
			&pb.Order{Status: pb.OrderStatus_ORDER_STATUS_REJECTED},
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
		_, err := svc.StrCtrl.UpdateOrder(
			context.Background(),
			req.GetPayload().GetOrderId(),
			&pb.Order{
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