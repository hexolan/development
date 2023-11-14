package clients

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kadm"

	"github.com/hexolan/development/experimentation/gateway-exper1/utils/config"
	"github.com/hexolan/development/experimentation/gateway-exper1/utils/errors"
)

func NewKafkaConn(conf *config.KafkaConfig, opts ...kgo.Opt) (*kgo.Client, error) {
	opts = append(opts, kgo.SeedBrokers(conf.Brokers...))
	kCl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to connect to Kafka", err)
	}

	return kCl, nil
}

func EnsureKafkaTopics(kcl *kgo.Client, topics ...string) error {
	ctx := context.Background()
	kadmCl := kadm.NewClient(kcl)

	_, err := kadmCl.CreateTopics(ctx, -1, -1, nil, topics...)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to create Kafka topics", err)
	}

	return nil
}
