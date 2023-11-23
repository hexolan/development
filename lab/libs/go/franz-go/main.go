package main

import (
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"

	"null.hexolan.dev/exp/app"
)

type KafkaConfig struct {
	Brokers []string
}

func NewKafkaConn(conf KafkaConfig, opts ...kgo.Opt) (*kgo.Client, error) {
	opts = append(opts, kgo.SeedBrokers(conf.Brokers...))
	kcl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return kcl, nil
}


func main() {
	kcl, err := NewKafkaConn(
		KafkaConfig{
			Brokers: []string{"localhost:9092"},
		},
		kgo.ConsumerGroup("something-service"),
		kgo.ConsumeTopics("topic1", "topic2"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("ensuring topics")
	app.EnsureKafkaTopics(kcl, "topic1", "topic2")
	
	fmt.Println("starting")
	go app.ProduceRecords(kcl)
	app.StartConsuming(kcl)
}