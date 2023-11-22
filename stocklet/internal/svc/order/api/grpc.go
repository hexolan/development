package api

import (
	"google.golang.org/grpc"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func NewGrpcServer(cfg *order.ServiceConfig, svc order.OrderService) *grpc.Server {
	svr := grpc.NewServer()

	pb.RegisterOrderServiceServer(svr, svc)
	
	serve.AttachGrpcUtils(svr, cfg.DevMode)

	return svr
}