package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"

	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func NewEventConsumer(kcl *kgo.Client) {
	ctx := context.Background()

	for {
		fetches := kcl.PollFetches(ctx)
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
				consumePlaceOrderEvents(ft)
			default:
				log.Error().Str("topic", ft.Topic).Msg("consumer: recieved records from unexpected topic")
			}
		})
	}
}

func consumePlaceOrderEvents(ft kgo.FetchTopic) {
	log.Info().Str("topic", ft.Topic).Msg("consumer: recieved records from topic")
	ft.EachRecord(func(record *kgo.Record) {
		// Unmarshal the event
		var event pb.PlaceOrderEvent
		err := proto.Unmarshal(record.Value, &event)
		if err != nil {
			log.Panic().Err(err).Msg("consumer: failed to unmarshal place order event")
		}

		// Process the event
		processPlaceOrderEvent(&event)

		//
		log.Debug().Str("value", string(record.Value)).Msg("")
	})
}

func processPlaceOrderEvent(event *pb.PlaceOrderEvent) {
	// Ignore events dispatched by the order service
	if event.Type == pb.PlaceOrderEvent_TYPE_UNSPECIFIED || event.Status == pb.PlaceOrderEvent_STATUS_UNSPECIFIED {
		return
	}

	// Mark the order as rejected if a failure status
	// was reported at any stage.
	if event.Status == pb.PlaceOrderEvent_STATUS_FAILURE {
		// todo
		return
	}
	
	// Otherwise,
	// If the event is from the last stage of the saga (shipping svc)
	// then mark the order as succesful.
	if event.Type == pb.PlaceOrderEvent_TYPE_SHIPPING {
		// todo: update order status and details
		// (append transaction id, etc to stored order)
	}
}