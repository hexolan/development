package order

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type ServiceConfig struct {
	// core configuration
	Shared *config.SharedConfig

	// dynamically loaded configuration
	Postgres *config.PostgresConfig
	Kafka *config.KafkaConfig
}

// load the service config with the core options loaded
func NewServiceConfig() (*ServiceConfig, error) {
	cfg := ServiceConfig{}

	// load the shared config options
	if err := cfg.Shared.Load(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
