.PHONY:	protoc clientrun serverrun

protoc:
	protoc -I ./proto --go-grpc_out=./proto --go_out=./proto proto/srv.proto proto/msg.proto

clientrun:
	go run client/client.go client/avg.go

serverrun:
	go run server/server.go server/avg.go