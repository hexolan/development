package auth

import (
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

// Interface for the service
type AuthService struct {
	pb.UnimplementedAuthServiceServer

	cfg *ServiceConfig
	pubJWKResponse *pb.GetJwksResponse
	
	StrCtrl *StorageController
}

// Interface for database methods
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {

}

// Create the auth service
func NewAuthService(cfg *ServiceConfig, strCtrl *StorageController) *AuthService {
	svc := &AuthService{
		cfg: cfg,
		pubJWKResponse: nil,

		StrCtrl: strCtrl,
	}

	return svc
}