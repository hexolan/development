package main;

import (
	"net"
	"net/http"
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "null.hexolan.dev/dev/protogen"
)

func newGrpcServer(svc pb.TestServiceServer) *grpc.Server {
	svr := grpc.NewServer()
	reflection.Register(svr)
	pb.RegisterTestServiceServer(svr, svc)

	return svr
}

func serveGrpcServer(svr *grpc.Server) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Str("port", "9090").Msg("failed to listen on RPC port")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}

func newHttpGateway() *runtime.ServeMux {
	ctx := context.Background()

	// == MUX OPTS ==
	// add HTTP headers as metadata to gRPC gateway requests
	muxOpt1 := runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher)
	// with solely muxOpt1 enabled:
	// {"level":"info","meta":{":authority":["localhost:9090"],"content-type":["application/grpc"],"grpcgateway-accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"],"grpcgateway-accept-language":["en-GB,en-US;q=0.9,en;q=0.8"],"grpcgateway-user-agent":["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.19.0.3"],"x-forwarded-host":["localhost"]},"time":"2023-12-02T23:43:45Z","message":"metadata extract"}

	// ADD SPECIAL METADATA TO MUX
	muxOpt2 := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		return metadata.MD{
			"gateway": {"true"},
		}
	})
	// with muxOpt1 enabled and muxOpt2:
	// {"level":"info","meta":{":authority":["localhost:9090"],"content-type":["application/grpc"],"gateway":["true"],"grpcgateway-accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"],"grpcgateway-accept-language":["en-GB,en-US;q=0.9,en;q=0.8"],"grpcgateway-cache-control":["max-age=0"],"grpcgateway-user-agent":["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.19.0.3"],"x-forwarded-host":["localhost"]},"time":"2023-12-03T00:15:34Z","message":"metadata extract"}
	// HIGHLIGHTING: "gateway":["true"]
	// == END OF MUX OPTS ==

	// create serve mux
	mux := runtime.NewServeMux(muxOpt1, muxOpt2)

	// register service to mux
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterTestServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Panic().Err(err).Msg("failed to register gRPC to gateway server")
	}

	return mux
}

func serveHttpGateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:90", mux)
}

func main() {
	rpcSvc := newRpcService()
	rpcSvr := newGrpcServer(rpcSvc)
	go serveGrpcServer(rpcSvr)

	gwayMux := newHttpGateway()
	serveHttpGateway(gwayMux)
}