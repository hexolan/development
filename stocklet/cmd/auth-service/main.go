package main

import (
	"github.com/hexolan/stocklet/internal/pkg/datastore"
	"github.com/hexolan/stocklet/internal/app/auth"
)

func main() {
	// todo: something alike
	cfg, err := config.LoadFromEnvironment()
	if err != nil {

	}

	db, err := datastore.NewPostgresConnection()
	if err != nil {

	}

	auth.NewAuthService(db)

	kafka, err := kafka.NewKafkaConnection()
}
