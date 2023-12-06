package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

func NewHttpGateway() *runtime.ServeMux {
	// todo: move into generic - miscutil.NewBaseGatewayMux or something
	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:" + serve.RpcPort, opts)
	if err != nil {
		log.Panic().Err(err).Msg("failed to register gRPC to gateway server")
	}

	return mux
}