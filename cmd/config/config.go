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
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
}

func MustParse() *Config {
	cfg := &Config{}
	err := env.Parse(&cfg, env.Options{RequiredIfNoDef: true})
	if err != nil {
		panic(fmt.Errorf("error parsing config: %w", err))
	}
	return cfg
}
