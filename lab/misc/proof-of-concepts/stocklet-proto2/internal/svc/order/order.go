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

package order

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

// Interface for the service
type OrderService struct {
	pb.UnimplementedOrderServiceServer
	
	pbVal *protovalidate.Validator

	StrCtrl StorageController
	EvtCtrl EventController
}

// Interface for database methods
//
// Allows implementing seperate controllers for different databases (e.g. Postgres, MongoDB, etc)
type StorageController interface {
	GetOrder(ctx context.Context, orderId string) (*pb.Order, error)
	CreateOrder(ctx context.Context, orderObj *pb.Order) (*pb.Order, error)
	UpdateOrder(ctx context.Context, orderId string, orderObj *pb.Order, mask *fieldmaskpb.FieldMask) (*pb.Order, error)
	DeleteOrder(ctx context.Context, orderId string) error
	
	GetOrderItems(ctx context.Context, orderId string) (*map[string]int32, error)
	SetOrderItems(ctx context.Context, orderId string, itemQuantities map[string]int32) (*map[string]int32, error)
	SetOrderItem(ctx context.Context, orderId string, itemId string, quantity int32) (*map[string]int32, error)
	DeleteOrderItem(ctx context.Context, orderId string, itemId string) (*map[string]int32, error)

	GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error)
}

// Interface for the messaging methods
//
// Allows flexibility to have seperate controllers for different messaging systems (e.g. Kafka, NATS, etc)
type EventController interface {
	PrepareConsumer(svc *OrderService) messaging.EventConsumerController
}