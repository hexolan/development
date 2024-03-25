package api

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

type kafkaConsumer struct {
	kCl *kgo.Client
}

func NewKafkaConsumer(kCl *kgo.Client) kafkaConsumer {
	return kafkaConsumer{
		kCl: kCl,
	}
}