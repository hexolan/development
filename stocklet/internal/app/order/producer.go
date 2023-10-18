package order

import (
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
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

func (repo evtRepository) GetOrder(req *order_v1.GetOrderRequest) (*order_v1.Order, error) {
	return repo.next.GetOrder(req)
}

func (repo evtRepository) UpdateOrder(req *order_v1.UpdateOrderRequest) (*order_v1.Order, error) {
	order, err := repo.next.UpdateOrder(req)

	if err == nil {
		// todo: dispatch order updated event
	}

	return order, err
}

func (repo evtRepository) DeleteOrder(req *order_v1.DeleteOrderRequest) error {
	err := repo.next.DeleteOrder(req)
	
	if err == nil {
		// todo: dispatch order deleted event
	}

	return err
}

func (repo evtRepository) CreateOrder(req *order_v1.CreateOrderRequest) (*order_v1.Order, error) {
	order, err := repo.next.CreateOrder(req)

	if err == nil {
		// todo: dispatch order created event
	}

	return order, err
}