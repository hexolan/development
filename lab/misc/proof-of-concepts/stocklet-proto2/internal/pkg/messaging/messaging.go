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

package messaging

// Consumer Interface
type EventConsumerController interface {
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