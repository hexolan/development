package auth

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type ServiceConfig struct {
	Postgres config.PostgresConfig
	Kafka config.KafkaConfig
}

func LoadServiceConfig() (*ServiceConfig, error) {
	pgConf, err := config.LoadPostgresConfig()
	if err != nil {
		return nil, err
	}

	kafkaConf, err := config.LoadKafkaConfig()
	if err != nil {
		return nil, err
	}

	return &ServiceConfig{
		Postgres: *pgConf,
		Kafka: *kafkaConf,
	}, nil
}