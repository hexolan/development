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

// Underlying utility method for preparing the JWKs
func (svc *AuthService) getJwks() (*pb.GetJwksResponse, error) {
	// Assemble the public JWK
	jwk, err := jwk.FromRaw(svc.cfg.ServiceOpts.PrivateKey.PublicKey)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "something went wrong parsing public key", err)
	}

	jwk.Set("use", "sig")  // denote use for signatures
	jwk.Set("alg", fmt.Sprintf("ES%v", svc.cfg.ServiceOpts.PrivateKey.Curve.Params().BitSize))  // dynamic support for ES256, ES384 and ES512

	// Attempt to marshal to protobuf
	// Convert the JWK to JSON
	jwkBytes, err := json.Marshal(jwk)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "something went wrong preparing the JWKs", err)
	}

	// Convert the JSON to Protobuf
	jwkPB := pb.ECPublicJWK{}
	err = protojson.Unmarshal(jwkBytes, &jwkPB)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "something went wrong preparing the JWKs", err)
	}

	// Succesfully assembled the JWK.
	// Store it for future requests
	svc.pubJWKResponse = &pb.GetJwksResponse{
		Keys: []*pb.ECPublicJWK{&jwkPB},
	}

	return svc.pubJWKResponse, nil
}

// Provide the ECDSA public key as part of a JSON Web Key set.
//
// This method is called by the API ingress for usage when validating JWT tokens.
func (svc *AuthService) GetJwks(ctx context.Context, req *pb.GetJwksRequest) (*pb.GetJwksResponse, error) {
	// Check if the JWK response has been prepared already
	if svc.pubJWKResponse != nil {
		return svc.pubJWKResponse, nil
	}

	// Prepare the JWK and return it.
	return svc.getJwks()
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