package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/biDstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:4321"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error while connecting to %v\n", addr)
	}

	c := pb.NewGreetServiceClient(conn)

	greetAll(c)
}
