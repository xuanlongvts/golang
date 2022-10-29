gen-proto:
	protoc --go_out=. hi/*.proto
	protoc --go-grpc_out=. hi/*.proto

run-server:
	go run server/main.go

run-client:
	go run client/main.go