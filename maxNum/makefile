# defining proto, client, server Dir 
P:=proto
C:=client
S:=server

# proto files
PROTO_FILES := $(P)/srvc.proto $(P)/msg.proto

protoc:
	protoc -I ./$(P) --go-grpc_out=./$(P) --go_out=./$(P) $(PROTO_FILES)

clientrun:
	go run $(C)/client.go $(C)/max.go

serverrun:
	go run $(S)/server.go $(S)/max.go

.PHONY: protoc clientrun serverrun