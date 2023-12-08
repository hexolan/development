package main

import (
	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/svc/auth"
	"github.com/hexolan/stocklet/internal/svc/auth/api"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	"github.com/hexolan/stocklet/internal/pkg/logging"
)

func loadConfig() *auth.ServiceConfig {
	// load the main service configuration
	cfg, err := auth.NewServiceConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// configure the logger
	logging.ConfigureLogger()

	return cfg
}

func usePostgresController(cfg *auth.ServiceConfig) (*auth.StorageController, error) {
	// load the postgres configuration
	if err := cfg.Postgres.Load(); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// todo:
	// instead of error return client

	// postgresController{}, client
	return nil, nil
}

func main() {
	cfg := loadConfig()

	// Create the controllers
	strC, _ := usePostgresController(cfg)

	// Create the service
	svc := auth.NewAuthService(cfg, strC)
	
	// Attach the API interfaces to the service
	grpcSvr := api.NewGrpcServer(cfg, svc)
	gatewayMux := api.NewHttpGateway()

	// Serve the API interfaces
	go serve.GrpcServer(grpcSvr)
	serve.HttpGateway(gatewayMux)
}

