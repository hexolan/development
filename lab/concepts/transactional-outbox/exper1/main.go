package main

func main() {
	svr := newGrpcServer()
	gw := newGrpcGateway()

	go serveGrpcGateway(gw)
	serveGrpcServer(svr)
}