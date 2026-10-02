package observability

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

type ShutdownFunc func(context.Context) error

func Init(
	ctx context.Context,
	serviceName string,
) (ShutdownFunc, error) {
	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, err
	}

	tracerProvider, err := newTracerProvider(ctx, res)
	if err != nil {
		return nil, err
	}

	meterProvider, err := newMeterProvider(ctx, res)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)

	shutDown := func(ctx context.Context) error {
		return errors.Join(
			tracerProvider.Shutdown(ctx),
			meterProvider.Shutdown(ctx),
		)
	}

	_ = res

	return shutDown, nil
}
