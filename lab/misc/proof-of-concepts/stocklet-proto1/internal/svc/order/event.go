package order

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func marshalEvent(event protoreflect.ProtoMessage) ([]byte, error) {
	wireEvt, err := proto.Marshal(event)
	if err != nil {
		return []byte{}, err
	}

	return wireEvt, nil
}

func PrepareCreatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Created_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_CREATED,
		Payload: order,
	}

	wireEvt, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvt, topic, nil
}

func PrepareUpdatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Updated_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_UPDATED,
		Payload: order,
	}

	wireEvt, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvt, topic, nil
}

func PrepareDeletedEvent(orderId string) ([]byte, string, error) {
	topic := messaging.Order_State_Deleted_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_DELETED,
		Payload: &pb.Order{
			Id: orderId,
		},
	}

	wireEvt, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvt, topic, nil
}