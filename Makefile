.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: protos/generate
proto/generate:
	protoc --go_out=. --go-grpc_out=. proto/server.proto