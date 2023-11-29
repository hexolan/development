package main;

import (
	"context"

	pb "null.hexolan.dev/dev/protogen"
)

type rpcService struct {
	pb.UnimplementedTestServiceServer
}

func newRpcService() pb.TestServiceServer {
	return rpcService{}
}

func (svc rpcService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "service says hello",
	}, nil
}
