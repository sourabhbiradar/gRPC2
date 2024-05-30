.PHONY: greetpb clientrun serverrun

greetpb:
	protoc --go-grpc_out=./proto --go_out=./proto proto/greet.proto

clientrun:
	go run client/client.go client/greet.go

serverrun:
	go run server/server.go server/greet.go
