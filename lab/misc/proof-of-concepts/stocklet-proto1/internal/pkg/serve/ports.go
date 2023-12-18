package serve

const (
	grpcPort string = "9090"
	gatewayPort string = "90"
)

func AddrToGrpc(host string) string {
	return host + ":" + grpcPort
}

func AddrToGateway(host string) string {
	return host + ":" + gatewayPort
}