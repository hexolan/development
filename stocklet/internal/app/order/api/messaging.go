package api

import (
	"github.com/rs/zerolog/log"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/hexolan/stocklet/internal/app/order"
	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/messaging"
)


func NewMessagingAPI(svc order.OrderRepository, kcl *kgo.Client) {

}
