package main

import (
	"github.com/hexolan/stocklet/internal/app/auth"
	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
)

func main() {
	// todo: something alike
	cfg, err := auth.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	kafka, err := messaging.NewKafkaConnection()
	if err != nil {
		panic(err)
	}

	auth.NewAuthService(db)

}
