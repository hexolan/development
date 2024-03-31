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

func marshalEvent(event protoreflect.ProtoMessage, topic string) ([]byte, string, error) {
	wireEvent, err := proto.Marshal(event)
	if err != nil {
		return []byte{}, "", err
	}

	return wireEvent, topic, nil
}

func PrepareOrderCreatedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Created_Topic
	event := &pb.OrderCreatedEvent{
		Revision: 1,

		OrderId: order.Id,
		CustomerId: order.CustomerId,
		ItemQuantities: order.Items,
	}

	return marshalEvent(event, topic)
}

func PrepareOrderPlacedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Pending_Topic
	event := &pb.OrderPlacedEvent{
		Revision: 1,

		OrderId: order.Id,
		CustomerId: order.CustomerId,
		ItemQuantities: order.Items,
	}

	return marshalEvent(event, topic)
}

func PrepareOrderRejectedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Rejected_Topic
	event := &pb.OrderRejectedEvent{
		Revision: 1,

		OrderId: order.Id,
		TransactionId: nil, // todo:
		ShippingId: nil, // todo:
	}

	return marshalEvent(event, topic)
}

func PrepareOrderApprovedEvent(order *pb.Order) ([]byte, string, error) {
	topic := messaging.Order_State_Approved_Topic
	event := &pb.OrderApprovedEvent{
		Revision: 1,

		OrderId: order.Id,
		TransactionId: "", // todo:
		ShippingId: "", // todo:
	}

	return marshalEvent(event, topic)
}
