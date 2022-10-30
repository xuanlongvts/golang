gen-proto:
	protoc --go_out=. sum_proto/*.proto
	protoc --go-grpc_out=. sum_proto/*.proto

run-server:
	go run server/main.go

run-client:
	go run client/main.go