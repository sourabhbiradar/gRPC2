package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/serverStream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:9999"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	greetMultiple(c)
}
