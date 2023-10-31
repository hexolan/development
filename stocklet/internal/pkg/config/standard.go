package config

type StandardConfig struct {
	Postgres PostgresConfig
	Kafka    KafkaConfig
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

	cfg.Postgres = *pgConf
	cfg.Kafka = *kafkaConf
	return nil
}
