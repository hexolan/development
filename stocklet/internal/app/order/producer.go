package order

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type EventProducer interface {
	
}

type eventProd struct {}

func NewEventProd(conf config.KafkaConfig) EventProducer {
	return eventProd{}
}