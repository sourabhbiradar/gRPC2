package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sourabhbiradar/gRPC2/withDeadlines/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GreetDeadline(ctx context.Context, in *pb.GreetReq) (*pb.GreetResp, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Clinet canceled request")
			return nil, status.Error(codes.Canceled, "Clinet canceled request")
		}
		time.Sleep(time.Second)
	}
	return &pb.GreetResp{
		Reply: "Greetings " + in.GetName(),
	}, nil
}
