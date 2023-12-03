package main

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// Common Gateway Mux
func newGatewayMux() *runtime.ServeMux {
	// add HTTP headers as metadata to gRPC gateway requests
	// DefaultHeaderMatcher is used by default - so no need for muxOpt1
	// muxOpt1 := runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher)
	// with solely muxOpt1 enabled:
	// {"level":"info","meta":{":authority":["localhost:9090"],"content-type":["application/grpc"],"grpcgateway-accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"],"grpcgateway-accept-language":["en-GB,en-US;q=0.9,en;q=0.8"],"grpcgateway-user-agent":["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.19.0.3"],"x-forwarded-host":["localhost"]},"time":"2023-12-02T23:43:45Z","message":"metadata extract"}
	// with auth header:
	// {"level":"info","meta":{":authority":["localhost:9090"],"authorization":["test"],"content-type":["application/grpc"],"gateway":["true"],"grpcgateway-accept":["*/*"],"grpcgateway-authorization":["test"],"grpcgateway-user-agent":["Insomnia/2023.5.7"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.19.0.3"],"x-forwarded-host":["localhost"]},"time":"2023-12-03T15:42:23Z","message":"metadata extract"}

	// ADD SPECIAL METADATA TO MUX
	muxOpt2 := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		return metadata.MD{
			"gateway": {"true"},
		}
	})
	// with muxOpt1 enabled and muxOpt2:
	// {"level":"info","meta":{":authority":["localhost:9090"],"content-type":["application/grpc"],"gateway":["true"],"grpcgateway-accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"],"grpcgateway-accept-language":["en-GB,en-US;q=0.9,en;q=0.8"],"grpcgateway-cache-control":["max-age=0"],"grpcgateway-user-agent":["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.19.0.3"],"x-forwarded-host":["localhost"]},"time":"2023-12-03T00:15:34Z","message":"metadata extract"}
	// HIGHLIGHTING: "gateway":["true"]

	// muxOpt3 - custom specific header matcher
	muxOpt3 := runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		switch key {
		case "Authorization":
			// return "authorization", true
			// authorization header is automatically forwaded
			// maybe change key here? - for now just work with the default mapping
			// - map no other default HTTP headers
			return "authorization", false
		default:
			// return key, false
			// for testing purposes:
			return "testing-" + key, true
		}
	})
	// with muxOpt3 and muxOpt2:
	// {"level":"info","meta":{":authority":["localhost:9090"],"authorization":["test","test"],"content-type":["application/grpc"],"gateway":["true"],"user-agent":["grpc-go/1.59.0"],"x-forwarded-for":["172.19.0.3"],"x-forwarded-host":["localhost"]},"time":"2023-12-03T15:47:55Z","message":"metadata extract"}
	// forwards only the authorization header from original HTTP request. although forwaded twice?
	// authorization header is always mapped to gRPC metadata
	// view: https://github.com/grpc-ecosystem/grpc-gateway/blob/bca0b834db86774f60c31db60b0ad9a79d23f3f8/README.md?plain=1#L594
	// instead set to false or set custom?

	// create serve mux
	// return runtime.NewServeMux(muxOpt1, muxOpt2)
	// return runtime.NewServeMux(muxOpt2)
	return runtime.NewServeMux(muxOpt2, muxOpt3)
}

// "What is the best way to know if a request came via HTTP?"
// https://github.com/grpc-ecosystem/grpc-gateway/issues/2783
//
// "Detecting request from gateway"
// https://github.com/grpc-ecosystem/grpc-gateway/issues/2159
//
// In short - adding an interceptor to the gateway gRPC client is the way
// to go about this.
//
// Functions as intended:
//
// > curl http://localhost/v1/foo/hello
// {"message":"from gateway true"}
//
// > grpcurl -plaintext localhost:9090 service.TestService/SayHello
// {"message":"from gateway false"}
//
func reqFromGateway(md *metadata.MD) bool {
	if len(md.Get("gateway")) == 1 {
		return true
	}

	return false
}

func reqIsInternal(md *metadata.MD) bool {
	return !reqFromGateway(md)
}

// JWT tokens (if provided) will be validated at the API gateway level (by Envoy - or when using Kubernetes; Istio)
// This is for handling afterwards, after being passed through gRPC-gateway (to ensure ability to interact/check within the services)
//
// todo: parsing authorization headers
// maybe using a common user header?
func reqIsAuthenticated(md *metadata.MD) bool {
	authorization := md.Get("authorization")
	if len(authorization) == 1 {
		return true
	}

	return false
}