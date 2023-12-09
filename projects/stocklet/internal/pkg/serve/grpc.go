package serve

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

func NewGrpcServeBase(cfg *config.SharedConfig) *grpc.Server {
	// todo: attaching OTEL metrics middleware
	svr := grpc.NewServer()

	// attach the health service
	svc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, svc)
	
	// enable reflection in dev mode
	// this is to make testing easier with tools like grpcurl and grpcui
	if cfg.DevMode {
		reflection.Register(svr)
	}

	return svr
}

func Grpc(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:" + GrpcPort)
	if err != nil {
		log.Panic().Err(err).Str("port", GrpcPort).Msg("failed to listen on gRPC port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}