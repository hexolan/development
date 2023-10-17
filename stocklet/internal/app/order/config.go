package order

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type ServiceConfig struct {
	config.StandardConfig
}

func NewServiceConfig() (*ServiceConfig, error) {
	cfg := ServiceConfig{}

	err := cfg.LoadStandardConfig()
	if err != nil {
		return nil, err
	}
	
	return &cfg, nil
}