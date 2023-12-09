package config

// Shared configuration implemented by all services
type SharedConfig struct {
	Otel *OtelConfig

	DevMode bool
}

// Load the options in the shared config
func (cfg *SharedConfig) Load() error {
	// determine application mode
	cfg.DevMode = false
	if mode := LoadFromEnv("MODE"); mode != nil {
		if *mode == "dev" || *mode == "development" {
			cfg.DevMode = true
		}
	}
	
	// load the Open Telemetry config
	if err := cfg.Otel.Load(); err != nil {
		return err
	}

	// config succesfully loaded
	return nil
}