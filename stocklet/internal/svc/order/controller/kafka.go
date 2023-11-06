package controller

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/svc/order"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type kafkaController struct {
	kcl *kgo.Client
}

func NewKafkaController(kcl *kgo.Client) order.EventController {
	return kafkaController{kcl: kcl}
}

func (c kafkaController) dispatchEvent(record *kgo.Record) {
	// todo: test
	ctx := context.Background()
	c.kcl.Produce(ctx, record, nil)
}

func (c kafkaController) DispatchCreatedEvent(req *pb.CreateOrderRequest, item *pb.Order) {
	// todo: test
	// ctx := context.Background()
	record := &kgo.Record{Topic: "order.created", Value: []byte("test")}
	// repo.kcl.Produce(ctx, record, nil)
	c.dispatchEvent(record)
}

func (c kafkaController) DispatchUpdatedEvent(req *pb.UpdateOrderRequest, item *pb.Order) {
	// todo:
}

func (c kafkaController) DispatchDeletedEvent(req *pb.CancelOrderRequest) {
	// todo:
}

func (c kafkaController) ProcessPlaceOrderEvent(event *pb.PlaceOrderEvent) {
	// Ignore events dispatched by the order service
	if event.Type == pb.PlaceOrderEvent_TYPE_UNSPECIFIED || event.Status == pb.PlaceOrderEvent_STATUS_UNSPECIFIED {
		return
	}

	// Mark the order as rejected if a failure status
	// was reported at any stage.
	if event.Status == pb.PlaceOrderEvent_STATUS_FAILURE {
		// todo
		return
	}
	
	// Otherwise,
	// If the event is from the last stage of the saga (shipping svc)
	// then mark the order as succesful.
	if event.Type == pb.PlaceOrderEvent_TYPE_SHIPPING {
		// todo: update order status and details
		// (append transaction id, etc to stored order)
	}
}