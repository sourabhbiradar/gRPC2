package main

import (
	"context"

	pb "github.com/gRPC/greetGRPC/proto"
)

func (s *server) Greet(ctx context.Context, req *pb.GreetReqt) (*pb.GreetResp, error) {
	return &pb.GreetResp{
		Reply: "Hello" + req.GetFirstName(),
	}, nil
}
