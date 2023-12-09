package otelutil

import (
	"context"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

// useful examples/demos:
// - https://github.com/open-telemetry/opentelemetry-demo/blob/59e33e5c115d6eb04b4dfeca75c19d782df66e7e/src/checkoutservice/main.go#L4
// - https://github.com/wavefrontHQ/opentelemetry-examples/blob/master/go-example/manual-instrumentation/main.go
func InitTracerProvider(otelCollGrpcAddr string, resource *sdkresource.Resource) *sdktrace.TracerProvider {
	exporter := InitTracerExporter(otelCollGrpcAddr)

	// create trace provider with exporter
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),

		// sdktrace.WithResource(InitTracerResource()),
		sdktrace.WithResource(resource), // the aim is for plug and play - so I can use this as a generic util for the services
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp
}

// establishes a connection to otel-collector over gRPC
// https://opentelemetry.io/docs/instrumentation/go/exporters/
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc#Option
func InitTracerExporter(otelCollGrpcAddr string) sdktrace.SpanExporter {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(otelCollGrpcAddr),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("new otlp trace grpc exporter failed")
	}

	return exporter
}

func InitTracerResource(svcName string) *sdkresource.Resource {
	ctx := context.Background()

	resource, err := sdkresource.New(
		ctx,
		sdkresource.WithAttributes(
			semconv.ServiceName(svcName),
			// attribute.String("application", "grpc"),
		),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create OT tracer resource")
	}

	return resource
}