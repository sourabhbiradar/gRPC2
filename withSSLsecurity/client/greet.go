package main

import (
	"context"
	"log"

	pb "github.com/gRPC/greetGRPC/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetReqt{
		FirstName: "Abc",
	})

	if err != nil {
		log.Fatalf("Failed to respond :%v\n", err)
	}

	log.Println(res.GetReply())
}
