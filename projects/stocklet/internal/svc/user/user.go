package user

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/user/v1"
)

// Interface for the service
type UserService struct {
	pb.UnimplementedUserServiceServer

	StrCtrl StorageController
	EvtCtrl EventController
}

// Interface for database methods
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {
	// todo: implement
}

// Interface for event methods
// Allows flexibility to have seperate controllers for different messaging systems (e.g. Kafka, NATS, etc)
type EventController interface {
	PrepareConsumer(svc *UserService) messaging.EventConsumerController

	// todo: implement methods
	// DispatchCreatedEvent()
	// DispatchUpdatedEvent()
	// DispatchDeletedEvent()
}

// Create the user service
func NewUserService(cfg *ServiceConfig, strCtrl StorageController, evtCtrl EventController) *UserService {
	return &UserService{
		StrCtrl: strCtrl,
		EvtCtrl: evtCtrl,
	}
}

// todo: implement svc methods
func (svc UserService) TODO(ctx context.Context, req *pb.TODORequest) (*pb.TODOResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}