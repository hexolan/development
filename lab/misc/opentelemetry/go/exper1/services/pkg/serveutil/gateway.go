package serveutil

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewGrpcGateway() (*runtime.ServeMux, []grpc.DialOption) {
	// gateway mux
	mux := runtime.NewServeMux()
	
	// base gateway options
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	
	return mux, opts
}

func ServeGrpcGateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:90", mux)
}