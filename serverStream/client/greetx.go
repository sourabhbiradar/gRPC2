package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/serverStream/proto"
)

func greetMultiple(c pb.GreetServiceClient) {
	req := &pb.GreetReq{
		FirstName: "Abc",
	}

	strm, err := c.GreetMultiple(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetMultiple : %v", err)
	}

	for {
		msg, err := strm.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream : %v\n", err)
		}

		fmt.Printf("GreetMultiple : %v\n", msg.Reply)
	}
}
