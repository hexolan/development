package main

import (
	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/logging"
	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/app/order"
	"github.com/hexolan/stocklet/internal/app/order/api/http"
	"github.com/hexolan/stocklet/internal/app/order/api/event"
)

func main() {
	// Load the required configurations
	cfg, err := order.NewServiceConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	logging.ConfigureLogger()

	// Open a database connection pool
	db, err := database.NewPostgresConn(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Create the service repository
	prod := order.NewEventProd(cfg.Kafka)
	svc := order.NewOrderService(db, prod)
	
	// Start the HTTP and Event APIs
	go http.NewHttpAPI(svc)
	event.NewEventAPI(svc)
}
