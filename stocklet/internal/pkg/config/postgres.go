package config

import (
	"errors"
)

func LoadPostgresConfig() (*PostgresConfig, error) {
	// todo: switch to custom errors interface
	return nil, errors.New("unable to load postgres configuration")

	return &PostgresConfig{
		Username: "",
	}, nil
}

type PostgresConfig struct {
	Username string
	Password string
	Host string
	Database string
}