package config

import (
	"strings"
)

type KafkaConfig struct {
	Brokers []string
}

func LoadKafkaConfig() (*KafkaConfig, error) {
	// Load configurations from env
	brokersOpt, err := requireFromEnv("KAFKA_BROKERS")
	if err != nil {
		return nil, err
	}

	// Comma seperate the brokers
	brokers := strings.Split(*brokersOpt, ",")
	
	return &KafkaConfig{
		Brokers: brokers,
	}, nil
}