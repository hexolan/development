package main

import (
	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

func main() {
	// Load the init container cfg options
	cfg := InitConfig{}
	if err := cfg.Load(); err != nil {
		log.Panic().Err(err).Msg("missing required configuration")
	}

	// If migrations or debezium are enabled,
	// then a database configuration will be required.
	if cfg.ApplyMigrations || cfg.ApplyDebezium {
		pgConf := config.PostgresConfig{}
		if err := pgConf.Load(); err == nil {
			// Using postgres as a database.
			if cfg.ApplyMigrations {
				applyPostgresMigrations(&pgConf)
			}

			if cfg.ApplyDebezium {
				applyPostgresOutbox(&cfg, &pgConf)
			}
		} else {
			log.Panic().Msg("unable to load any db configs (unable to perform migrations or apply connector cfgs)")
		}
	}

	log.Info().Msg("completed init for service")
}