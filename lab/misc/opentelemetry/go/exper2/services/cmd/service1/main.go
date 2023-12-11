package main

import (
	"os"
	"errors"
	"context"
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"encoding/base64"

	"github.com/rs/zerolog/log"

	"null.hexolan.dev/dev/pkg/service1"
	"null.hexolan.dev/dev/pkg/otelutil"
	"null.hexolan.dev/dev/pkg/serveutil"
	pb "null.hexolan.dev/dev/protogen"
)

func main() {
	// todo: potentially attaching OTEL to Zerolog (look into this)
	// could use a hook like the Telegram example: https://betterstack.com/community/guides/logging/zerolog/
	//
	// https://opentelemetry.io/docs/instrumentation/
	// traces and metrics: implemented (and stable)
	// logging not yet implemented

	privateKey, err := loadPrivateKey()
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	// Init Open Telemetry
	var otelCollGrpcAddr string = "otel-collector:4317"
	otelResource := otelutil.InitTracerResource("service1")
	otelutil.InitTracerProvider(otelCollGrpcAddr, otelResource)

	// Create Service
	svc := service1.NewRpcService(privateKey)

	// GRPC Server
	grpcSvr := serveutil.NewGrpcServer()
	pb.RegisterAuthServiceServer(grpcSvr, svc)
	go serveutil.ServeGrpcServer(grpcSvr)

	// Gateway Server
	gwMux, gwOpts := serveutil.NewGrpcGateway()
	if err := pb.RegisterAuthServiceHandlerFromEndpoint(context.Background(), gwMux, "localhost:9090", gwOpts); err != nil {
		log.Panic().Err(err).Msg("failed to register gRPC to gateway server")
	}
	serveutil.ServeGrpcGateway(gwMux)
}

func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	// PEM private key file exposed as an environment variable encoded in base64 
	pkB64, exists := os.LookupEnv("AUTH_PRIVATE_KEY")
	if !exists || pkB64 == "" {
		return nil, errors.New("not found priv key")
	}

	// Decode from base64
	pkBytes, err := base64.StdEncoding.DecodeString(pkB64)
	if err != nil {
		return nil, errors.New("provided 'AUTH_PRIVATE_KEY' is not valid base64")
	}

	// Decode the PEM key
	pkBlock, _ := pem.Decode(pkBytes)
	if pkBlock == nil {
		return nil, errors.New("provided 'AUTH_PRIVATE_KEY' is not valid PEM format")
	}

	// Parse the block to a ecdsa.PrivateKey object
	privateKey, err := x509.ParseECPrivateKey(pkBlock.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse provided 'AUTH_PRIVATE_KEY' to ECDSA private key")
	}

	return privateKey, nil
}