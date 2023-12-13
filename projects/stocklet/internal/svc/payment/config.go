package payment

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

// Payment Service Configuration
type ServiceConfig struct {
	// Core Configuration
	Shared *config.SharedConfig

	// Dynamically loaded configuration
	Postgres *config.PostgresConfig
	Kafka *config.KafkaConfig
}

// load the base service configuration
func NewServiceConfig() (*ServiceConfig, error) {
	cfg := ServiceConfig{}

	// Load the core configuration
	if err := cfg.Shared.Load(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
