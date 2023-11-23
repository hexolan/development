package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/hexolan/experimentation/libs/go/grpc-gateway/app"
	"github.com/hexolan/experimentation/libs/go/grpc-gateway/app/helloworld"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	rpc_svr := app.NewRPCServer()

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, rpc_svr)

	if err := s.Serve(lis); err != nil {
		panic("err")
	}
}