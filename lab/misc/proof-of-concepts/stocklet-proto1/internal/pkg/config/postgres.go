package config

import (
	"fmt"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (conf *PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
}

func (cfg *PostgresConfig) Load() error {
	// Load configurations from env
	if opt, err := RequireFromEnv("PG_USER"); err != nil {
		return err
	} else {
		cfg.Username = opt
	}

	if opt, err := RequireFromEnv("PG_PASS"); err != nil {
		return err
	} else {
		cfg.Password = opt
	}

	if opt, err := RequireFromEnv("PG_HOST"); err != nil {
		return err
	} else {
		cfg.Host = opt
	}

	if opt, err := RequireFromEnv("PG_PORT"); err != nil {
		cfg.Port = "5432"
	} else {
		cfg.Port = opt
	}

	if opt, err := RequireFromEnv("PG_DB"); err != nil {
		return err
	} else {
		cfg.Database = opt		
	}

	// config properties succesfully loaded
	return nil
}
