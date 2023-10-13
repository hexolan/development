package config

type KafkaConfig struct {
	Username string
	Password string
	Host string
}

func LoadKafkaConfig() (*KafkaConfig, error) {
	// Load configurations from env
	username, err := requireFromEnv("KAFKA_USER")
	if err != nil {
		return nil, err
	}
	
	password, err := requireFromEnv("KAFKA_PASS")
	if err != nil {
		return nil, err
	}
	
	host, err := requireFromEnv("KAFKA_HOST")
	if err != nil {
		return nil, err
	}

	return &KafkaConfig{
		Username: *username,
		Password: *password,
		Host: *host,
	}, nil
}