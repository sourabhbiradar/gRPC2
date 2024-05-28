package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sourabhbiradar/gRPC2/biDstream/proto"
)

func greetAll(c pb.GreetServiceClient) {
	strm, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream :%v\n", err)
	}

	names := []*pb.GreetReq{
		{Name: "Raju"},
		{Name: "Shyam"},
		{Name: "Babu Rao"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range names {
			log.Printf("Send Name : %v\n", req)
			strm.Send(req)
			time.Sleep(time.Second)
		}
		strm.CloseSend() // after sending all close
	}()

	go func() {
		for {
			res, err := strm.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving : %v\n", err)
				break
			}
			log.Printf("Received : %v\n", res.Reply)
		}
		close(waitc)
	}()
	<-waitc
}
