package serve

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewGrpcBase() *grpc.Server {
	svr := grpc.NewServer()
	reflection.Register(svr)
	return svr
}

func NewGrpcGatewayBase() (*runtime.ServeMux, []grpc.DialOption) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	return mux, opts
}