package main

import (
	"null.hexolan.dev/dev/pkg/rpc"
	"null.hexolan.dev/dev/pkg/serve"
	"null.hexolan.dev/dev/pkg/gateway"
	"null.hexolan.dev/dev/pkg/database"
)

func main() {
	db := database.NewDatabaseInterface("postgresql://postgres:postgres@test-service-postgres:5432/postgres?sslmode=disable")

	svr := rpc.NewGrpcServer(db)
	gw := gateway.NewGrpcGateway()

	go serve.GrpcGateway(gw)
	serve.GrpcServer(svr)
}