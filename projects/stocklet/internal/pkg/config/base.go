package config

// Base configuration implemented by all services
type BaseConfig struct {
	Shared *SharedConfig
}

// Load the base configuration
func (cfg *BaseConfig) LoadBaseConfig() error {
	return cfg.Shared.Load()
}

// Shared configuration loaded under the base config
type SharedConfig struct {
	OtelConfig *OtelConfig
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
	if err := cfg.OtelConfig.Load(); err != nil {
		return err
	}

	// config succesfully loaded
	return nil
}