package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DB *PostgresConfig
}

type PostgresConfig struct {
	DSN string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file, %w", err)
	}

	dbDsn := os.Getenv("POSTGRES_DSN")
	if dbDsn == "" {
		return nil, fmt.Errorf("empty .env file, %w", errors.New("POSTGRES_DSN is not set"))
	}

	return &Config{
		DB: &PostgresConfig{
			DSN: dbDsn,
		},
	}, nil
}
