package controller

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type kafkaController struct {
	kCl *kgo.Client
}

func NewKafkaController(kCl *kgo.Client) order.EventController {
	return kafkaController{kCl: kCl}
}

func (c kafkaController) dispatchEvent(record *kgo.Record) {
	// todo: test
	ctx := context.Background()
	c.kCl.Produce(ctx, record, nil)
}

func (c kafkaController) marshalEvent(evt protoreflect.ProtoMessage) []byte {
	wireEvt, err := proto.Marshal(evt)
	if err != nil {
		// todo: handling
		panic(err)
	}

	return wireEvt
}

func (c kafkaController) DispatchCreatedEvent(order *pb.Order) {
	record := &kgo.Record{
		Topic: messaging.Order_State_Created_Topic,
		Value: c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_CREATED,
				Payload: order,
			},
		),
	}

	c.dispatchEvent(record)
}

func (c kafkaController) DispatchUpdatedEvent(order *pb.Order) {
	record := &kgo.Record{
		Topic: messaging.Order_State_Updated_Topic,
		Value: c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_UPDATED,
				Payload: order,
			},
		),
	}

	c.dispatchEvent(record)
}

func (c kafkaController) DispatchDeletedEvent(req *pb.CancelOrderRequest) {
	// todo: improve assembly of payload (dispatch whole order?)

	record := &kgo.Record{
		Topic: messaging.Order_State_Deleted_Topic,
		Value: c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_DELETED,
				Payload: &pb.Order{
					Id: req.GetOrderId(),
				},
			},
		),
	}

	c.dispatchEvent(record)
}

func (c kafkaController) DispatchPlaceOrderEvent(evt *pb.PlaceOrderEvent) {
	// todo:
	record := &kgo.Record{
		Topic: messaging.Order_PlaceOrder_Order_Topic,
		Value: c.marshalEvent(evt),
	}

	c.dispatchEvent(record)
}