package main

import (
	"github.com/hexolan/stocklet/internal/pkg/database"
	"github.com/hexolan/stocklet/internal/app/order"
	"github.com/hexolan/stocklet/internal/app/order/http"
	"github.com/hexolan/stocklet/internal/app/order/events"
)

func main() {
	cfg, err := order.NewServiceConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.NewPostgresConn(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// todo: open kafka conn

	prod := order.NewEventProd(cfg.Kafka)
	svc := order.NewOrderService(db, prod)
	
	http.NewRestAPI(svc)
	go events.NewEventConsumer(svc)
}
