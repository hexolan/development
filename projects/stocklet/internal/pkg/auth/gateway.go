// Copyright (C) 2024 Declan Teevan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package auth

import (
	"context"
	"encoding/json"
	"encoding/base64"

	"google.golang.org/grpc/metadata"

	"github.com/hexolan/stocklet/internal/pkg/errors"
)

const JWTPayloadHeader string = "X-JWT-Payload"

// todo: sync claims with auth svc
type JWTClaims struct {
	Subject string `json:"sub"`
	IssuedAt int64 `json:"iat"`
	Expiry int64 `json:"exp"`
}

func GetGatewayUser(ctx context.Context) (*JWTClaims, error) {
	// check the gRPC request has metadata
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "no metadata provided")
	}

	// check the request came through gRPC-gateway
	if !isGatewayRequest(&md) {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "not a gateway request")
	}

	// check a token payload has been passed down
	// Envoy validates JWT tokens and provides a "X-JWT-Payload" header. (containing base64 JWT token claims)
	authorization := md.Get("jwt-payload")
	if len(authorization) != 1 {
		return nil, errors.NewServiceError(errors.ErrCodeUnauthorised, "invalid jwt")
	}

	// user is authenticated
	// attempt to decode the jwt-payload
	bytes, err := base64.StdEncoding.DecodeString(authorization[0])
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeUnauthorised, "malformed jwt", err)
	}

	var claims JWTClaims
	err = json.Unmarshal(bytes, &claims)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeUnauthorised, "malformed jwt", err)
	}


	return &claims, nil
}

func isGatewayRequest(md *metadata.MD) bool {
	if len(md.Get("from-gateway")) != 0 {
		return true
	}

	return false
}

func getTokenPayload(md *metadata.MD) (*JWTClaims, error) {
	// Envoy validates JWT tokens and provides a header containing 
	// valid auth token claims in base64 format.
	authorization := md.Get("jwt-payload")
	if len(authorization) != 1 {
		return nil, errors.NewServiceError(errors.ErrCodeUnauthorised, "auth token not provided")
	}

	// User is authenticated
	// Attempt to decode the jwt-payload
	bytes, err := base64.StdEncoding.DecodeString(authorization[0])
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeUnauthorised, "malformed jwt", err)
	}

	var claims JWTClaims
	err = json.Unmarshal(bytes, &claims)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeUnauthorised, "malformed jwt", err)
	}

	return &claims, nil
}