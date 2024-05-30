P:=proto
C:=client
S:=server

PROTO_FILES := $(P)/srvc.proto $(P)/msg.proto

protoc:
	protoc -I ./$(P) --go-grpc_out=./$(P) --go_out=./$(P) $(PROTO_FILES)

clientrun:
	go run $(C)/client.go $(C)/greet.go

serverrun:
	go run $(S)/server.go $(S)/greet.go

.PHONY: protoc clientrun serverrun
