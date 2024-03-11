package main

import (
	"github.com/rs/zerolog/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

func applyPostgresMigrations(conf *config.PostgresConfig) {
    m, err := migrate.New("file:///migrations", conf.GetDSN())
	if err != nil {
		log.Panic().Err(err).Msg("migrate: failed to open client")
	}

    err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Info().Err(err).Msg("migrate: migrations up to date")
		} else {
			log.Panic().Err(err).Msg("migrate: raised when performing db migration")
		}
	}

	log.Info().Msg("migrate: succesfully performed postgres migrations")
}