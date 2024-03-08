package config

type OtelConfig struct {
	// Env Var: "OTEL_COLLECTOR_GRPC"
	CollectorGrpc string
}

func (cfg *OtelConfig) Load() error {
	// Load configurations from env
	if collectorGrpc, err := RequireFromEnv("OTEL_COLLECTOR_GRPC"); err != nil {
		return err
	} else {
		cfg.CollectorGrpc = collectorGrpc
	}

	// Succesfully loaded all config properties
	return nil
}
