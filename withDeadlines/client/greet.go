package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sourabhbiradar/gRPC2/withDeadlines/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func greetFunc(c pb.GreetServiceClient, timeout time.Duration) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetReq{
		Name: "Jack",
	}

	res, err := c.GreetDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok { // ok == true then its grpc error
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline Exceeded")
				return
			} else {
				log.Fatalf("Unexpected gRPC error : %v\n", err)
			}
		} else {
			log.Fatalf("Non-gRPC error : %v\n", err)
		}
	}

	log.Printf("%s\n", res.Reply)
}
