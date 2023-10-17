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

func EnsureKafkaStreamTopics(conf config.KafkaConfig, topics []string) error {
	const TopicPartitions = 8

	tm, err := goka.NewTopicManager(conf.Brokers, goka.DefaultConfig(), goka.NewTopicManagerConfig())
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "failed to create topic manager", err)
	}

	defer tm.Close()

	for _, topic := range topics {
		err = tm.EnsureStreamExists(topic, TopicPartitions)
	}

	return nil
}