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

package config

import (
	"fmt"
)

type PostgresConfig struct {
	// Env Var: "PG_USER"
	Username string

	// Env Var: "PG_PASS"
	Password string

	// Env Var: "PG_HOST"
	Host string

	// Env Var: "PG_PORT" (optional)
	// Defaults to "5432"
	Port string
	
	// Env Var: "PG_DB"
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
