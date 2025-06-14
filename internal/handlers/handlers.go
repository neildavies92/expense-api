package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bryx/expense-api/internal/database"
	"github.com/bryx/expense-api/internal/errors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	db *database.DB
}

func NewHandler(db *database.DB) *Handler {
	return &Handler{db: db}
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
			r.Get("/", h.handleGetExpenses)
		})
	})
	return r
}

func (h *Handler) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Handler) handleGetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.db.GetExpenses()
	if err != nil {
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
}

func (h *Handler) handlePostExpense(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Expense       string  `json:"expense"`
		ExpenseAmount float64 `json:"expense_amount"`
		DueDate       int     `json:"due_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		status := errors.HTTPStatus(errors.ErrInvalidInput)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": errors.ErrorMessage(errors.ErrInvalidInput),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   body,
	})
}
