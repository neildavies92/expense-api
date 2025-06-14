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

func (db *DB) GetExpenses() ([]Expense, error) {
	query := `SELECT id, expense, expense_amount, due_date FROM expenses ORDER BY id DESC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.ID, &e.Expense, &e.ExpenseAmount, &e.DueDate); err != nil {
			return nil, errors.ErrInvalidInput
		}
		expenses = append(expenses, e)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.ErrInvalidInput
	}

	return expenses, nil
}
