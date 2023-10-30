package order

import (
	"context"	

	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type EventProducer struct {
	kcl *kgo.Client
}

func NewEventProducer(kcl *kgo.Client) EventProducer {
	return EventProducer{kcl: kcl}
}

func (prod EventProducer) dispatchEvent(record *kgo.Record) {
	// todo: test
	ctx := context.Background()
	prod.kcl.Produce(ctx, record, nil)
}

func (prod EventProducer) DispatchCreatedEvent(req *order_v1.CreateOrderRequest, item *order_v1.Order) {
	// todo: test
	// ctx := context.Background()
	record := &kgo.Record{Topic: "order.created", Value: []byte("test")} 
	// repo.kcl.Produce(ctx, record, nil)
	prod.dispatchEvent(record)
}

func (prod EventProducer) DispatchUpdatedEvent(req *order_v1.UpdateOrderRequest, item *order_v1.Order) {
	// todo: 
}

func (prod EventProducer) DispatchDeletedEvent(req *order_v1.CancelOrderRequest) {
	// todo: 
}