package dispatch

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type EventProducer interface {
	DispatchCreatedEvent() error
	DispatchUpdatedEvent() error
	DispatchDeletedEvent() error
}

type eventProducer struct {}

func NewEventProducer(conf config.KafkaConfig) EventProducer {
	return eventProducer{}
}

func (prod eventProducer) DispatchCreatedEvent() error {
	return nil // todo:
}

func (prod eventProducer) DispatchUpdatedEvent() error {
	return nil // todo:
}

func (prod eventProducer) DispatchDeletedEvent() error {
	return nil // todo:
}