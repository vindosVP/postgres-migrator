package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/vindosVP/migrator/cmd/config"
)

var (
	buildCommit = "N/A"
	buildTime   = "N/A"
	version     = "N/A"
)

func main() {
	cfg := config.MustParse()
	fmt.Printf("buildCommit: %s\n", buildCommit)
	fmt.Printf("buildTime: %s\n", buildTime)
	fmt.Printf("version: %s\n", version)
	if _, err := os.Stat(cfg.MigrationsPath); errors.Is(err, os.ErrNotExist) {
		panic(fmt.Errorf("failed to open migrations file: %w", err))
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.MigrationsPath),
		postgresDSN(cfg),
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(fmt.Errorf("failed to apply migrations: %w", err))
	}
	fmt.Println("migrations applied")
}

func postgresDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)
}
