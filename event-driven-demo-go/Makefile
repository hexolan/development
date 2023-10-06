build-proto:
	protoc -I./schema/protos \
	--go_out=module=github.com/hexolan/event-driven-demo-go/protos:./internal/protos --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/event-driven-demo-go/protos:./internal/protos --go-grpc_opt=paths=import \
	order.proto events/order-placed-event.proto models/order.proto