package auth

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type Config struct {
	Postgres config.PostgresConfig
	Kafka config.KafkaConfig
}

func LoadConfig() (*Config, error) {
	postgresCfg, err := config.LoadPostgresConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := config.LoadKafkaConfig()
	if err != nil {
		return err
	}

	return &Config{
		Postgres: postgresCfg,
		Kafka: kafkaCfg,
	}, nil
}