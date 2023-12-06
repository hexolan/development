package config

import (
	"fmt"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Database string
}

func (conf PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Database,
	)
}

func LoadPostgresConfig() (*PostgresConfig, error) {
	// Load configurations from env
	username, err := RequireFromEnv("PG_USER")
	if err != nil {
		return nil, err
	}

	password, err := RequireFromEnv("PG_PASS")
	if err != nil {
		return nil, err
	}

	host, err := RequireFromEnv("PG_HOST")
	if err != nil {
		return nil, err
	}

	database, err := RequireFromEnv("PG_DB")
	if err != nil {
		return nil, err
	}

	return &PostgresConfig{
		Username: username,
		Password: password,
		Host:     host,
		Database: database,
	}, nil
}
