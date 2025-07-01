package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGDB  *PostgresConfig
	SQLDB *SQLConfig
}

type PostgresConfig struct {
	DSN string
}

type SQLConfig struct {
	DSN string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env.dev")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file, %w", err)
	}

	pgDsn := os.Getenv("POSTGRES_DSN")
	if pgDsn == "" {
		return nil, fmt.Errorf("empty .env file, %w", errors.New("POSTGRES_DSN is not set"))
	}

	sqlDSN := os.Getenv(("SQL_DSN"))
	if sqlDSN == "" {
		return nil, fmt.Errorf("empty .env file, %w", errors.New("SQL_DSN is not set"))
	}

	return &Config{
		PGDB: &PostgresConfig{
			DSN: pgDsn,
		},
		SQLDB: &SQLConfig{
			DSN: sqlDSN,
		},
	}, nil
}
