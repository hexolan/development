package rpc

import (
	"context"

	"google.golang.org/grpc"

	"null.hexolan.dev/dev/pkg/serve"
	"null.hexolan.dev/dev/pkg/database"
	pb "null.hexolan.dev/dev/protogen/testpb"
)

type rpcService struct {
	db database.DatabaseInterface

	pb.UnimplementedTestServiceServer
}

func NewGrpcServer(db database.DatabaseInterface) *grpc.Server {
	svr := serve.NewGrpcBase()

	svc := rpcService{db: db}
	pb.RegisterTestServiceServer(svr, svc)

	return svr
}

func (svc rpcService) GetTestItem(ctx context.Context, req *pb.GetTestItemRequest) (*pb.GetTestItemResponse, error) {
	item, err := svc.db.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTestItemResponse{Data: item}, nil
}

func (svc rpcService) CreateTestItem(ctx context.Context, req *pb.CreateTestItemRequest) (*pb.CreateTestItemResponse, error) {
	item, err := svc.db.CreateItem(ctx, req.Item)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTestItemResponse{Data: item}, nil
}