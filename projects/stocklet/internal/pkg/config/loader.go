package config

import (
	"os"

	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func LoadFromEnv(name string) *string {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return nil
	}

	return &value
}

func RequireFromEnv(name string) (string, error) {
	value := LoadFromEnv(name)
	if value == nil {
		return "", errors.NewServiceErrorf(errors.ErrCodeService, "failed to load required config option (%s)", name)
	}

	return *value, nil
}
