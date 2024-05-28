package main

import (
	"io"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/biDstream/proto"
)

func (s *server) GreetEveryone(strm pb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := strm.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream : %v\n", err)
		}

		res := "Hello " + req.Name + " !"

		err = strm.Send(&pb.GreetResp{
			Reply: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data : %v\n", err)
		}

	}
}
