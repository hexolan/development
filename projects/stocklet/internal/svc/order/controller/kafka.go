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

func (c kafkaController) dispatchEvent(topic string, wireEvt []byte) {
	ctx := context.Background()

	c.kCl.Produce(
		ctx,
		&kgo.Record{
			Topic: topic,
			Value: wireEvt,
		},
		nil,
	)
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
	c.dispatchEvent(
		messaging.Order_State_Created_Topic,
		c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_CREATED,
				Payload: order,
			},
		),
	)
}

func (c kafkaController) DispatchUpdatedEvent(order *pb.Order) {
	c.dispatchEvent(
		messaging.Order_State_Updated_Topic,
		c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_UPDATED,
				Payload: order,
			},
		),
	)
}

func (c kafkaController) DispatchDeletedEvent(req *pb.CancelOrderRequest) {
	// todo: improve assembly of payload (dispatch whole order?)
	c.dispatchEvent(
		messaging.Order_State_Deleted_Topic,
		c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_DELETED,
				Payload: &pb.Order{
					Id: req.GetOrderId(),
				},
			},
		),
	)
}

func (c kafkaController) DispatchPlaceOrderEvent(evt *pb.PlaceOrderEvent) {
	// todo:
	c.dispatchEvent(
		messaging.Order_PlaceOrder_Order_Topic,
		c.marshalEvent(evt),
	)
}