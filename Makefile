PROTO_FILES_PATH=proto
PROTO_OUT=proto

run-proto:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/vehicles.proto