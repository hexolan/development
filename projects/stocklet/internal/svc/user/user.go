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

package user

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/bufbuild/protovalidate-go"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/user/v1"
	commonpb "github.com/hexolan/stocklet/internal/pkg/protogen/common/v1"
)

// Interface for the service
type UserService struct {
	pb.UnimplementedUserServiceServer
	
	store StorageController
	pbVal *protovalidate.Validator
}

// Interface for database methods
// Flexibility for implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {

}

// Interface for event consumption
// Flexibility for seperate controllers for different messaging systems (e.g. Kafka, NATS, etc)
type ConsumerController interface {
	messaging.ConsumerController

	Attach(svc pb.UserServiceServer)
}

// Create the shipping service
func NewUserService(cfg *ServiceConfig, store StorageController) *UserService {
	// Initialise the protobuf validator
	pbVal, err := protovalidate.New()
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialise protobuf validator")
	}

	// Initialise the service
	return &UserService{
		store: store,
		pbVal: pbVal,
	}
}

func (svc UserService) ServiceInfo(ctx context.Context, req *commonpb.ServiceInfoRequest) (*commonpb.ServiceInfoResponse, error) {
	return &commonpb.ServiceInfoResponse{
		Name: "user",
		Source: "https://github.com/hexolan/stocklet",
		SourceLicense: "AGPL-3.0",
	}, nil
}

func (svc UserService) ViewUser(ctx context.Context, req *pb.ViewUserRequest) (*pb.ViewUserResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}