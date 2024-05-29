package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sourabhbiradar/gRPC2/maxNum/proto"
)

func findMax(c pb.MaxNumServiceClient) {
	strm, err := c.MaxNum(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream : %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 5, 3, 6, 2, 17}

		for _, number := range numbers {
			log.Printf("Sending number : %v\n", number)
			strm.Send(&pb.NumReq{
				Number: number,
			})
			time.Sleep(time.Second)
		}
		strm.CloseSend()
	}()

	go func() {
		for {
			res, err := strm.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving stream :%v\n", err)
				break
			}
			log.Printf("Received new maximum number :%v\n ", res.Max)
		}
		close(waitc)
	}()
	<-waitc
}
