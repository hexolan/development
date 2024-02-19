package gateway

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"null.hexolan.dev/dev/pkg/serve"
	pb "null.hexolan.dev/dev/protogen/testpb"
)

func NewGrpcGateway() *runtime.ServeMux {
	mux, opts := serve.NewGrpcGatewayBase()

	ctx := context.Background()
	err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Panic().Err(err).Msg("gateway svc registration")
	}

	return mux
}