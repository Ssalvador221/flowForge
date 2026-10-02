package main

import (
	"FlowForge/internal/observability"
	"FlowForge/internal/transport/http"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const serviceName = "flowforge-api"

func main() {
	logger := observability.NewLogger()

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

	// HTTP API service init (blocks until SIGINT/SIGTERM)
	server := http.NewServer(":8080", serviceName, logger)
	if err := server.Start(ctx); err != nil {
		logger.Error("http server error", slog.Any("error", err))
	}
}
