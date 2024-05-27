package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/factors/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:1111"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	c := pb.NewFactorServiceClient(conn)

	findFactors(c)
}
