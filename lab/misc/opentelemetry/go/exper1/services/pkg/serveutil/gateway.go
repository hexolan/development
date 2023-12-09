package serveutil

import (
	"net/http"

	// "github.com/rs/zerolog/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	// "go.opentelemetry.io/otel/propagation"
)

func gatewayHeaderMatcher(key string) (string, bool) {
	// log.Info().Str("header-key", key).Msg("header matcher")
	/*
	without using an auth header:
		exper1-service1-1        | {"level":"info","header-key":"User-Agent","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Accept-Language","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"X-Envoy-Expected-Rq-Timeout-Ms","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Cache-Control","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Ch-Ua-Platform","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Accept","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Accept-Encoding","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Ch-Ua","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Fetch-Mode","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Fetch-User","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"X-Request-Id","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Traceparent","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Fetch-Site","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Ch-Ua-Mobile","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Upgrade-Insecure-Requests","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Sec-Fetch-Dest","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"X-Forwarded-Proto","time":"2023-12-09T00:04:34Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Tracestate","time":"2023-12-09T00:04:34Z","message":"header matcher"}
	*/

	/*
	using an auth header,
	envoy config is set to forward payload with "x-jwt-payload",
	this is checking if auto capitalization is taking place (as docs say they set x-request-id lowercase and is being recieved uppercase by this matcher):
		exper1-service1-1        | {"level":"info","header-key":"X-Jwt-Payload","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Traceparent","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Tracestate","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"User-Agent","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Authorization","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"X-Request-Id","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Test-Header","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"Accept","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"X-Forwarded-Proto","time":"2023-12-09T00:19:46Z","message":"header matcher"}
		exper1-service1-1        | {"level":"info","header-key":"X-Envoy-Expected-Rq-Timeout-Ms","time":"2023-12-09T00:19:46Z","message":"header matcher"}

	is indeed being capitalised when checked here.

	here is the passed through values:
		{"level":"info","meta":{":authority":["localhost:9090"],"authorization":["eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiaWF0IjoxNzAyMDgwODQyLCJleHAiOjE3MDIwODQ0NDJ9.RT6dZCIYi52U1zmXXZUTJVpt-Yn5jiseWUxKAr-90ZITMztbujnlwsxVah51gN-YaGgYE61FmBx1nesCQNFtVg"],"content-type":["application/grpc"],"grpc-accept-encoding":["gzip"],"misctest-accept":["*\/*"],"misctest-test-header":["abc"],"misctest-traceparent":["00-f1dc41a56f7319fd26581886c19d1a99-25877c4556848fbc-01"],"misctest-tracestate":[""],"misctest-user-agent":["Insomnia/2023.5.7"],"misctest-x-envoy-expected-rq-timeout-ms":["15000"],"misctest-x-forwarded-proto":["http"],"traceparent":["00-a27d413bd0a046be8e0a7d9944170870-e4ec5fab9467f799-01"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.30.0.4"],"x-forwarded-host":["localhost"],"x-jwt-payload":["eyJzdWIiOiIxIiwiaWF0IjoxNzAyMDgwODQyLCJleHAiOjE3MDIwODQ0NDJ9"],"x-request-id":["8c1848b7-0655-93be-aa1d-9b23eb0c7dab"]},"time":"2023-12-09T00:36:35Z","message":"metadata extract"}
	
	from that:
	x-request-id (fully passed thru HTTP header): 8c1848b7-0655-93be-aa1d-9b23eb0c7dab
	misctest-traceparent (passed thru HTTP header): 00-f1dc41a56f7319fd26581886c19d1a99-25877c4556848fbc-01
	traceparent (generated by OTEL Golang SDK?): 	00-a27d413bd0a046be8e0a7d9944170870-e4ec5fab9467f799-01
		> propegated downstream when recieved by gRPC server (from gateway client request)
	
	testing what happens when misctest-traceparent (traceparent http header) is passed thru.
	hopefully it propegates the trace.
	*/

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

// https://github.com/iamrajiv/opentelemetry-grpc-gateway-boilerplate
func NewGrpcGateway() (*runtime.ServeMux, []grpc.DialOption) {
	// gateway mux
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gatewayHeaderMatcher),
	)
	
	// base gateway options

	// somewhat related to trying to pass through (propegate)
	// the generated envoy request-ids
	// https://github.com/open-telemetry/opentelemetry-specification/issues/2912

	opts := []grpc.DialOption{
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	
	return mux, opts
}

func ServeGrpcGateway(mux *runtime.ServeMux) error {
	return http.ListenAndServe("0.0.0.0:90", mux)
}