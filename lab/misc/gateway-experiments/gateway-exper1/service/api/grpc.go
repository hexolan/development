package api

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"null.hexolan.dev/exp/service"
	pb "null.hexolan.dev/exp/protogen/testingv1"
)

func NewGrpcServer(svc service.Service) *grpc.Server {
	svr := grpc.NewServer()

	hSvc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, hSvc)

	// for dev usage (with grpcui)
	reflection.Register(svr)

	pb.RegisterTestServiceServer(svr, svc)

	return svr
}