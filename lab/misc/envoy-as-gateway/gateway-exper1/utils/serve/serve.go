package serve

import (
	"net"
	"net/http"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func GrpcServer(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Str("port", "9090").Msg("failed to listen on RPC port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}

func HttpGateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:90", mux)
}