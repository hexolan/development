package controller

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type kafkaConsumerController struct {
	svc *order.OrderService
	kCl *kgo.Client

	ctx context.Context
	ctxCancel context.CancelFunc
}

func newKafkaConsumerController(svc *order.OrderService, kCl *kgo.Client) messaging.EventConsumerController {
	// create a cancellable context for the consumer
	ctx, ctxCancel := context.WithCancel(context.Background())

	// add the consumption topics
	kCl.AddConsumeTopics(
		messaging.Order_PlaceOrder_Catchall,
	)

	// create the consumer controller
	return kafkaConsumerController{
		svc: svc,
		kCl: kCl,
		
		ctx: ctx,
		ctxCancel: ctxCancel,
	}
}

func (c kafkaConsumerController) Start() {
	for {
		fetches := c.kCl.PollFetches(c.ctx)
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

func (c kafkaConsumerController) Stop() {
	// cancel the consumer context
	c.ctxCancel()
}

func (c kafkaConsumerController) consumePlaceOrderTopic(ft kgo.FetchTopic) {
	log.Info().Str("topic", ft.Topic).Msg("consumer: recieved records from topic")
	ft.EachRecord(func(record *kgo.Record) {
		// Unmarshal the event
		var event pb.PlaceOrderEvent
		err := proto.Unmarshal(record.Value, &event)
		if err != nil {
			log.Panic().Err(err).Msg("consumer: failed to unmarshal place order event")
		}

		// Process the Event
		ctx := context.Background()
		c.svc.ProcessPlaceOrderEvent(ctx, &event)

		//
		log.Debug().Str("value", string(record.Value)).Msg("")
	})
}

