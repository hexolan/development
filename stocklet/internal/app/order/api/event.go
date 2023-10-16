package api

import (
	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/app/order"
)

func NewEventAPI(svc order.OrderRepository, kafkaConf config.KafkaConfig) {

}

func consumeTopicOrderCreate() {
	
}