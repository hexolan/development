package controller

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"null.hexolan.dev/exp/service"
	pb "null.hexolan.dev/exp/protogen/testingv1"
)

type kafkaController struct {
	kCl *kgo.Client
}

func NewKafkaController(kCl *kgo.Client) service.EventController {
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
		panic(err)
	}

	return wireEvt
}

func (c kafkaController) DispatchCreatedEvent(item *pb.Item) {
	c.dispatchEvent(
		"testing.state.created",
		c.marshalEvent(
			&pb.ItemStateEvent{
				Type: pb.ItemStateEvent_TYPE_CREATED,
				Payload: item,
			},
		),
	)
}

func (c kafkaController) DispatchUpdatedEvent(item *pb.Item) {
	c.dispatchEvent(
		"testing.state.updated",
		c.marshalEvent(
			&pb.ItemStateEvent{
				Type: pb.ItemStateEvent_TYPE_UPDATED,
				Payload: item,
			},
		),
	)
}

func (c kafkaController) DispatchDeletedEvent(item *pb.Item) {
	c.dispatchEvent(
		"testing.state.deleted",
		c.marshalEvent(
			&pb.ItemStateEvent{
				Type: pb.ItemStateEvent_TYPE_DELETED,
				Payload: item,
			},
		),
	)
}