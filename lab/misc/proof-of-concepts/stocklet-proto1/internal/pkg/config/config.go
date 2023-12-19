package config

import (
	"os"

	"github.com/hexolan/stocklet/internal/pkg/errors"
)

// Load an option from an environment variable
func loadFromEnv(name string) *string {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return nil
	}

	return &value
}

// Require an option from an environment variable
func RequireFromEnv(name string) (string, error) {
	value := loadFromEnv(name)
	if value == nil {
		return "", errors.NewServiceErrorf(errors.ErrCodeService, "failed to load required cfg option (%s)", name)
	}

	return *value, nil
}

// Shared configuration implemented by all services
type SharedConfig struct {
	Otel OtelConfig

	DevMode bool
}

// Load the options in the shared config
func (cfg *SharedConfig) Load() error {
	// determine application mode
	cfg.DevMode = false
	if mode, err := RequireFromEnv("MODE"); err == nil {
		if mode == "dev" || mode == "development" {
			cfg.DevMode = true
		}
	}
	
	// load the Open Telemetry config
	//cfg.Otel = OtelConfig{}
	if err := cfg.Otel.Load(); err != nil {
		return err
	}

	// config succesfully loaded
	return nil
}