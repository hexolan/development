package serveutil

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func gatewayHeaderMatcher(key string) (string, bool) {
	switch key {
	case "Authorization":
		// authorization header is automatically forwarded
		return "authorization", false
	case "X-Jwt-Payload":
		// custom envoy header payload after validation
		// base64 encoded (and validated) jwt payload
		return "x-jwt-payload", true
	case "X-Request-Id":
		// custom envoy header payload after validation
		// base64 encoded (and validated) jwt payload
		return "x-request-id", true
	case "Traceparent":
		// opentelemetry trace parent
		return "traceparent", true
	case "Tracestate":
		// opentelemetry
		return "tracestate", true
	default:
		// for testing purposes:
		return "misctest-" + key, true
		// return key, false
	}
}

func NewGrpcGateway() (*runtime.ServeMux, []grpc.DialOption) {
	// gateway mux
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gatewayHeaderMatcher),
	)

	// base gateway options
	// insecure transport and OTEL grpc client tracing
	opts := []grpc.DialOption{
		grpc.WithStatsHandler(
			otelgrpc.NewClientHandler(
				otelgrpc.WithMessageEvents(
					otelgrpc.ReceivedEvents,
					otelgrpc.SentEvents,
				),
			),
		),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	
	return mux, opts
}

func ServeGrpcGateway(mux *runtime.ServeMux) error {
	// https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/tracing/
	// return http.ListenAndServe("0.0.0.0:90", mux)

	// attach OTEL handler
	return http.ListenAndServe(
		"0.0.0.0:90",
		otelhttp.NewHandler(
			mux,
			"grpc-gateway",
			/*
			otelhttp.WithMessageEvents(
				otelhttp.ReadEvents,
				otelhttp.WriteEvents,
			),
			*/
			otelhttp.WithSpanNameFormatter(
				func (operation string, r *http.Request) string {
					return operation + ": " + r.RequestURI
				},
			),
		),
	)
}