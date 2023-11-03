package api

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func NewGrpcServer(svc pb.OrderServiceServer) *grpc.Server {
	svr := grpc.NewServer()

	healthSvc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, healthSvc)

	pb.RegisterOrderServiceServer(svr, svc)

	return svr
}

// todo: move to pkg method?
func ServeGrpcServer(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Msg("failed to listen on RPC port (:9090)")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}