package main

import (
	"google.golang.org/grpc/metadata"
)

// "What is the best way to know if a request came via HTTP?"
// https://github.com/grpc-ecosystem/grpc-gateway/issues/2783
//
// "Detecting request from gateway"
// https://github.com/grpc-ecosystem/grpc-gateway/issues/2159
//
// In short - adding an interceptor to the gateway gRPC client is the way
// to go about this.
//
// Functions as intended
/*
D:\Git\GitHub\Hexolan>curl http://localhost/v1/foo/hello
  {"message":"from gateway true"}
D:\Git\GitHub\Hexolan>grpcurl -plaintext localhost:9090 service.TestService/SayHello
  {"message":"from gateway false"}
*/
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
func reqIsAuthenticated() {

}