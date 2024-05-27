package main

import (
	"context"

	pb "github.com/sourabhbiradar/gRPC2/sumURPC/proto"
)

func (s *server) Sum(ctx context.Context, req *pb.SumReqt) (*pb.SumResp, error) {
	return &pb.SumResp{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}
