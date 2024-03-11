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

	log.Info().Str("svc", cfg.ServiceName).Msg("completed init for service")
}