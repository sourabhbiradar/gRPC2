package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sourabhbiradar/gRPC2/clientStream/proto"
)

func sendNames(c pb.GreetServiceClient) {
	reqs := []*pb.GreetReq{
		{Name: "Abc"},
		{Name: "Xyz"},
		{Name: "Gtr"},
	}
	strm, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed calling LongGreet() : %v\n", err)
	}

	for _, req := range reqs {
		strm.Send(req)
		time.Sleep(time.Second)
	}

	res, err := strm.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving stream of response : %v\n", err)
	}

	log.Printf("LongGreet : %s\n", res)
}
