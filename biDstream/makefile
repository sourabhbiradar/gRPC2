.PHONY: protoc runserver runclient

protoc:
	protoc -I ./proto --go-grpc_out=./proto --go_out=./proto proto/srv.proto proto/msg.proto

runserver:
	go run server/server.go server/greetE1.go

runclient:
	go run client/client.go client/greetE1.go