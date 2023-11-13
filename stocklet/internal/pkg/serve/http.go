package serve

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func ServeHttpGateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:" + HttpPort, mux)
}