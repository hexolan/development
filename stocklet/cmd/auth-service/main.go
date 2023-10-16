package main

import (
	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/app/auth"
	"github.com/hexolan/stocklet/internal/app/auth/http"
	"github.com/hexolan/stocklet/internal/app/auth/events"
)

func main() {
	cfg, err := auth.NewServiceConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgresConn(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	/*
	// todo: open kafka conn (maybe using goka or sarama directly?)
	kafka, err := messaging.NewKafkaConnection()
	if err != nil {
		panic(err)
	}
	*/

	prod := auth.NewEventProd(cfg.Kafka)
	svc := auth.NewAuthService(db, prod)
	
	http.NewRestAPI(svc)
	go events.NewEventConsumer(svc)
}
