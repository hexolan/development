// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/svc/auth"
	"github.com/hexolan/stocklet/internal/svc/auth/api"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	"github.com/hexolan/stocklet/internal/pkg/metrics"
)

func loadConfig() *auth.ServiceConfig {
	// Load the main service configuration
	cfg, err := auth.NewServiceConfig()
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	// Configure metrics (logging and OTEL)
	metrics.ConfigureLogger()
	metrics.InitTracerProvider(&cfg.Shared.Otel, "auth")

	return cfg
}

func usePostgresController(cfg *auth.ServiceConfig) (*auth.StorageController, error) {
	// load the postgres configuration
	if err := cfg.Postgres.Load(); err != nil {
		log.Panic().Err(err).Msg("")
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
	grpcSvr := api.AttachSvcToGrpc(cfg, svc)
	gwMux := api.AttachSvcToGateway(cfg, svc)

	// Serve the API interfaces
	go serve.Gateway(gwMux)
	serve.Grpc(grpcSvr)
}

