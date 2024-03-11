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
	"github.com/hexolan/stocklet/internal/pkg/config"
)

// Init Container Configuration
type InitConfig struct {
	// Name of the service (e.g. 'auth' or 'order')
	//
	// Env Var: "INIT_SVC_NAME"
	ServiceName string
	
	// Env Var: "INIT_MIGRATIONS" (optional. accepts 'false')
	// Defaults to true
	ApplyMigrations bool

	// 'ApplyDebezium' will default to false unless
	// the debezium host is provided.
	//
	// Env Var: "INIT_DEBEZIUM_HOST" (optional)
	// e.g. "http://debezium:8083"
	ApplyDebezium bool
	DebeziumHost string
}

func (opts *InitConfig) Load() error {
	// ServiceName
	opt, err := config.RequireFromEnv("INIT_SVC_NAME") 
	if err != nil {
		return err
	}
	opts.ServiceName = opt

	// ApplyMigrations
	opts.ApplyMigrations = true
	if opt, _ := config.RequireFromEnv("INIT_MIGRATIONS"); opt == "false" {
		opts.ApplyMigrations = false
	}

	// ApplyDebezium and DebeziumHost
	if opt, err := config.RequireFromEnv("INIT_DEBEZIUM_HOST"); err == nil {
		opts.ApplyDebezium = true
		opts.DebeziumHost = opt
	}

	return nil
}