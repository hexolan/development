package serve

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

func NewGatewayServeBase(cfg *config.SharedConfig) (*runtime.ServeMux, []grpc.DialOption) {
	mux := runtime.NewServeMux()

	// todo: attach open telemetry interface
	clientOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	return mux, clientOpts
}

func Gateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:" + GatewayPort, mux)
}