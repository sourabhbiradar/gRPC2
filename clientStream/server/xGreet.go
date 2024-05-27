package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/clientStream/proto"
)

func (s *server) LongGreet(strm pb.GreetService_LongGreetServer) error {

	res := ""

	for {
		req, err := strm.Recv()
		if err == io.EOF {
			return strm.SendAndClose(&pb.GreetResp{
				Reply: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading request stream : %v\n", err)
		}
		res += fmt.Sprintf("Hello %s!", req.Name)

	}
}
