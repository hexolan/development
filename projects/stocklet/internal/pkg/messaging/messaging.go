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

package messaging

type ConsumerController interface {
	Start()
	Stop()
}

// Topic Definitions
const (
	// Order Topics
	Order_State_Topic = "order.state"
	Order_State_Created_Topic = Order_State_Topic + ".created"
	Order_State_Updated_Topic = Order_State_Topic + ".updated"
	Order_State_Deleted_Topic = Order_State_Topic + ".deleted"

	// PlaceOrder Saga Topic
	Order_PlaceOrder_Topic = "order.placeorder"
	Order_PlaceOrder_Order_Topic = Order_PlaceOrder_Topic + ".order"
	Order_PlaceOrder_Warehouse_Topic = Order_PlaceOrder_Topic + ".warehouse"
	Order_PlaceOrder_Payment_Topic = Order_PlaceOrder_Topic + ".payment"
	Order_PlaceOrder_Shipping_Topic = Order_PlaceOrder_Topic + ".shipping"
)