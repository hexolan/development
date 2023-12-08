package api

import (
	"google.golang.org/grpc"

	"github.com/hexolan/stocklet/internal/svc/auth"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

func NewGrpcServer(cfg *auth.ServiceConfig, svc *auth.AuthService) *grpc.Server {
	// todo: move into generic?
	// miscutil.NewBaseGrpcServer or something?
	svr := grpc.NewServer()
	serve.AttachGrpcUtils(svr, cfg.Shared.DevMode)

	pb.RegisterAuthServiceServer(svr, svc)

	return svr
}