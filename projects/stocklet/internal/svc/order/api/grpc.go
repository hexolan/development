package api

import (
	"google.golang.org/grpc"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func AttachSvcToGrpc(cfg *order.ServiceConfig, svc *order.OrderService) *grpc.Server {
	svr := serve.NewGrpcServeBase(cfg.Shared)
	pb.RegisterOrderServiceServer(svr, svc)
	return svr
}