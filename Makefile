proto:
	protoc --proto_path=. --go-grpc_out=module=github.com/vulpes-ferrilata/catan-service-proto:. --go_out=module=github.com/vulpes-ferrilata/catan-service-proto:. protobuf/*.proto protobuf/models/*.proto

.PHONY: proto