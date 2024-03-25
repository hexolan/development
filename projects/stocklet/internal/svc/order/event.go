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

package order

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/hexolan/stocklet/internal/pkg/messaging"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func marshalEvent(event protoreflect.ProtoMessage) ([]byte, error) {
	wireEvent, err := proto.Marshal(event)
	if err != nil {
		return []byte{}, err
	}

	return wireEvent, nil
}

func PrepareCreatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Created_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_CREATED,
		Payload: order,
	}

	wireEvent, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvent, topic, nil
}

func PrepareUpdatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Updated_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_UPDATED,
		Payload: order,
	}

	wireEvent, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvent, topic, nil
}

func PrepareDeletedEvent(orderId string) ([]byte, string, error) {
	topic := messaging.Order_State_Deleted_Topic
	event := &pb.OrderStateEvent{
		Type: pb.OrderStateEvent_TYPE_DELETED,
		Payload: &pb.Order{
			Id: orderId,
		},
	}

	wireEvent, err := marshalEvent(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvent, topic, nil
}