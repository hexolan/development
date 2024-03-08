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