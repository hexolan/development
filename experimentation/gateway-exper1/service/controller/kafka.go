package controller

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/development/experimentation/gateway-exper1/service"
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
		// todo: handling
		panic(err)
	}

	return wireEvt
}