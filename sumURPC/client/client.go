package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/sumURPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:9898"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}

	c := pb.NewSumServiceClient(conn)

	getSum(c)
}
