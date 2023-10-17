package api

import (
	"os"
	"os/signal"
	"syscall"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	"github.com/hexolan/stocklet/internal/app/order"
)


func NewEventAPI(svc order.OrderRepository, kafkaConf config.KafkaConfig) {
	group := goka.DefineGroup(
		"order-service",
		goka.Input(messaging.TopicTestThing, new(codec.Bytes), topicTestThingHandler(svc)),
	)

	processor, err := messaging.NewKafkaProcessor(kafkaConf, group)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start Kafka processor")
	}

	// run the processor (turn to generic call in pkg?)
	ctx, cancel := context.WithCancel(context.Background())
	doneProcessing := make(chan bool)
	go func() {
		defer close(doneProcessing)
		if err = processor.Run(ctx); err != nil {
			log.Fatal().Err(err).Msg("error whilst processing Kafka topics")
		} else {
			log.Info().Msg("kafka processor has stopped")
		}
	}()

	// todo: improve somehow?
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait
	cancel()
	<-doneProcessing
}

func topicTestThingHandler(svc order.OrderRepository) goka.ProcessCallback {
	return func(ctx goka.Context, msg interface{}) {
		
	}
}