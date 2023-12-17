package api

import (
	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
)

func AttachSvcToConsumer(cfg *order.ServiceConfig, svc *order.OrderService) messaging.EventConsumerController {
	return svc.EvtCtrl.PrepareConsumer(svc)
}