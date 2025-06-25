package repositories

import (
	"context"
	"fmt"
	"github.com/donskova1ex/application_aggregator/config"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(ctx context.Context, postgresConfig *config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", postgresConfig.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}
	return db, nil

}
