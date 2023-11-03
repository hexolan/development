package order

import (
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

func ProcessPlaceOrderEvent(event *pb.PlaceOrderEvent) {
	// Ignore events dispatched by the order service
	if event.Type == pb.PlaceOrderEvent_TYPE_UNSPECIFIED || event.Status == pb.PlaceOrderEvent_STATUS_UNSPECIFIED {
		return
	}

	// Mark the order as rejected if a failure status
	// was reported at any stage.
	if event.Status == pb.PlaceOrderEvent_STATUS_FAILURE {
		// todo
		return
	}
	
	// Otherwise,
	// If the event is from the last stage of the saga (shipping svc)
	// then mark the order as succesful.
	if event.Type == pb.PlaceOrderEvent_TYPE_SHIPPING {
		// todo: update order status and details
		// (append transaction id, etc to stored order)
	}
}