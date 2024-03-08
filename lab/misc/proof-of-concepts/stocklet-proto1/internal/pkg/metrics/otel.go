package metrics

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

// Initiate the OpenTelemetry tracer provider
func InitTracerProvider(cfg *config.OtelConfig, svcName string) *sdktrace.TracerProvider {
	// create resoure
	resource := initTracerResource(svcName)
	
	// create trace exporter (to OTEL-collector)
	exporter := initTracerExporter(cfg.CollectorGrpc)

	// create trace provider with exporter
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),

		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp
}

// Establishes a connection to otel-collector over gRPC
func initTracerExporter(collectorEndpoint string) sdktrace.SpanExporter {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(collectorEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to start otlp grpc exporter")
	}

	return exporter
}

// Prepare a tracer resource to use with the tracing provider
func initTracerResource(svcName string) *sdkresource.Resource {
	ctx := context.Background()

	resource, err := sdkresource.New(
		ctx,
		sdkresource.WithAttributes(
			semconv.ServiceName(svcName),
		),
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create otel tracer resource")
	}

	return resource
}