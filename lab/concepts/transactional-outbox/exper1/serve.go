package main

import (
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func serveGrpcServer(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Str("port", "9090").Msg("listening on rpc port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("serving grpc server")
	}
}

func serveGrpcGateway(mux *runtime.ServeMux) error {
	lis, err := net.Listen("tcp", "0.0.0.0:90")
	if err != nil {
		log.Panic().Err(err).Str("port", "90").Msg("listening on gateway port")
	}

	err = http.Serve(lis, mux)
	return err
}