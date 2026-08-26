package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/victoroluborode/leaflet-backend/internal/config"
)

func NewLogger(cfg *config.ObservabilityConfig) zerolog.Logger {
	var logLevel zerolog.Level
	switch cfg.Logging.Level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	var logger zerolog.Logger
	if cfg.Logging.Format == "json" {
		logger = zerolog.New(os.Stdout).Level(logLevel).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(consoleWriter).Level(logLevel).With().Timestamp().Logger()
	}

	return logger
}