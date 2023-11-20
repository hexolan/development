package main

import (
	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/development/experimentation/gateway-exper1/service"
	"github.com/hexolan/development/experimentation/gateway-exper1/service/api"
	"github.com/hexolan/development/experimentation/gateway-exper1/service/controller"
	"github.com/hexolan/development/experimentation/gateway-exper1/utils/serve"
	"github.com/hexolan/development/experimentation/gateway-exper1/utils/config"
	"github.com/hexolan/development/experimentation/gateway-exper1/utils/clients"
)


func usePostgresController(pgCfg *config.PostgresConfig, evtC service.EventController) (service.StorageController, *pgxpool.Pool) {
	// open a Postgres connection
	pCl, err := clients.NewPostgresConn(pgCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// create the data controller
	dbC := controller.NewPostgresController(pCl, evtC)
	return dbC, pCl
}

func useKafkaController(kCfg *config.KafkaConfig) (service.EventController, *kgo.Client) {
	// open a Kafka connection
	kCl, err := clients.NewKafkaConn(
		kCfg,
		kgo.ConsumerGroup("test-service"),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// create the event controller
	evtC := controller.NewKafkaController(kCl)
	return evtC, kCl
}

func main() {
	// Create the controllers
	evtC, kcl := useKafkaController(
		&config.KafkaConfig{
			Brokers: []string{"kafka:19092"},
		},
	)
	defer kcl.Close()
	
	strC, db := usePostgresController(
		&config.PostgresConfig{
			Username: "postgres",
			Password: "postgres",
			Host: "test-service-postgres:5432",
			Database: "postgres",
		},
		evtC,
	)
	defer db.Close()

	// Create the service
	svc := service.NewTestService(evtC, strC)
	
	// Attach the API interfaces to the service
	grpcSvr := api.NewGrpcServer(svc)
	gatewayMux := api.NewHttpGateway()

	// Serve the API interfaces
	go serve.GrpcServer(grpcSvr)
	serve.HttpGateway(gatewayMux)
}
