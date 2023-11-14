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