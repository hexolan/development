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

package serve

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

// todo: logger for GRPC
func NewGrpcServeBase(cfg *config.SharedConfig) *grpc.Server {
	// attach OTEL metrics middleware
	svr := grpc.NewServer(
		grpc.StatsHandler(
			otelgrpc.NewServerHandler(),
		),
	)

	// attach the health service
	svc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, svc)
	
	// enable reflection in dev mode
	// this is to make testing easier with tools like grpcurl and grpcui
	if cfg.DevMode {
		reflection.Register(svr)
	}

	return svr
}

func Grpc(svr *grpc.Server) {
	lis, err := net.Listen("tcp", AddrToGrpc("0.0.0.0"))
	if err != nil {
		log.Panic().Err(err).Str("port", grpcPort).Msg("failed to listen on gRPC port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}