PROTOC_BASE := protoc -I ./schema/protobufs \
--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import

protoc: protoc-auth protoc-order protoc-payment protoc-product protoc-shipping protoc-user protoc-warehouse

protoc-auth:
	$(PROTOC_BASE) auth/v1/service.proto auth/v1/models.proto

protoc-order:
	$(PROTOC_BASE) order/v1/models.proto order/v1/service.proto order/v1/events.proto

protoc-payment:
	$(PROTOC_BASE) payment/v1/service.proto payment/v1/models.proto

protoc-product:
	$(PROTOC_BASE) product/v1/service.proto product/v1/models.proto

protoc-shipping:
	$(PROTOC_BASE) shipping/v1/service.proto shipping/v1/models.proto

protoc-user:
	$(PROTOC_BASE) user/v1/service.proto user/v1/models.proto

protoc-warehouse:
	$(PROTOC_BASE) warehouse/v1/service.proto warehouse/v1/models.proto