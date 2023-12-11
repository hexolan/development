package serveutil

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func NewGrpcServer() *grpc.Server {
	//svr := grpc.NewServer()
	// otel grpc server handler
	svr := grpc.NewServer(
		grpc.StatsHandler(
			otelgrpc.NewServerHandler(
				otelgrpc.WithMessageEvents(
					otelgrpc.ReceivedEvents,
					otelgrpc.SentEvents,
				),
			),
		),
	)

	reflection.Register(svr)

	return svr
}

func ServeGrpcServer(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Str("port", "9090").Msg("failed to listen on RPC port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}
