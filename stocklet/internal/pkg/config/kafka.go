package config

import (
	"errors"
)

func LoadKafkaConfig() (*KafkaConfig, error) {
	// todo: switch to custom errors interface
	return nil, errors.New("unable to load kafka configuration")
	return KafkaConfig{}, nil
}

type KafkaConfig struct {
	Username string
	Password string
	Host string
}