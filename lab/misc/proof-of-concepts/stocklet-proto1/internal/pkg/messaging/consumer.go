package messaging

type EventConsumerController interface {
	Start()
	Stop()
}