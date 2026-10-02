package http

import (
	"errors"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.opentelemetry.io/otel/trace"
)

// requestLogger logs every request as structured JSON, correlated with the active trace
func requestLogger(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		// Run the error handler now so the logged status reflects the final response
		if err != nil {
			if handlerErr := c.App().ErrorHandler(c, err); handlerErr != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		status := c.Response().StatusCode()
		attrs := []slog.Attr{
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.String("route", c.FullPath()),
			slog.Int("status", status),
			slog.Float64("latency_ms", float64(time.Since(start).Microseconds())/1000),
			slog.String("ip", c.IP()),
			slog.String("request_id", requestid.FromContext(c)),
		}

		if sc := trace.SpanContextFromContext(c.Context()); sc.IsValid() {
			attrs = append(attrs,
				slog.String("trace_id", sc.TraceID().String()),
				slog.String("span_id", sc.SpanID().String()),
			)
		}

		level := slog.LevelInfo
		switch {
		case status >= fiber.StatusInternalServerError:
			level = slog.LevelError
		case status >= fiber.StatusBadRequest:
			level = slog.LevelWarn
		}

		logger.LogAttrs(c.Context(), level, "http request", attrs...)
		return nil
	}
}

// errorHandler returns a JSON error body and logs unexpected (5xx) failures
func errorHandler(logger *slog.Logger) fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := "internal server error"

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
			message = fiberErr.Message
		}

		if code >= fiber.StatusInternalServerError {
			logger.ErrorContext(c.Context(), "request failed",
				slog.String("path", c.Path()),
				slog.Any("error", err),
			)
		}

		return c.Status(code).JSON(fiber.Map{"error": message})
	}
}
