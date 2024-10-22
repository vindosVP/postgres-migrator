package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DB             DB
	MigrationsPath string `env:"MIGRATIONS_PATH"`
}

type DB struct {
	Host     string `env:"MIGRATIONS_DB_HOST"`
	Port     int    `env:"MIGRATIONS_DB_PORT"`
	Username string `env:"MIGRATIONS_DB_USERNAME"`
	Password string `env:"MIGRATIONS_DB_PASSWORD"`
	Database string `env:"MIGRATIONS_DB_DATABASE"`
}

func MustParse() *Config {
	cfg := &Config{}
	err := env.Parse(cfg, env.Options{RequiredIfNoDef: true})
	if err != nil {
		panic(fmt.Errorf("error parsing config: %w", err))
	}
	return cfg
}
