package main

import (
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/pkg/logging"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/svc/order/service"
)

func main() {
	// Load the required configurations
	cfg, err := order.NewServiceConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	logging.ConfigureLogger()

	// Open a Postgres connection
	db, err := database.NewPostgresConn(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer db.Close()

	// Open a Kafka connection
	kcl, err := messaging.NewKafkaConn(cfg.Kafka, kgo.ConsumerGroup("order-service"), kgo.ConsumeTopics("orders"))
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer kcl.Close()

	// Ensure the required Kafka topics exist
	err = messaging.EnsureKafkaTopics(kcl, "orders", "orders2")
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	// Create the service
	service.NewOrderService(db, kcl)
}
