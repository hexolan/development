package api

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func NewHttpGateway(rpcSvc pb.OrderServiceServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	// https://github.com/grpc-ecosystem/grpc-gateway/issues/1458
	// note: no gRPC handler support when using this method
	err := pb.RegisterOrderServiceHandlerServer(
		ctx,
		mux,
		rpcSvc,
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to register gRPC to gateway server")
	}

	return http.ListenAndServe(":8080", mux)
}
