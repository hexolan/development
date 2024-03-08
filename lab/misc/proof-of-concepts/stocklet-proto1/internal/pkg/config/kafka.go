package config

import (
	"strings"
)

type KafkaConfig struct {
	// Env Var: "KAFKA_BROKERS"
	// Comma delimited from env var.
	Brokers []string
}

func (cfg *KafkaConfig) Load() error {
	// load configurations from env
	brokersOpt, err := RequireFromEnv("KAFKA_BROKERS")
	if err != nil {
		return err
	}

	// comma seperate the kafka brokers
	cfg.Brokers = strings.Split(brokersOpt, ",")
	
	// config options were succesfully set and loaded
	return nil
}
