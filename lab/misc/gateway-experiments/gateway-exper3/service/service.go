package main;

import (
	"fmt"
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"

	pb "null.hexolan.dev/dev/protogen"
)

type rpcService struct {
	pb.UnimplementedTestServiceServer
}

func newRpcService() *rpcService {
	return &rpcService{}
}

func (svc rpcService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// https://www.hward.com/golang-grpc-context-client-server/
	md, _ := metadata.FromIncomingContext(ctx)
	log.Info().Any("meta", md).Msg("metadata extract")

	log.Info().Any("val", ctx.Value("Authorization")).Msg("")
	log.Info().Any("ctx", ctx).Msg("")

	fromGateway := reqFromGateway(&md)
	return &pb.HelloReply{
		Message: "from gateway " + fmt.Sprintf("%v", fromGateway),
	}, nil
}
