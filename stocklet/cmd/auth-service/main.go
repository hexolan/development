package main

import (
	"github.com/hexolan/stocklet/internal/auth"
)

func main() {
	// todo: something alike
	cfg, err := config.LoadFromEnvironment()
	if err != nil {
		
	}

	db, err := postgres.NewDatabaseConnection()
	if err != nil {

	}

	auth.NewAuthService(db)

	kafka, err := kafka.NewKafkaConnection()
}
