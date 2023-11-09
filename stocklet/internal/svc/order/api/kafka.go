package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"

	"github.com/hexolan/stocklet/internal/svc/order"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type kafkaConsumer struct {
	svc order.OrderService
	
	kcl *kgo.Client
}

func NewKafkaConsumer(svc order.OrderService, kcl *kgo.Client) kafkaConsumer {
	return kafkaConsumer{
		svc: svc,
		kcl: kcl,
	}
}

func (c kafkaConsumer) StartConsuming() {
	ctx := context.Background()

	for {
		fetches := c.kcl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Panic().Any("kafka-errs", errs).Msg("consumer: unrecoverable kafka errors")
		}

		fetches.EachTopic(func(ft kgo.FetchTopic) {
			switch ft.Topic {
			case "topic1":
				// consumeTopic1(svc, ft)
				// todo:
				// topic schema
				// order.placeorder.*
				// order.placeorder (publ. )
				// order.placeorder.warehouse
				// order.placeorder.payment
				// order.placeorder.shipping
				c.consumePlaceOrderTopic(ft)
			default:
				log.Error().Str("topic", ft.Topic).Msg("consumer: recieved records from unexpected topic")
			}
		})
	}
}

func (c kafkaConsumer) consumePlaceOrderTopic(ft kgo.FetchTopic) {
	log.Info().Str("topic", ft.Topic).Msg("consumer: recieved records from topic")
	ft.EachRecord(func(record *kgo.Record) {
		// Unmarshal the event
		var event pb.PlaceOrderEvent
		err := proto.Unmarshal(record.Value, &event)
		if err != nil {
			log.Panic().Err(err).Msg("consumer: failed to unmarshal place order event")
		}

		// Process the event
		// will currently result in error
		c.svc.EvtCtrl.ProcessPlaceOrderEvent(&event)

		//
		log.Debug().Str("value", string(record.Value)).Msg("")
	})
}

