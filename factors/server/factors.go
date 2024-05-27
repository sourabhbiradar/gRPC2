package main

import (
	"log"

	pb "github.com/sourabhbiradar/gRPC2/factors/proto"
)

func (s *server) Factor(num *pb.FactorReq, strm pb.FactorService_FactorServer) error {
	log.Printf("Number :%v\n", num)

	number := num.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			strm.Send(&pb.FactorResp{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
