package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "null.hexolan.dev/dev/protogen/testpb"
)

type rpcService struct {
	db DatabaseInterface

	pb.UnimplementedTestServiceServer
}

func newGrpcServer(db DatabaseInterface) *grpc.Server {
	svr := newGrpcBase()

	svc := rpcService{db: db}
	pb.RegisterTestServiceServer(svr, svc)

	return svr
}

func newGrpcGateway() *runtime.ServeMux {
	mux, opts := newGrpcGatewayBase()

	ctx := context.Background()
	err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Panic().Err(err).Msg("gateway svc registration")
	}

	return mux
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