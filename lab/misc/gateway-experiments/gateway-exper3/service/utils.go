package main

// JWT tokens (if provided) will be validated at the API gateway level (by Envoy - or when using Kubernetes; Istio)
// This is for handling afterwards, after being passed through gRPC-gateway (to ensure ability to interact/check within the services)

// todo:
func reqIsInternal() {
	// check if from gRPC or if gateway
}

func reqIsAuthenticated() {

}