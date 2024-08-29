.PHONY: proto
proto:
	protoc \
	--go_out=./app/inventory \
	--go-grpc_out=./app/inventory \
	./app/protobuf/*.proto

.PHONY: run-service
run-service:
	cd ./app && go run ./cmd/main.go
