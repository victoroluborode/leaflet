package config

import (
	"os"
	"strings"



	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Primary  PrimaryConfig  `koanf:"primary" validate:"required"`
	Server   ServerConfig   `koanf:"server" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
}

type PrimaryConfig struct {
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct {
	Port string `koanf:"port" validate:"required"`
	ReadTimeout int `koanf:"read_timeeout" validate:"required"`
	WriteTimeout int    `koanf:"write_timeout" validate:"required"`
	IdleTimeout  int    `koanf:"idle_timeout" validate:"required"`
}

type DatabaseConfig struct {
	Host    string `koanf:"host" validate:"required"`
	Port    int    `koanf:"port" validate:"required"`
	User    string `koanf:"user" validate:"required"`
	Password string `koanf:"password"`
	Name    string `koanf:"name" validate:"required"`
	SSLMode string `koanf:"ssl_mode" validate:"required"`
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	k := koanf.New(".")

	if err := k.Load(env.Provider("LEAFLET_", ".", func(s string) string {
	return strings.ToLower(strings.TrimPrefix(s, "LEAFLET_"))
}), nil); err != nil {
	logger.Fatal().Err(err).Msg("could not load env variables")
}

	mainConfig := &Config{}
	if err := k.Unmarshal("", mainConfig); err != nil {
	logger.Fatal().Err(err).Msg("could not unmarshal config")
}
	if err := validator.New().Struct(mainConfig); err != nil {
	logger.Fatal().Err(err).Msg("config validation failed")
}

 return mainConfig, nil
}