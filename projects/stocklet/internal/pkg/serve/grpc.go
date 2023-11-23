package serve

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func GrpcServer(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:" + RpcPort)
	if err != nil {
		log.Panic().Err(err).Str("port", RpcPort).Msg("failed to listen on RPC port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}

func AttachGrpcUtils(svr *grpc.Server, devMode bool) {
	// attach the health service
	svc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, svc)
	
	// enable reflection in dev mode
	// this is to make testing easier with tools like grpcurl and grpcui
	if devMode {
		reflection.Register(svr)
	}
}