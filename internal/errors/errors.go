package errors

import (
	"errors"
	"net/http"
)

var (
	// Database errors
	ErrNotFound          = errors.New("resource not found")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrInvalidInput      = errors.New("invalid input")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
)

// HTTPStatus returns the appropriate HTTP status code for the given error
func HTTPStatus(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrDuplicateUsername):
		return http.StatusConflict
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// ErrorMessage returns a user-friendly error message for the given error
func ErrorMessage(err error) string {
	switch {
	case errors.Is(err, ErrNotFound):
		return "Resource not found"
	case errors.Is(err, ErrDuplicateUsername):
		return "Username already exists"
	case errors.Is(err, ErrInvalidInput):
		return "Invalid input"
	case errors.Is(err, ErrUnauthorized):
		return "Unauthorized"
	case errors.Is(err, ErrForbidden):
		return "Forbidden"
	default:
		return "Internal server error"
	}
}
