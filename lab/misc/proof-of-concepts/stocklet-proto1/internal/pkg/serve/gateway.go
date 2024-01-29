package serve

import (
	"context"
	"net/http"
	
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	
	"github.com/hexolan/stocklet/internal/pkg/config"
)

// todo: remove as not needed anymore (just using default for now)
// overriding gRPC status has prevented leaking of internal error information
func withGatewayErrorHandler() runtime.ErrorHandlerFunc {
	return func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, req, err)
	}
}

// todo: improve - http logger for debug
func withGatewayLogger(h http.Handler) http.Handler {
    return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
			log.Info().Str("path", r.URL.Path).Str("reqURI", r.RequestURI).Str("remoteAddr", r.RemoteAddr).Msg("")
		},
	)
}

func NewGatewayServeBase(cfg *config.SharedConfig) (*runtime.ServeMux, []grpc.DialOption) {
	// create the base runtime ServeMux 
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(
			withGatewayErrorHandler(),
		),
	)

	// attach open telemetry instrumentation through the gRPC client options
	clientOpts := []grpc.DialOption{
		grpc.WithStatsHandler(
			otelgrpc.NewClientHandler(),
		),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	return mux, clientOpts
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
		Handler: withGatewayLogger(handler),
	}

	// serve
	return svr.ListenAndServe()
}