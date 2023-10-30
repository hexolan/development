package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func NewHttpGateway(rpcSvc pb.OrderServiceServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()


	// https://github.com/grpc-ecosystem/grpc-gateway/issues/1458
	// note: no gRPC handler support when using this method
	return pb.RegisterOrderServiceHandlerServer(
		ctx,
		mux,
		rpcSvc,
	)
}