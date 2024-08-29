.PHONY: protoc
protoc:
	protoc \
	--go_out=./app/service \
	--go-grpc_out=./app/service \
	./app/protobuf/*.proto

.PHONY: run-service
run-service:
	cd ./app && go run ./cmd/main.go
