.PHONY: proto/generate
proto/generate:
	protoc --go_out=. --go-grpc_out=. proto/server.proto