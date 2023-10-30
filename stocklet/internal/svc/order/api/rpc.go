package api

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	"github.com/hexolan/stocklet/internal/svc/order"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type rpcServicer struct {
	repo order.DatabaseRepo
	prod order.EventProducer

	pb.UnimplementedOrderServiceServer
}

func NewRPCServicer(db *pgxpool.Pool, kcl *kgo.Client) rpcServicer {
	repo := order.NewDatabaseRepo(db)
	prod := order.NewEventProducer(kcl)

	return rpcServicer{
		repo: repo,
		prod: prod,
	}
}

func (svc rpcServicer) GetOrder(context.Context, *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc rpcServicer) GetOrders(context.Context, *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc rpcServicer) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc rpcServicer) UpdateOrder(context.Context, *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc rpcServicer) CancelOrder(context.Context, *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc rpcServicer) DeleteUserData(context.Context, *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}