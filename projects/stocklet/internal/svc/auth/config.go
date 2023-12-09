package auth

import (
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"encoding/base64"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

// Service Configuration
type ServiceConfig struct {
	// core configuration
	Shared *config.SharedConfig
	ServiceOpts *ServiceConfigOpts

	// dynamically loaded configuration
	Postgres *config.PostgresConfig
}

// Load the core auth service opts
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

// Service specific cfg options
type ServiceConfigOpts struct {
	PrivateKey *ecdsa.PrivateKey
}

// Load the ECDSA private key.
// Used for signing JWT tokens and validating at the API ingress.
func (opts *ServiceConfigOpts) loadPrivateKey() error {
	// PEM private key file exposed as an environment variable encoded in base64 
	pkB64, err := config.RequireFromEnv("AUTH_PRIVATE_KEY") 
	if err != nil {
		return err
	}

	// Decode from base64
	pkBytes, err := base64.StdEncoding.DecodeString(pkB64)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "provided 'AUTH_PRIVATE_KEY' is not valid base64", err)
	}

	// Decode the PEM key
	pkBlock, _ := pem.Decode(pkBytes)
	if pkBlock == nil {
		return errors.NewServiceError(errors.ErrCodeService, "provided 'AUTH_PRIVATE_KEY' is not valid PEM format")
	}

	// Parse the block to a ecdsa.PrivateKey object
	privKey, err := x509.ParseECPrivateKey(pkBlock.Bytes)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "failed to parse provided 'AUTH_PRIVATE_KEY' to ECDSA private key", err)
	}

	opts.PrivateKey = privKey
	return nil
}

// Load the service specific configuration options
func (opts *ServiceConfigOpts) Load() error {
	// load the private key
	if err := opts.loadPrivateKey(); err != nil {
		return err
	}

	return nil
}