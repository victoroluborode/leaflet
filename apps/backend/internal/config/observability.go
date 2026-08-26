package config

type ObservabilityConfig struct {
	Logging LoggingConfig `koanf:"logging" validate:"required"`
}

type LoggingConfig struct {
	Level  string `koanf:"level" validate:"required,oneof=debug info warn error"`
	Format string `koanf:"format" validate:"required,oneof=json console"`
}

func DefaultObservabilityConfig() *ObservabilityConfig {
	return &ObservabilityConfig{
		Logging: LoggingConfig{
			Level:  "info",
			Format: "console",
		},
	}
}