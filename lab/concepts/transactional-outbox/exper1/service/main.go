package main

func main() {
	db := NewDatabaseInterface()

	svr := newGrpcServer(db)
	gw := newGrpcGateway()

	// go db.StartProducer()
	go serveGrpcGateway(gw)
	serveGrpcServer(svr)
}