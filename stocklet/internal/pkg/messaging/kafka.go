package messaging

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kadm"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func NewKafkaConn(conf config.KafkaConfig, opts ...kgo.Opt) (*kgo.Client, error) {
	// todo: passing options
	// kgo.ConsumerGroup("something-service")
	// kgo.ConsumeTopics("topic1", "topic2")
	
	opts = append(opts, kgo.SeedBrokers(conf.Brokers...))
	kcl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to connect to Kafka", err)
	}

	return kcl, nil
}

func EnsureKafkaTopics(kcl *kgo.Client, topics ...string) error {
	kadmcl := kadm.NewClient(kcl)

	ctx := context.Background()
	// or something: context.WithTimeout(ctx, time.Second * 30)

	// todo: check functionality
	_, err := kadmcl.CreateTopics(ctx, -1, -1, nil, topics...)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to create topics")
	}
}