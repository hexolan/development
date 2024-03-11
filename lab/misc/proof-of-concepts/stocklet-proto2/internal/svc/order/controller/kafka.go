// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

type kafkaController struct {
	kCl *kgo.Client
}

func NewKafkaController(kCl *kgo.Client) order.EventController {
	// Ensure the required Kafka topics exist
	err := messaging.EnsureKafkaTopics(
		kCl,

		messaging.Order_State_Created_Topic,
		messaging.Order_State_Updated_Topic,
		messaging.Order_State_Deleted_Topic,

		messaging.Order_PlaceOrder_Order_Topic,
		messaging.Order_PlaceOrder_Warehouse_Topic,
		messaging.Order_PlaceOrder_Payment_Topic,
		messaging.Order_PlaceOrder_Shipping_Topic,
	)
	if err != nil {
		log.Warn().Err(err).Msg("kafka: raised attempting to ensure svc topics")
	}

	// Create the controller
	return kafkaController{kCl: kCl}
}

func (c kafkaController) PrepareConsumer(svc *order.OrderService) messaging.EventConsumerController {
	// Create a cancellable context for the consumer
	ctx, ctxCancel := context.WithCancel(context.Background())

	// Add the consumption topics
	c.kCl.AddConsumeTopics(
		messaging.Order_PlaceOrder_Warehouse_Topic,
		messaging.Order_PlaceOrder_Payment_Topic,
		messaging.Order_PlaceOrder_Shipping_Topic,
	)

	// Create the consumer controller
	return kafkaConsumerController{
		svc: svc,
		kCl: c.kCl,
		
		ctx: ctx,
		ctxCancel: ctxCancel,
	}
}

type kafkaConsumerController struct {
	svc *order.OrderService
	kCl *kgo.Client

	ctx context.Context
	ctxCancel context.CancelFunc
}

func (c kafkaConsumerController) Start() {
	for {
		fetches := c.kCl.PollFetches(c.ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Panic().Any("kafka-errs", errs).Msg("consumer: unrecoverable kafka errors")
		}

		fetches.EachTopic(func(ft kgo.FetchTopic) {
			switch ft.Topic {
			case messaging.Order_PlaceOrder_Warehouse_Topic:
				c.consumePlaceOrderTopic(ft)
			case messaging.Order_PlaceOrder_Payment_Topic:
				c.consumePlaceOrderTopic(ft)
			case messaging.Order_PlaceOrder_Shipping_Topic:
				c.consumePlaceOrderTopic(ft)
			default:
				log.Warn().Str("topic", ft.Topic).Msg("consumer: recieved records from unexpected topic")
			}
		})
	}
}

func (c kafkaConsumerController) Stop() {
	// Cancel the consumer context
	c.ctxCancel()
}

func (c kafkaConsumerController) consumePlaceOrderTopic(ft kgo.FetchTopic) {
	log.Info().Str("topic", ft.Topic).Msg("consumer: recieved records from topic")

	// Process each message from the topic
	ft.EachRecord(func(record *kgo.Record) {
		// Unmarshal the event
		var event pb.PlaceOrderEvent
		err := proto.Unmarshal(record.Value, &event)
		if err != nil {
			log.Panic().Err(err).Msg("consumer: failed to unmarshal place order event")
		}

		// Process the event
		ctx := context.Background()
		c.svc.ProcessPlaceOrderEvent(ctx, &event)
	})
}