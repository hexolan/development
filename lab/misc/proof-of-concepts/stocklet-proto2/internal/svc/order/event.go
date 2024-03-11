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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func marshalEvent(event protoreflect.ProtoMessage) ([]byte, error) {
	wireEvt, err := proto.Marshal(event)
	if err != nil {
		return []byte{}, err
	}

	return wireEvt, nil
}

func PrepareCreatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Created_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_CREATED,
		Payload: order,
	}

	wireEvt, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvt, topic, nil
}

func PrepareUpdatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Updated_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_UPDATED,
		Payload: order,
	}

	wireEvt, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvt, topic, nil
}

func PrepareDeletedEvent(orderId string) ([]byte, string, error) {
	topic := messaging.Order_State_Deleted_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_DELETED,
		Payload: &pb.Order{
			Id: orderId,
		},
	}

	wireEvt, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvt, topic, nil
}