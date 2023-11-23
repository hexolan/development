package app

import (
	"context"

	"github.com/hexolan/experimentation/libs/go/grpc-gateway/app/helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func NewRPCServer() helloworld.GreeterServer {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}