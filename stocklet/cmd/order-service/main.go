package main

import (
	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/logging"
	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	"github.com/hexolan/stocklet/internal/app/order"
	"github.com/hexolan/stocklet/internal/app/order/api"
)

func main() {
	// Load the required configurations
	cfg, err := order.NewServiceConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	logging.ConfigureLogger()

	// Open a database connection
	db, err := database.NewPostgresConn(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Ensure the required Kafka topics exist
	err = messaging.EnsureKafkaStreamTopics(
		cfg.Kafka, 
		[]string{
			string(order.TopicTestThing),
		},
	)
	if err != nil {
		// todo: check if this raises err if topic already exists
		// may not always need to be fatal (only if conn cannot be established)
		log.Fatal().Err(err).Msg("")
	}

	// Create the service repositories
	dbRepo := order.NewDBRepository(db)
	evtRepo := order.NewEventRepository(dbRepo, cfg.Kafka)
	svc := order.NewServiceRepository(evtRepo)
	
	// Start the HTTP and event interfaces
	go api.NewHttpAPI(svc)
	api.NewEventAPI(svc, cfg.Kafka)
}
