package serve

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	
	"github.com/hexolan/stocklet/internal/pkg/config"
)

func NewGatewayServeBase(cfg *config.SharedConfig) (*runtime.ServeMux, []grpc.DialOption) {
	mux := runtime.NewServeMux()

	// attach open telemetry instrumentation
	clientOpts := []grpc.DialOption{
		grpc.WithStatsHandler(
			otelgrpc.NewClientHandler(),
		),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	return mux, clientOpts
}

// todo: improve - http logger for debug
func withLogger(h http.Handler) http.Handler {
    return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
			log.Info().Str("path", r.URL.Path).Str("reqURI", r.RequestURI).Str("remoteAddr", r.RemoteAddr).Msg("")
		},
	)
}

func Gateway(mux *runtime.ServeMux) error {
	// create OTEL instrumentation handler
	handler := otelhttp.NewHandler(
		mux,
		"grpc-gateway",
		otelhttp.WithSpanNameFormatter(
			func (operation string, r *http.Request) string {
				return operation + ": " + r.RequestURI
			},
		),
	)

	// create gateway HTTP server
	svr := &http.Server{
		Addr: AddrToGateway("0.0.0.0"),
		// Handler: handler,
		Handler: withLogger(handler),
	}

	// serve
	return svr.ListenAndServe()
}