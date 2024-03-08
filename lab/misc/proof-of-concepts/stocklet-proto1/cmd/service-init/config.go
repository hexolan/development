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
	
	// Env Var: "INIT_MIGRATIONS" (optional)
	// accepted: 'false'
	// Defaults to true
	ApplyMigrations bool

	// Env Var: "INIT_DEBEZIUM_HOST" (optional)
	// e.g. "http://debezium:8083"
	//
	// 'ApplyDebezium' will default to false unless
	// the debezium host is provided.
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