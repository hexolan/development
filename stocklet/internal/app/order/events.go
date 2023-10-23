package order

import (
	"context"	

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"

	order_v1 "github.com/hexolan/stocklet/internal/pkg/protobuf/order/v1"
)

type evtRepository struct {
	next OrderRepository
	kcl *kgo.Client
}

func NewEventRepository(next OrderRepository, kcl *kgo.Client) OrderRepository {
	return evtRepository{
		next: next,
		kcl: kcl,
	}
}

func (repo evtRepository) dispatchEvent(record *kgo.Record) {
	// todo: test
	ctx := context.Background()
	repo.kcl.Produce(ctx, record, nil)
}

func (repo evtRepository) dispatchCreatedEvent(req *order_v1.CreateOrderRequest, item *order_v1.Order) {
	// todo: test
	// ctx := context.Background()
	record := &kgo.Record{Topic: "order.created", Value: []byte("test")} 
	// repo.kcl.Produce(ctx, record, nil)
	repo.dispatchEvent(record)
}

func (repo evtRepository) dispatchUpdatedEvent(req *order_v1.UpdateOrderRequest, item *order_v1.Order) {
	// todo: 
}

func (repo evtRepository) dispatchDeletedEvent(req *order_v1.DeleteOrderRequest) {
	// todo: 
}

func (repo evtRepository) GetOrder(req *order_v1.GetOrderRequest) (*order_v1.Order, error) {
	return repo.next.GetOrder(req)
}

func (repo evtRepository) CreateOrder(req *order_v1.CreateOrderRequest) (*order_v1.Order, error) {
	order, err := repo.next.CreateOrder(req)

	if err == nil {
		repo.dispatchCreatedEvent(req, order)
	}

	return order, err
}

func (repo evtRepository) UpdateOrder(req *order_v1.UpdateOrderRequest) (*order_v1.Order, error) {
	order, err := repo.next.UpdateOrder(req)

	if err == nil {
		repo.dispatchUpdatedEvent(req, order)
	}

	return order, err
}

func (repo evtRepository) DeleteOrder(req *order_v1.DeleteOrderRequest) error {
	err := repo.next.DeleteOrder(req)
	
	if err == nil {
		repo.dispatchDeletedEvent(req)
	}

	return err
}