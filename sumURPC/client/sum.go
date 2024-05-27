package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sourabhbiradar/gRPC2/sumURPC/proto"
)

func getSum(c pb.SumServiceClient) {
	var a, b int32

	fmt.Print("Enter first number: ")
	fmt.Scanf("%d", &a)

	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &b)

	res, err := c.Sum(context.Background(), &pb.SumReqt{
		FirstNumber:  a,
		SecondNumber: b,
	})

	if err != nil {
		log.Fatalf("could not get sum : %v\n", err)
	}

	fmt.Println(res.Result)
}
