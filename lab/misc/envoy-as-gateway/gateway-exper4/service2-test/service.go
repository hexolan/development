package main;

import (
	"context"

	pb "null.hexolan.dev/dev/protogen"
)

type rpcService struct {
	pb.UnimplementedTestServiceServer
}

func newRpcService() *rpcService {
	return &rpcService{}
}

func (svc rpcService) ViewTest(ctx context.Context, req *pb.ViewTestRequest) (*pb.ViewTestReply, error) {
	return &pb.ViewTestReply{
		Message: "test",
	}, nil
}
