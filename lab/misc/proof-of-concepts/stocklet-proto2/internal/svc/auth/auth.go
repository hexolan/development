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
	"fmt"
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

// Interface for the service
type AuthService struct {
	pb.UnimplementedAuthServiceServer

	cfg *ServiceConfig
	publicJWK *pb.ECPublicJWK
	
	store StorageController
}

// Interface for database methods
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {
	
}

// Create the auth service
func NewAuthService(cfg *ServiceConfig, store StorageController) *AuthService {
	publicJWK := preparePublicJwk(cfg)
	
	svc := &AuthService{
		store: store,
		publicJWK: publicJWK,
	}

	return svc
}

// Converts the ECDSA key to a publica JWK.
func preparePublicJwk(cfg *ServiceConfig) *pb.ECPublicJWK {
	// Assemble the public JWK
	jwk, err := jwk.FromRaw(cfg.ServiceOpts.PrivateKey.PublicKey)
	if err != nil {
		log.Panic().Err(err).Msg("something went wrong parsing public key from private key")
	}

	// denote use for signatures
	jwk.Set("use", "sig")

	// envoy includes support for ES256, ES384 and ES512
	alg := fmt.Sprintf("ES%v", cfg.ServiceOpts.PrivateKey.Curve.Params().BitSize)
	if alg != "ES256" && alg != "ES384" && alg != "ES512" {
		log.Panic().Err(err).Msg("unsupported bitsize for private key")
	}
	jwk.Set("alg", alg)

	// Convert the JWK to JSON
	jwkBytes, err := json.Marshal(jwk)
	if err != nil {
		log.Panic().Err(err).Msg("something went wrong preparing the public JWK (json marshal)")
	}

	// Unmarshal the JSON to Protobuf format
	jwkPB := pb.ECPublicJWK{}
	err = protojson.Unmarshal(jwkBytes, &jwkPB)
	if err != nil {
		log.Panic().Err(err).Msg("something went wrong preparing the public JWK (protonjson unmarshal)")
	}

	return &jwkPB
}

// Provide the JWK ECDSA public key as part of a JSON Web Key set.
// This method is called by the API ingress for usage when validating inbound JWT tokens.
func (svc *AuthService) GetJwks(ctx context.Context, req *pb.GetJwksRequest) (*pb.GetJwksResponse, error) {
	return &pb.GetJwksResponse{Keys: []*pb.ECPublicJWK{svc.publicJWK}}, nil
}

// todo: docs
func (svc *AuthService) LoginPassword(ctx context.Context, req *pb.LoginPasswordRequest) (*pb.LoginPasswordResponse, error) {
	// todo: implement
	log.Info().Msg("testing")
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

// todo: docs
func (svc *AuthService) SetPassword(ctx context.Context, req *pb.SetPasswordRequest) (*pb.SetPasswordResponse, error) {
	// todo: implement
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

// todo: docs
func (svc *AuthService) DeleteUserData(ctx context.Context, req *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	// todo: implement
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}