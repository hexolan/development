package main

import (
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// https://thegodev.com/transactional-outbox-pattern/
// https://medium.com/@adamszpilewicz/outbox-pattern-in-go-with-postgres-alternative-to-rabbitmq-aa203d9cc9d0
//
// https://pradeepl.com/blog/transactional-outbox-pattern/
// https://medium.com/trendyol-tech/transaction-log-tailing-with-debezium-part-1-aeb968d72220
//
// https://pkritiotis.io/outbox-pattern-in-go/
//

// specific tools:
// could use Debezium
//
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
	log.Info().Msg("producer started")
	log.Info().Msg("producer not implemented")
}

// other notes:
// potentially have this as a seperate instance?? (event producer)
// unsure on scalability of log publisher

// the event "producer" is effectively as "message relay"
// > https://pradeepl.com/blog/transactional-outbox-pattern/ (**** useful)
