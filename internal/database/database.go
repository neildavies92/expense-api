package database

import (
	"database/sql"
	"log/slog"

	"github.com/neildavies92/expense-api/config"
	"github.com/neildavies92/expense-api/internal/errors"
	"github.com/neildavies92/expense-api/internal/models"

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

func (db *DB) GetExpenses() ([]models.Expense, error) {
	slog.Info("executing GetExpenses query")

	query := `SELECT id, expense, expense_amount, due_date FROM expenses ORDER BY id DESC`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("failed to execute GetExpenses query",
			"error", err,
			"query", query,
		)
		return nil, errors.ErrNotFound
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var e models.Expense
		if err := rows.Scan(&e.ID, &e.Expense, &e.ExpenseAmount, &e.DueDate); err != nil {
			slog.Error("failed to scan expense row",
				"error", err,
				"query", query,
			)
			return nil, errors.ErrInvalidInput
		}
		expenses = append(expenses, e)
	}

	if err := rows.Err(); err != nil {
		slog.Error("error occurred while iterating expense rows",
			"error", err,
			"query", query,
		)
		return nil, errors.ErrInvalidInput
	}

	slog.Info("GetExpenses query completed successfully",
		"count", len(expenses),
		"query", query,
	)
	return expenses, nil
}

func (db *DB) GetExpenseByID(id int64) (*models.Expense, error) {
	slog.Info("executing GetExpenseByID query", "id", id)

	query := `SELECT id, expense, expense_amount, due_date FROM expenses WHERE id = $1`
	var expense models.Expense
	err := db.QueryRow(query, id).Scan(&expense.ID, &expense.Expense, &expense.ExpenseAmount, &expense.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("expense not found",
				"id", id,
				"query", query,
			)
			return nil, errors.ErrNotFound
		}
		slog.Error("failed to execute GetExpenseByID query",
			"error", err,
			"id", id,
			"query", query,
		)
		return nil, errors.ErrInvalidInput
	}

	slog.Info("GetExpenseByID query completed successfully",
		"id", id,
		"expense", expense.Expense,
		"query", query,
	)
	return &expense, nil
}
