package api

import (
	"google.golang.org/grpc"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func NewGrpcServer(svc order.OrderService) *grpc.Server {
	svr := grpc.NewServer()
	serve.AttachHealthService(svr)

	pb.RegisterOrderServiceServer(svr, svc)

	return svr
}