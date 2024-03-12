// Copyright (C) 2024 Declan Teevan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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

	// Create the storage controller
	// todo: defering client closure
	store, _ := usePostgresController(cfg)

	// Create the service (& API interfaces)
	svc := auth.NewAuthService(cfg, store)
	grpcSvr := api.PrepareGrpc(cfg, svc)
	gatewayMux := api.PrepareGateway(cfg)

	// Serve the API interfaces
	go serve.Gateway(gatewayMux)
	serve.Grpc(grpcSvr)
}

