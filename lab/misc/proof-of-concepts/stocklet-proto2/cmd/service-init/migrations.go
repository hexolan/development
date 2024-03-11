// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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