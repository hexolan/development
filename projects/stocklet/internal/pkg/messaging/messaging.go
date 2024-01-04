package messaging

// Consumer Interface
type EventConsumerController interface {
	Start()
	Stop()
}

// Topic Definitions
const (
	// RegEx
	Order_State_Catchall string = "order\\.state\\."
	Order_PlaceOrder_Catchall string = "order\\.placeorder\\."

	// Order Topics
	Order_State_Topic string = "order.state"
	Order_State_Created_Topic string = "order.state.created"
	Order_State_Updated_Topic string = "order.state.updated"
	Order_State_Deleted_Topic string = "order.state.deleted"

	// PlaceOrder Saga Topic
	Order_PlaceOrder_Topic string = "order.placeorder"
	Order_PlaceOrder_Order_Topic string = "order.placeorder.order"
	Order_PlaceOrder_Warehouse_Topic string = "order.placeorder.warehouse"
	Order_PlaceOrder_Payment_Topic string = "order.placeorder.payment"
	Order_PlaceOrder_Shipping_Topic string = "order.placeorder.shipping"
)