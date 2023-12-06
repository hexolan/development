package auth

import (
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"encoding/base64"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

type ServiceConfig struct {
	config.StandardConfig

	PrivateKey *ecdsa.PrivateKey
}

// Load the ECDSA private key.
// Used for signing JWT tokens and validating at the API ingress.
func loadPrivateKey() (*ecdsa.PrivateKey, error) {
	// PEM private key file exposed as an environment variable encoded in base64 
	pkB64, err := config.RequireFromEnv("AUTH_PRIVATE_KEY") 
	if err != nil {
		return nil, err
	}

	// Decode from base64
	pkBytes, err := base64.StdEncoding.DecodeString(pkB64)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "provided 'AUTH_PRIVATE_KEY' is not valid base64", err)
	}

	// Decode the PEM key
	pkBlock, _ := pem.Decode(pkBytes)
	if pkBlock == nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "provided 'AUTH_PRIVATE_KEY' is not valid PEM format")
	}

	// Parse the block to a ecdsa.PrivateKey object
	privKey, err := x509.ParseECPrivateKey(pkBlock.Bytes)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to parse provided 'AUTH_PRIVATE_KEY' to ECDSA private key", err)
	}

	return privKey, nil
}

// Load the auth service configurations
func NewServiceConfig() (*ServiceConfig, error) {
	// Load the private key
	privKey, err := loadPrivateKey()
	if err != nil {
		return nil, err
	}

	// Initialise the configuration struct
	cfg := ServiceConfig{
		PrivateKey: privKey,
	}

	// Load the other standard config options
	err = cfg.LoadStandardConfig()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
