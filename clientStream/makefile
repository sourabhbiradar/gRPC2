.PHONY:	protoc run_client run_server

protoc:
	protoc -I ./proto --go-grpc_out=./proto --go_out=./proto proto/ser.proto proto/msg.proto

run_client:
	go run client/client.go client/xGreet.go

run_server:
	go run server/server.go server/xGreet.go