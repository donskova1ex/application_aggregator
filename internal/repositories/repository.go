package repositories

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type PostgresRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewPostgresRepository(db *sqlx.DB, log *slog.Logger) *PostgresRepository {
	return &PostgresRepository{
		db:  db,
		log: log,
	}
}
