package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "github.com/hexolan/development/experimentation/gateway-exper1/protogen/testingv1"
)

func NewHttpGateway() *runtime.ServeMux {
	ctx := context.Background()
	
	// todo: pass thru cancel
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Panic().Err(err).Msg("failed to register gRPC to gateway server")
	}

	return mux
}