package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/maxNum/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:6666"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect grpc server :%v\n", err)
	}
	defer conn.Close()

	c := pb.NewMaxNumServiceClient(conn)

	findMax(c)
}
