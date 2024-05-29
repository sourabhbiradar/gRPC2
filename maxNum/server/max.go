package main

import (
	"io"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/maxNum/proto"
)

func (s *server) MaxNum(strm pb.MaxNumService_MaxNumServer) error {

	var maximum int32 = 0

	for {
		req, err := strm.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving request stream : %v\n", err)
		}

		if num := req.Number; num > maximum {
			maximum = num
			strm.Send(&pb.NumResp{
				Max: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending response : %v\n", err)
			}
		}

	}
}
