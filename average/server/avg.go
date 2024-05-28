package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/average/proto"
)

func (s *server) Average(strm pb.AverageService_AverageServer) error {
	var sum int32 = 0
	count := 0

	for {
		req, err := strm.Recv()

		if err == io.EOF {
			return strm.SendAndClose(&pb.AvgRes{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading inputs : %v\n", err)
		}

		fmt.Printf("Received number :%v\n", req.Number)
		sum += req.Number
		count++
	}
}
