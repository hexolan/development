package main

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// https://thegodev.com/transactional-outbox-pattern/
// https://medium.com/@adamszpilewicz/outbox-pattern-in-go-with-postgres-alternative-to-rabbitmq-aa203d9cc9d0

// specific tools:
// * https://github.com/obsidiandynamics/goharvest
//   > specific to Postgres -> Kafka (will not fit usecase)

type EventProducer interface {
	Start()
}
type eventProducer struct {
	pCl *pgxpool.Pool
}

func NewEventProducer(pCl *pgxpool.Pool) EventProducer {
	return &eventProducer{}
}

func (ep eventProducer) Start() {

}