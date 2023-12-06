package auth

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
)

type ServiceConfig struct {
	config.StandardConfig
}

func NewServiceConfig() (*ServiceConfig, error) {
	// todo: loading in ES256 private key
	// base64 encoded env variable?
	// decode base64 -> bytes -> jwx library (JWK)
	// gen JWK to serve from API route
	// use private key for generation of JWTs for authed
	// users

	// example of public key JWK for ES256:
	//
	// ensure use is set to signature
	/*
	{
		"kty": "EC",
		"use": "sig",
		"crv": "P-256",
		"x": "i2bzTXnapYy0SCe3nTwSroLV2hHmpjxPObCplD61c8c",
		"y": "2R9Uupcmx_4LaFYjAgIcPe8YMA0LuKxlGuuUKl8OKN4",
		"alg": "ES256"
	}
	*/

	cfg := ServiceConfig{}

	err := cfg.LoadStandardConfig()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
