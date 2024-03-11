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

package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/svc/auth"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

func AttachSvcToGateway(cfg *auth.ServiceConfig, svc *auth.AuthService) *runtime.ServeMux {
	mux, clientOpts := serve.NewGatewayServeBase(&cfg.Shared)

	ctx := context.Background()
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, serve.AddrToGrpc("localhost"), clientOpts)
	if err != nil {
		log.Panic().Err(err).Msg("failed to register svc to gateway server")
	}

	return mux
}