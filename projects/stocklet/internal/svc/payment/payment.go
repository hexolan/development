package payment

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/payment/v1"
)

// Interface for the service
type PaymentService struct {
	pb.UnimplementedPaymentServiceServer

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
	PrepareConsumer(svc *PaymentService) messaging.EventConsumerController

	// todo: implement methods
	// DispatchCreatedEvent()
	// DispatchUpdatedEvent()
	// DispatchDeletedEvent()
}

// Create the payment service
func NewPaymentService(cfg *ServiceConfig, strCtrl StorageController, evtCtrl EventController) *PaymentService {
	return &PaymentService{
		StrCtrl: strCtrl,
		EvtCtrl: evtCtrl,
	}
}

// todo: implement svc methods
func (svc PaymentService) TODO(ctx context.Context, req *pb.TODORequest) (*pb.TODOResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}