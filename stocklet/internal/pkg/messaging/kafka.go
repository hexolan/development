package messaging

import (
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func NewKafkaEmitter(conf config.KafkaConfig, topic goka.Stream) (*goka.Emitter, error) {
	emitter, err := goka.NewEmitter(conf.Brokers, topic, new(codec.Bytes))
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to create Kafka emitter", err)
	}

	return emitter, nil
}

func NewKafkaProcessor(conf config.KafkaConfig, group *goka.GroupGraph) (*goka.Processor, error) {
	processor, err := goka.NewProcessor(conf.Brokers, group)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to create Kafka processor", err)
	}

	return processor, nil
}

func EnsureKafkaTopics(conf config.KafkaConfig, topics []string) error {
	// todo
	return nil
}