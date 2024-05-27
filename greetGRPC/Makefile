.PHONY: greetpb  # declares greet as a phony target, meaning it is not a real file but just a label for a recipe to be executed when called.

greetpb:
protoc --go-grpc_out=./proto --go_out=./proto proto/greet.proto

# cmd : make greetpb