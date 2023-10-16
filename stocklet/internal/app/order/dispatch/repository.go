package dispatch

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/app/order"
)

type eventRepository struct {
	next *order.OrderRepository
	prod EventProducer
}

func NewEventRepository(next *order.OrderRepository, prod EventProducer) order.OrderRepository {
	return eventRepository{
		next: next,
	}
}