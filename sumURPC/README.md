# sumURPC

Unary RPC 

`Sum` requests 2 numbers , adds 'em up & returns sum.

`NEW` : importing `msg.proto` into `sum.proto`

`cmd` : protoc -I=./proto --go_out=./proto --go-grpc_out=./proto proto/msg.proto proto/sum.proto

`-I=./proto`  : tells proto compiler where to look for proto files to import from