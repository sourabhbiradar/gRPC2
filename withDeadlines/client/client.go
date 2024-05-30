package main

import (
	"log"
	"time"

	pb "github.com/sourabhbiradar/gRPC2/withDeadlines/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:1111"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to grpc server : %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)

	// greetFunc(c, 5*time.Second)
	greetFunc(c, 2*time.Second) // deadline exceeded
}
