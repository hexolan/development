package main

import (
	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/pkg/serve"
	"github.com/hexolan/stocklet/internal/pkg/storage"
	"github.com/hexolan/stocklet/internal/pkg/metrics"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/svc/order/api"
	"github.com/hexolan/stocklet/internal/svc/order/controller"
)

func loadConfig() *order.ServiceConfig {
	// load the core service configuration
	cfg, err := order.NewServiceConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// configure the logger
	metrics.ConfigureLogger()

	// configure OTEL
	metrics.InitTracerProvider(
		cfg.Shared.Otel,
		"order",
	)

	return cfg
}

func usePostgresController(cfg *order.ServiceConfig, evtC order.EventController) (order.StorageController, *pgxpool.Pool) {
	// load the postgres configuration
	if err := cfg.Postgres.Load(); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// open a Postgres connection
	pCl, err := storage.NewPostgresConn(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// create the data controller
	dbC := controller.NewPostgresController(pCl, evtC)
	return dbC, pCl
}

func useKafkaController(cfg *order.ServiceConfig) (order.EventController, *kgo.Client) {
	// open a Kafka connection
	kCl, err := messaging.NewKafkaConn(
		cfg.Kafka,
		kgo.ConsumerGroup("order-service"),
		
		// todo: exper with REGEX consumption
		kgo.ConsumeRegex(),
		kgo.ConsumeTopics(
			messaging.Order_PlaceOrder_Catchall,
		),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// ensure the required Kafka topics exist
	err = messaging.EnsureKafkaTopics(
		kCl,

		messaging.Order_State_Created_Topic,
		messaging.Order_State_Updated_Topic,
		messaging.Order_State_Deleted_Topic,

		messaging.Order_PlaceOrder_Order_Topic,
		messaging.Order_PlaceOrder_Payment_Topic,
		messaging.Order_PlaceOrder_Shipping_Topic,
		messaging.Order_PlaceOrder_Warehouse_Topic,
	)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	// create the event controller
	evtC := controller.NewKafkaController(kCl)
	return evtC, kCl
}

func main() {
	cfg := loadConfig()

	// Create the controllers
	evtC, kCl := useKafkaController(cfg)
	defer kCl.Close()
	
	strC, pCl := usePostgresController(cfg, evtC)
	defer pCl.Close()

	// Create the service
	svc := order.NewOrderService(cfg, strC, evtC)
	
	// Attach the API interfaces to the service
	grpcSvr := api.AttachSvcToGrpc(cfg, svc)
	gwMux := api.AttachSvcToGateway(cfg, svc)
	// todo: api.AttachSvcToConsumer(cfg, svc)

	// Serve the API interfaces
	go serve.Gateway(gwMux)
	serve.Grpc(grpcSvr)
	// todo: also starting the consumer after gRPC started - serve.Consume(consumer) or something
}
