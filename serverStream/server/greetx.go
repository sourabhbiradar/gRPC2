package main

import (
	"fmt"

	pb "github.com/sourabhbiradar/gRPC2/serverStream/proto"
)

func (s *server) GreetMultiple(req *pb.GreetReq, strm pb.GreetService_GreetMultipleServer) error {
	for i := 0; i < 5; i++ {
		res := fmt.Sprintf("Hello %s number %d", req.GetFirstName(), i)

		strm.Send(&pb.GreetResp{
			Reply: res,
		})
	}
	return nil
}
