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

func (conf *PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Database,
	)
}

func (cfg *PostgresConfig) Load() error {
	// Load configurations from env
	if username, err := RequireFromEnv("PG_USER"); err != nil {
		return err
	} else {
		cfg.Username = username
	}

	if password, err := RequireFromEnv("PG_PASS"); err != nil {
		return err
	} else {
		cfg.Password = password
	}

	if host, err := RequireFromEnv("PG_HOST"); err != nil {
		return err
	} else {
		cfg.Host = host
	}

	if database, err := RequireFromEnv("PG_DB"); err != nil {
		return err
	} else {
		cfg.Database = database		
	}

	// config properties succesfully loaded
	return nil
}
