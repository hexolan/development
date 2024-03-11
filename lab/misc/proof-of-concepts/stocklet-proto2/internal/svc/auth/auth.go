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