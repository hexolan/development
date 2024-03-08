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
	// Load the core service configuration
	cfg, err := order.NewServiceConfig()
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	// Configure metrics (logging and OTEL)
	metrics.ConfigureLogger()
	metrics.InitTracerProvider(
		&cfg.Shared.Otel,
		"order",
	)

	return cfg
}

func usePostgresController(cfg *order.ServiceConfig) (order.StorageController, *pgxpool.Pool) {
	// load the Postgres configuration
	if err := cfg.Postgres.Load(); err != nil {
		log.Panic().Err(err).Msg("")
	}

	// open a Postgres connection
	pCl, err := storage.NewPostgresConn(&cfg.Postgres)
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	strC := controller.NewPostgresController(pCl)
	return strC, pCl
}

func useKafkaController(cfg *order.ServiceConfig) (order.EventController, *kgo.Client) {
	// load the Kafka configuration
	if err := cfg.Kafka.Load(); err != nil {
		log.Panic().Err(err).Msg("")
	}

	// open a Kafka connection
	kCl, err := messaging.NewKafkaConn(
		&cfg.Kafka,
		kgo.ConsumerGroup("order-service"),
	)
	if err != nil {
		log.Panic().Err(err).Msg("")
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
	
	strC, pCl := usePostgresController(cfg)
	defer pCl.Close()

	// Create the service
	svc := order.NewOrderService(cfg, strC, evtC)
	
	// Attach the API interfaces to the service
	grpcSvr := api.AttachSvcToGrpc(cfg, svc)
	gwMux := api.AttachSvcToGateway(cfg, svc)
	consCtrl := api.AttachSvcToConsumer(cfg, svc)

	// Serve the API interfaces
	go serve.Gateway(gwMux)
	go consCtrl.Start()  // todo: starting consumer after gRPC; though this should work for now since calling directly (goroutine/thread safety needs a look into though)
	serve.Grpc(grpcSvr)
}
