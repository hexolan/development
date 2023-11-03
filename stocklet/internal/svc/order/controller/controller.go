package controller

import (
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type DataController interface {
	GetOrder() error
}

type EventController interface {
	DispatchCreatedEvent(req *pb.CreateOrderRequest, item *pb.Order)
	DispatchUpdatedEvent(req *pb.UpdateOrderRequest, item *pb.Order)
	DispatchDeletedEvent(req *pb.CancelOrderRequest)
}