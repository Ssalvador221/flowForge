package http

import (
	"context"
	"errors"
	"log/slog"
	"time"

	fiberotel "github.com/gofiber/contrib/v3/otel"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

const (
	readTimeout     = 10 * time.Second
	writeTimeout    = 10 * time.Second
	idleTimeout     = 60 * time.Second
	shutdownTimeout = 10 * time.Second
)

type Server struct {
	app    *fiber.App
	addr   string
	logger *slog.Logger
}

// Init API HTTP Server
func NewServer(addr, serviceName string, logger *slog.Logger) *Server {
	app := fiber.New(fiber.Config{
		AppName:      serviceName,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		ErrorHandler: errorHandler(logger),
	})

	// Health checks are registered before tracing so probes don't generate spans
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())

	// Middlewares must be registered before the routes they should wrap
	app.Use(
		fiberotel.Middleware(
			fiberotel.WithTraceResponseHeader("X-Trace-ID"),
		),
		recover.New(recover.Config{EnableStackTrace: true}),
		requestid.New(),
		requestLogger(logger),
	)

	RegisterRoutes(app)

	return &Server{
		app:    app,
		addr:   addr,
		logger: logger,
	}
}

// Start blocks serving HTTP until ctx is cancelled, then shuts down gracefully
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("http server starting", slog.String("addr", s.addr))

	err := s.app.Listen(s.addr, fiber.ListenConfig{
		GracefulContext:       ctx,
		ShutdownTimeout:       shutdownTimeout,
		DisableStartupMessage: true,
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	s.logger.Info("http server stopped")
	return nil
}
