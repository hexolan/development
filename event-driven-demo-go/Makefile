build:
	protoc --go_out=. --go_opt=MOrderService.proto=./builds --go-grpc_out=. --go-grpc_opt=MOrderService.proto=./builds OrderService.proto

build2:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    order.proto models/order.proto events/order.proto

build3:
	protoc --go_out=./protobuild --go_opt=paths=source_relative \
	--go-grpc_out=./protobuild --go-grpc_opt=paths=source_relative \
	order.proto events/order.proto models/order.proto

build4:
	protoc --go_out=./protobuild \
	--go-grpc_out=./protobuild \
	order.proto events/order.proto models/order.proto

build5:
	protoc --go_out=./protobuild --go_opt=paths=import \
	--go-grpc_out=./protobuild --go-grpc_opt=paths=import \
	order.proto events/order.proto models/order.proto

build6:
	protoc --go_out=module=github.com/hexolan/test:./protobuild --go_opt=paths=import \
	--go-grpc_out=module=github.com/hexolan/test:./protobuild --go-grpc_opt=paths=import \
	order.proto events/order.proto models/order.proto