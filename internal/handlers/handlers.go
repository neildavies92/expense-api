package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/neildavies92/expense-api/internal/database"
	"github.com/neildavies92/expense-api/internal/errors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	db *database.DB
}

func NewHandler(db *database.DB) *Handler {
	return &Handler{db: db}
}

// requestContext returns common request attributes for logging
func requestContext(r *http.Request) []any {
	return []any{
		"method", r.Method,
		"path", r.URL.Path,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	}
}

func SetupRoutes(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/health", h.handleHealthCheck)

	r.Route("/expense", func(r chi.Router) {
		r.Get("/", h.handleGetExpenses)
		r.Post("/", h.handlePostExpense)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.handleGetExpense)
		})
	})
	return r
}

func (h *Handler) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	slog.Info("health check requested", requestContext(r)...)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	slog.Info("health check completed",
		append(requestContext(r), "status", http.StatusOK)...,
	)
}

func (h *Handler) handlePostExpense(w http.ResponseWriter, r *http.Request) {
	slog.Info("creating new expense", requestContext(r)...)

	var body struct {
		Expense       string  `json:"expense"`
		ExpenseAmount float64 `json:"expense_amount"`
		DueDate       int     `json:"due_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		slog.Error("failed to decode request body",
			append(requestContext(r), "error", err)...,
		)

		status := errors.HTTPStatus(errors.ErrInvalidInput)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": errors.ErrorMessage(errors.ErrInvalidInput),
		})
		return
	}

	slog.Info("expense data received",
		append(requestContext(r),
			"expense", body.Expense,
			"amount", body.ExpenseAmount,
			"due_date", body.DueDate,
		)...,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   body,
	})

	slog.Info("expense created successfully",
		append(requestContext(r),
			"status", http.StatusOK,
			"expense", body.Expense,
		)...,
	)
}

func (h *Handler) handleGetExpenses(w http.ResponseWriter, r *http.Request) {
	slog.Info("fetching all expenses", requestContext(r)...)

	expenses, err := h.db.GetExpenses()
	if err != nil {
		slog.Error("failed to fetch expenses",
			append(requestContext(r), "error", err)...,
		)

		status := errors.HTTPStatus(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": errors.ErrorMessage(err),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"data":   expenses,
	})

	slog.Info("expenses fetched successfully",
		append(requestContext(r),
			"status", http.StatusOK,
			"count", len(expenses),
		)...,
	)
}

func (h *Handler) handleGetExpense(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	slog.Info("fetching expense by ID",
		append(requestContext(r), "id", idStr)...,
	)

	if idStr == "" {
		slog.Error("missing expense ID parameter", requestContext(r)...)

		status := errors.HTTPStatus(errors.ErrInvalidInput)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": errors.ErrorMessage(errors.ErrInvalidInput),
		})
		return
	}

	// Parse the ID from string to int64
	var id int64
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		slog.Error("failed to parse expense ID",
			append(requestContext(r),
				"error", err,
				"id_string", idStr,
			)...,
		)

		status := errors.HTTPStatus(errors.ErrInvalidInput)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": errors.ErrorMessage(errors.ErrInvalidInput),
		})
		return
	}

	expense, err := h.db.GetExpenseByID(id)
	if err != nil {
		slog.Error("failed to fetch expense by ID",
			append(requestContext(r),
				"error", err,
				"id", id,
			)...,
		)

		status := errors.HTTPStatus(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": errors.ErrorMessage(err),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"data":   expense,
	})

	slog.Info("expense fetched successfully",
		append(requestContext(r),
			"status", http.StatusOK,
			"id", id,
			"expense", expense.Expense,
		)...,
	)
}
