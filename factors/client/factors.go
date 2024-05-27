package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/factors/proto"
)

func findFactors(c pb.FactorServiceClient) {
	var n int64
	fmt.Print("Enter a number : ")
	fmt.Scanf("%d", &n)
	req := &pb.FactorReq{
		Number: n,
	}

	strm, err := c.Factor(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Factor() procedure : %v\n", err)
	}

	for {
		res, err := strm.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to read stream :%v\n ", err)
		}
		fmt.Printf("Factors : %v\n", res.Result)
	}
}
