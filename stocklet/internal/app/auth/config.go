package auth

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type ServiceConfig struct {
	Postgres config.PostgresConfig
	Kafka config.KafkaConfig
}

func LoadServiceConfig() (*ServiceConfig, error) {
	pgCfg, err := config.LoadPostgresConfig()
	if err != nil {
		return nil, err
	}

	kafkaCfg, err := config.LoadKafkaConfig()
	if err != nil {
		return nil, err
	}

	return &ServiceConfig{
		Postgres: *pgCfg,
		Kafka: *kafkaCfg,
	}, nil
}