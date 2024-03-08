package auth

import (
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"encoding/base64"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

// Auth Service Configuration
type ServiceConfig struct {
	// Core configuration
	Shared config.SharedConfig
	ServiceOpts ServiceConfigOpts

	// Dynamically loaded configuration
	Postgres config.PostgresConfig
}

// load the base service configuration
//
// This involves loading the service specific options (ServiceConfigOpts)
func NewServiceConfig() (*ServiceConfig, error) {
	cfg := ServiceConfig{}

	// load the shared config options
	if err := cfg.Shared.Load(); err != nil {
		return nil, err
	}

	// load the service config opts
	if err := cfg.ServiceOpts.Load(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Auth-service specific config options
type ServiceConfigOpts struct {
	// Env Var: "AUTH_PRIVATE_KEY"
	// to be provided in base64 format
	PrivateKey *ecdsa.PrivateKey
}

// Load the ServiceConfigOpts
//
// PrivateKey is loaded and decoded from the base64
// encoded PEM file exposed in the 'AUTH_PRIVATE_KEY'
// environment variable.
func (opts *ServiceConfigOpts) Load() error {
	// load the private key
	if err := opts.loadPrivateKey(); err != nil {
		return err
	}

	return nil
}

// Load the ECDSA private key.
//
// Used for signing JWT tokens.
// The public key is also served in JWK format, from this service,
// for use when validating the tokens at the API ingress.
func (opts *ServiceConfigOpts) loadPrivateKey() error {
	// PEM private key file exposed as an environment variable encoded in base64 
	opt, err := config.RequireFromEnv("AUTH_PRIVATE_KEY") 
	if err != nil {
		return err
	}

	// Decode from base64
	pkBytes, err := base64.StdEncoding.DecodeString(opt)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "the provided 'AUTH_PRIVATE_KEY' is not valid base64", err)
	}

	// Decode the PEM key
	pkBlock, _ := pem.Decode(pkBytes)
	if pkBlock == nil {
		return errors.NewServiceError(errors.ErrCodeService, "the provided 'AUTH_PRIVATE_KEY' is not valid PEM format")
	}

	// Parse the block to a ecdsa.PrivateKey object
	privKey, err := x509.ParseECPrivateKey(pkBlock.Bytes)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "failed to parse the provided 'AUTH_PRIVATE_KEY' to an EC private key", err)
	}

	opts.PrivateKey = privKey
	return nil
}