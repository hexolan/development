package api

import (
	"context"
	
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
	
	"github.com/hexolan/stocklet/internal/svc/order"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type eventConsumer struct {

}

// plan:
// this is the generic consumption interface
// a generic method is called in the controller to attach the consumer interface to the EventController or something.
// this returns a method to actually start the consumption after prepartation?
func AttachSvcToConsumer(cfg *order.ServiceConfig, svc *order.OrderService) error {
	// todo: implement
	return nil
}

func (c eventConsumer) ConsumePlaceOrderTopic()
// on second thoughts this may need a rethink

// =====================================================
// before refactoring:
// todo: clean up
type kafkaConsumer struct {
	svc order.OrderService
	
	kCl *kgo.Client
}

func NewKafkaConsumer(svc order.OrderService, kCl *kgo.Client) kafkaConsumer {
	return kafkaConsumer{
		svc: svc,
		kCl: kCl,
	}
}

func (c kafkaConsumer) StartConsuming() {
	ctx := context.Background()

	for {
		fetches := c.kCl.PollFetches(ctx)
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

		// Process the Event
		ctx := context.Background()
		c.svc.ProcessPlaceOrderEvent(ctx, &event)

		//
		log.Debug().Str("value", string(record.Value)).Msg("")
	})
}

