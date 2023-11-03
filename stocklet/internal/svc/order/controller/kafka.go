package controller

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type KafkaProducer struct {
	kcl *kgo.Client
}

func NewKafkaProducer(kcl *kgo.Client) KafkaProducer {
	return KafkaProducer{kcl: kcl}
}

func (p KafkaProducer) dispatchEvent(record *kgo.Record) {
	// todo: test
	ctx := context.Background()
	p.kcl.Produce(ctx, record, nil)
}

func (p KafkaProducer) DispatchCreatedEvent(req *pb.CreateOrderRequest, item *pb.Order) {
	// todo: test
	// ctx := context.Background()
	record := &kgo.Record{Topic: "order.created", Value: []byte("test")}
	// repo.kcl.Produce(ctx, record, nil)
	p.dispatchEvent(record)
}

func (p KafkaProducer) DispatchUpdatedEvent(req *pb.UpdateOrderRequest, item *pb.Order) {
	// todo:
}

func (p KafkaProducer) DispatchDeletedEvent(req *pb.CancelOrderRequest) {
	// todo:
}
