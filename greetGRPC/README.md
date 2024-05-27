# gRPC 

# unary RPC implementation

cmd : make greetpb 

cmd : server >> go run server.go greet.go

cmd : client >> go run client.go greet.go