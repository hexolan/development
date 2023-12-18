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
		log.Fatal().Err(err).Msg("failed to register svc to gateway server")
	}

	return mux
}