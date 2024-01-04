package serve

// Port Definitions
const (
	grpcPort string = "9090"
	gatewayPort string = "90"
)

// Get an address to a gRPC server
func AddrToGrpc(host string) string {
	return host + ":" + grpcPort
}

// Get an address to a gRPC-gateway interface
func AddrToGateway(host string) string {
	return host + ":" + gatewayPort
}