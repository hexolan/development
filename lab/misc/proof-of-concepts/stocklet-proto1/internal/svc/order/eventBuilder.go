package order

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type eventBuilder struct {}

var _ EventBuilder = (*eventBuilder)(nil)

func (eb eventBuilder) marshalProto(evt protoreflect.ProtoMessage) []byte {
	wireEvt, err := proto.Marshal(evt)
	if err != nil {
		// todo: proper error handling
		log.Panic().Err(err).Msg("error marshaling protobuf event")
	}

	return wireEvt
}

func (eb eventBuilder) PrepareCreatedEvent(order *pb.Order) (string, []byte) {
	topic := messaging.Order_State_Created_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_CREATED,
		Payload: order,
	}

	return topic, eb.marshalProto(event)
}

func (eb eventBuilder) PrepareUpdatedEvent(order *pb.Order) (string, []byte) {
	topic := messaging.Order_State_Updated_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_UPDATED,
		Payload: order,
	}

	return topic, eb.marshalProto(event)
}

func (eb eventBuilder) PrepareDeletedEvent(req *pb.CancelOrderRequest) (string, []byte) {
	topic := messaging.Order_State_Deleted_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_DELETED,
		Payload: &pb.Order{
			Id: req.GetId(),
		},
	}

	return topic, eb.marshalProto(event)
}