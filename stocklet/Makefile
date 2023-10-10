build-proto:
	protoc -I./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/pkg/protos:./pkg/protos --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/pkg/protos:./pkg/protos --go-grpc_opt=paths=import \
	services/order.proto events/order-placed-event.proto models/order.proto