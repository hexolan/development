package messaging

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kadm"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func NewKafkaConn(conf config.KafkaConfig, opts ...kgo.Opt) (*kgo.Client, error) {
	opts = append(opts, kgo.SeedBrokers(conf.Brokers...))
	kcl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to connect to Kafka", err)
	}

	return kcl, nil
}

func EnsureKafkaTopics(kcl *kgo.Client, topics ...string) error {
	ctx := context.Background()
	kadmcl := kadm.NewClient(kcl)

	_, err := kadmcl.CreateTopics(ctx, -1, -1, nil, topics...)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to create Kafka topics", err)
	}
}
