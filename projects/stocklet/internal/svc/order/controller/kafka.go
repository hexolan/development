package controller

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type kafkaController struct {
	kCl *kgo.Client
}

func NewKafkaController(kCl *kgo.Client) order.EventController {
	// ensure the required Kafka topics exist
	err := messaging.EnsureKafkaTopics(
		kCl,

		messaging.Order_State_Created_Topic,
		messaging.Order_State_Updated_Topic,
		messaging.Order_State_Deleted_Topic,

		messaging.Order_PlaceOrder_Order_Topic,
		messaging.Order_PlaceOrder_Payment_Topic,
		messaging.Order_PlaceOrder_Shipping_Topic,
		messaging.Order_PlaceOrder_Warehouse_Topic,
	)
	if err != nil {
		log.Error().Err(err).Msg("raised attempting to creating Kafka topics")
	}

	// create the controller
	return kafkaController{kCl: kCl}
}

func (c kafkaController) PrepareConsumer(svc *order.OrderService) order.EventConsumerController {
	return newKafkaConsumerController(svc, c.kCl)
}

func (c kafkaController) dispatchEvent(topic string, wireEvt []byte) {
	ctx := context.Background()

	c.kCl.Produce(
		ctx,
		&kgo.Record{
			Topic: topic,
			Value: wireEvt,
		},
		nil,
	)
}

func (c kafkaController) marshalEvent(evt protoreflect.ProtoMessage) []byte {
	wireEvt, err := proto.Marshal(evt)
	if err != nil {
		// todo: handling
		panic(err)
	}

	return wireEvt
}

func (c kafkaController) DispatchCreatedEvent(order *pb.Order) {
	c.dispatchEvent(
		messaging.Order_State_Created_Topic,
		c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_CREATED,
				Payload: order,
			},
		),
	)
}

func (c kafkaController) DispatchUpdatedEvent(order *pb.Order) {
	c.dispatchEvent(
		messaging.Order_State_Updated_Topic,
		c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_UPDATED,
				Payload: order,
			},
		),
	)
}

func (c kafkaController) DispatchDeletedEvent(req *pb.CancelOrderRequest) {
	// todo: improve assembly of payload (dispatch whole order?)
	c.dispatchEvent(
		messaging.Order_State_Deleted_Topic,
		c.marshalEvent(
			&pb.OrderStateEvent{
				Type: pb.OrderStateEvent_TYPE_DELETED,
				Payload: &pb.Order{
					Id: req.GetOrderId(),
				},
			},
		),
	)
}

func (c kafkaController) DispatchPlaceOrderEvent(evt *pb.PlaceOrderEvent) {
	// todo:
	c.dispatchEvent(
		messaging.Order_PlaceOrder_Order_Topic,
		c.marshalEvent(evt),
	)
}

// ====== EVENT CONSUMER ========
type kafkaConsumerController struct {
	svc *order.OrderService
	kCl *kgo.Client

	ctx context.Context
	ctxCancel context.CancelFunc
}

func newKafkaConsumerController(svc *order.OrderService, kCl *kgo.Client) order.EventConsumerController {
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

