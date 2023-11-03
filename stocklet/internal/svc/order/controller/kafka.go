package controller

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type kafkaController struct {
	kcl *kgo.Client
}

func NewKafkaController(kcl *kgo.Client) EventController {
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
