package api

import (
	"inno-accounting/internal/adapters/jwt"
	"inno-accounting/internal/adapters/postgres"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	PostgresURI *postgres.Config `toml:"storage"`
	JWT *jwt.Config `toml:"jwt"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		PostgresURI: &postgres.Config{},
	}
}