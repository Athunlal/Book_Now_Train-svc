proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/train.proto
	
wire:
	go run github.com/google/wire/cmd/wire

run:
	go run cmd/api/main.go