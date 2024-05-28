package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/average/proto"
)

func findAvg(c pb.AverageServiceClient) {
	strm, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error opening stream : %v\n", err)
	}

	numbers := []int32{3, 7, 8, 20, 5}

	for _, number := range numbers {
		strm.Send(&pb.AvgReq{
			Number: number,
		})
	}

	res, err := strm.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response : %v\n", err)
	}

	fmt.Printf("Average of numbers : %f", res.Result)
}
