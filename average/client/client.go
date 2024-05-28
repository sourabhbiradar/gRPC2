package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/average/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:5432"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connection Error : %v\n", err)
	}

	c := pb.NewAverageServiceClient(conn)

	findAvg(c)
}
