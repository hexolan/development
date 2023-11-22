package config

type StandardConfig struct {
	DevMode bool
	Postgres PostgresConfig
	Kafka KafkaConfig
}

func (cfg *StandardConfig) LoadStandardConfig() error {
	pgConf, err := LoadPostgresConfig()
	if err != nil {
		return err
	}

	kafkaConf, err := LoadKafkaConfig()
	if err != nil {
		return err
	}

	cfg.DevMode = false
	mode := loadFromEnv("MODE")
	if mode != nil {
		if *mode == "dev" || *mode == "development" {
			cfg.DevMode = true
		}
	}

	cfg.Postgres = *pgConf
	cfg.Kafka = *kafkaConf
	return nil
}
