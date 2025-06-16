package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/neildavies92/expense-api/config"
	"github.com/neildavies92/expense-api/internal/database"
	"github.com/neildavies92/expense-api/internal/errors"
	"github.com/neildavies92/expense-api/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load configuration", "error", errors.ErrorMessage(err))
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		slog.Error("failed to connect to database", "error", errors.ErrorMessage(err))
		os.Exit(1)
	}
	defer db.Close()

	handler := handlers.NewHandler(db)
	router := handlers.SetupRoutes(handler)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		slog.Info("starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", "error", errors.ErrorMessage(err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "error", errors.ErrorMessage(err))
		os.Exit(1)
	}
}
