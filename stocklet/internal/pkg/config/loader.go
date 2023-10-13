package config

import (
	"os"

	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func loadFromEnv(name string) *string {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return nil
	}

	return &value
}

func requireFromEnv(name string) (*string, error) {
	value := loadFromEnv(name)
	if value == nil {
		return nil, errors.NewServiceErrorf(errors.ErrCodeService, "failed to load required config option (%s)", name)
	}

	return value, nil
}