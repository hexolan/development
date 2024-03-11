// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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