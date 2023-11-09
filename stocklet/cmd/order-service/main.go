package main

import (
	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/pkg/logging"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/svc/order/api"
	"github.com/hexolan/stocklet/internal/svc/order/controller"
)

func loadConfig() *order.ServiceConfig {
	// load the service configuration
	cfg, err := order.NewServiceConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// configure the logger
	logging.ConfigureLogger()

	return cfg
}

func usePostgresController(cfg *order.ServiceConfig) (order.DataController, *pgxpool.Pool) {
	// open a Postgres connection
	db, err := database.NewPostgresConn(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// create the data controller
	dbC := controller.NewPostgresController(db)
	return dbC, db
}

func useKafkaController(cfg *order.ServiceConfig) (order.EventController, *kgo.Client) {
	// open a Kafka connection
	kcl, err := messaging.NewKafkaConn(cfg.Kafka, kgo.ConsumerGroup("order-service"), kgo.ConsumeTopics("orders"))
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// ensure the required Kafka topics exist
	err = messaging.EnsureKafkaTopics(kcl, "orders", "orders2")
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	// create the event controller
	evtC := controller.NewKafkaController(kcl)
	return evtC, kcl
}

func main() {
	cfg := loadConfig()

	// todo: clean up variable names for controllers and client conns

	// Create the controllers
	evtC, kcl := useKafkaController(cfg)
	defer kcl.Close()
	
	dbC, db := usePostgresController(cfg)
	defer db.Close()

	// Create the service
	svc := order.NewOrderService(dbC, evtC)

	// Attach the API interfaces to the service
	go api.NewKafkaConsumer(svc, kcl)
	
	grpcSvr := api.NewGrpcServer(svc)
	go api.ServeGrpcServer(grpcSvr)
	api.NewHttpGateway(svc)  // todo: change for client conn - use grpcSvr instead of svc - kafka consumer can maintain
}
