package main;

import (
	"net"
	"net/http"
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "go.hexolan.dev/dev/service1/protogen"
)

func newGrpcServer(svc pb.Experiment1ServiceServer) *grpc.Server {
	svr := grpc.NewServer()
	reflection.Register(svr)
	pb.RegisterExperiment1ServiceServer(svr, svc)

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

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterExperiment1ServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Panic().Err(err).Msg("failed to register gRPC to gateway server")
	}

	return mux
}

func serveHttpGateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:90", mux)
}

type rpcService struct {
	pb.UnimplementedExperiment1ServiceServer
}

func newRpcService() pb.Experiment1ServiceServer {
	return rpcService{}
}

func (svc rpcService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "service1 says hello",
	}, nil
}

func main() {
	rpcSvc := newRpcService()
	rpcSvr := newGrpcServer(rpcSvc)
	go serveGrpcServer(rpcSvr)

	gwayMux := newHttpGateway()
	serveHttpGateway(gwayMux)
}