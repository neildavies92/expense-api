package database

import (
	"database/sql"
	"log/slog"

	"github.com/bryx/expense-api/config"
	"github.com/bryx/expense-api/internal/errors"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewConnection(cfg config.DatabaseConfig) (*DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, errors.ErrInvalidInput
	}

	if err := db.Ping(); err != nil {
		return nil, errors.ErrNotFound
	}

	slog.Info("successfully connected to database")
	return &DB{db}, nil
}
