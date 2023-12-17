package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func AttachSvcToGateway(cfg *order.ServiceConfig, svc *order.OrderService) *runtime.ServeMux {
	mux, clientOpts := serve.NewGatewayServeBase(cfg.Shared)

	ctx := context.Background()
	err := pb.RegisterOrderServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:" + serve.GatewayPort,
		clientOpts,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to register svc to gateway server")
	}

	return mux
}