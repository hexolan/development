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

package warehouse

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/warehouse/v1"
	eventpb "github.com/hexolan/stocklet/internal/pkg/protogen/events/v1"
	commonpb "github.com/hexolan/stocklet/internal/pkg/protogen/common/v1"
)

// Interface for the service
type WarehouseService struct {
	pb.UnimplementedWarehouseServiceServer
	
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

	Attach(svc pb.WarehouseServiceServer)
}

// Create the shipping service
func NewWarehouseService(cfg *ServiceConfig, store StorageController) *WarehouseService {
	// Initialise the protobuf validator
	pbVal, err := protovalidate.New()
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialise protobuf validator")
	}

	// Initialise the service
	return &WarehouseService{
		store: store,
		pbVal: pbVal,
	}
}

func (svc WarehouseService) ServiceInfo(ctx context.Context, req *commonpb.ServiceInfoRequest) (*commonpb.ServiceInfoResponse, error) {
	return &commonpb.ServiceInfoResponse{
		Name: "warehouse",
		Source: "https://github.com/hexolan/stocklet",
		SourceLicense: "AGPL-3.0",
	}, nil
}

func (svc WarehouseService) ViewProductStock(ctx context.Context, req *pb.ViewProductStockRequest) (*pb.ViewProductStockResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc WarehouseService) ViewReservation(ctx context.Context, req *pb.ViewReservationRequest) (*pb.ViewReservationResponse, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc WarehouseService) ProcessOrderPendingEvent(ctx context.Context, req *eventpb.OrderPendingEvent) (*emptypb.Empty, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc WarehouseService) ProcessShipmentAllocationEvent(ctx context.Context, req *eventpb.ShipmentAllocationEvent) (*emptypb.Empty, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}

func (svc WarehouseService) ProcessPaymentProcessedEvent(ctx context.Context, req *eventpb.PaymentProcessedEvent) (*emptypb.Empty, error) {
	return nil, errors.NewServiceError(errors.ErrCodeService, "todo")
}