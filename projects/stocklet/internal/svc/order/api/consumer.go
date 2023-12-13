package api

import (
	"github.com/hexolan/stocklet/internal/svc/order"
)

func AttachSvcToConsumer(cfg *order.ServiceConfig, svc *order.OrderService) order.EventConsumerController {
	return svc.EvtCtrl.PrepareConsumer(svc)
}