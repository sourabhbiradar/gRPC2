package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/clientStream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:2222"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connected to server : %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	sendNames(c)
}
