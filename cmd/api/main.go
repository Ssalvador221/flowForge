package main

import (
	"FlowForge/internal/application/jobs"
	"FlowForge/internal/infrastructure/database"
	"FlowForge/internal/infrastructure/database/db"
	"FlowForge/internal/observability"
	"FlowForge/internal/transport/http"
	"FlowForge/internal/transport/http/handlers"
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

const serviceName = "flowforge-api"

func main() {
	logger := observability.NewLogger()

	// Loads .env for local runs 
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		logger.Error("failed to load .env", slog.Any("error", err))
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Observability Init service
	shutDown, err := observability.Init(ctx, serviceName)
	if err != nil {
		logger.Error("failed to init observability", slog.Any("error", err))
		os.Exit(1)
	}

	defer func() {
		// Fresh context: ctx is already cancelled at this point
		flushCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := shutDown(flushCtx); err != nil {
			logger.Error("failed to shutdown observability", slog.Any("error", err))
		}
	}()

	// Database pool init
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.Error("DATABASE_URL is not set")
		os.Exit(1)
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		logger.Error("failed to create db pool", slog.Any("error", err))
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		logger.Error("failed to connect to db", slog.Any("error", err))
		os.Exit(1)
	}

	jobRepo := database.NewJobRepository(db.New(pool))
	jobSVC := jobs.NewService(jobRepo, logger)

	// HTTP API service init
	server := http.NewServer(":8080", serviceName, logger, http.Handlers{
		Jobs: handlers.NewJobHandler(jobSVC),
	})

	if err := server.Start(ctx); err != nil {
		logger.Error("http server error", slog.Any("error", err))
	}
}
