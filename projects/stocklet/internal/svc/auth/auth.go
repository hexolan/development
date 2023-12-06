package auth

import (
	"fmt"
	"context"
	"crypto/ecdsa"
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
	
	StrCtrl *StorageController

	privKey *ecdsa.PrivateKey
	pubJWKResponse *pb.GetJwksResponse
}

// Interface for database methods
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {

}

// Initialise the service
func NewAuthService(strCtrl *StorageController, privKey *ecdsa.PrivateKey) *AuthService {
	svc := &AuthService{
		StrCtrl: strCtrl,

		privKey: privKey,
		pubJWKResponse: nil,
	}

	return svc
}

// Underlying utility method for preparing the JWKs
func (svc *AuthService) getJwks() (*pb.GetJwksResponse, error) {
	// Assemble the public JWK
	jwk, err := jwk.FromRaw(svc.privKey.PublicKey)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "something went wrong parsing public key", err)
	}

	jwk.Set("use", "sig")  // denote use for signatures
	jwk.Set("alg", fmt.Sprintf("EC%v", svc.privKey.Curve.Params().BitSize))  // ease support for both EC256 and EC512

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
	log.Info().Msg("testing")
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

// todo: docs
func (svc *AuthService) SetPassword(ctx context.Context, req *pb.SetPasswordRequest) (*pb.SetPasswordResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

// todo: docs
func (svc *AuthService) DeleteUserData(ctx context.Context, req *pb.DeleteUserDataRequest) (*pb.DeleteUserDataResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}