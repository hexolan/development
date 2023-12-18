package api

import (
	"google.golang.org/grpc"

	"github.com/hexolan/stocklet/internal/svc/auth"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

func AttachSvcToGrpc(cfg *auth.ServiceConfig, svc *auth.AuthService) *grpc.Server {
	svr := serve.NewGrpcServeBase(&cfg.Shared)
	pb.RegisterAuthServiceServer(svr, svc)
	return svr
}