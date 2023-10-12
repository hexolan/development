protoc-auth:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	auth/v1/service.proto auth/v1/models.proto

protoc-order:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	order/v1/service.proto order/v1/models.proto

protoc-payment:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	payment/v1/service.proto payment/v1/models.proto

protoc-product:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	product/v1/service.proto product/v1/models.proto

protoc-shipping:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	shipping/v1/service.proto shipping/v1/models.proto

protoc-user:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	user/v1/service.proto user/v1/models.proto

protoc-warehouse:
	protoc -I ./schema/protobufs \
	--go_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/stocklet/internal/pkg/protobuf:./internal/pkg/protobuf --go-grpc_opt=paths=import \
	warehouse/v1/service.proto warehouse/v1/models.proto

protoc: protoc-auth protoc-order protoc-payment protoc-product protoc-shipping protoc-user protoc-warehouse