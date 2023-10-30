package service

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/svc/order/api"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func newRPCServer(rpcSvc pb.OrderServiceServer) {
	svr := grpc.NewServer()

	healthSvc := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, healthSvc)

	pb.RegisterOrderServiceServer(svr, rpcSvc)

	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Msg("failed to listen on RPC port (:9090)")
	}

	err = svr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}

func NewOrderService(db *pgxpool.Pool, kcl *kgo.Client) {
	rpcSvc := api.NewRPCServicer(db, kcl)

	// Create the RPC server
	go newRPCServer(rpcSvc)

	// Create the HTTP grpc-gateway interface
	go api.NewHttpGateway(rpcSvc)
	
	// Begin consuming events from Kafka
	api.NewEventConsumer(kcl)
}