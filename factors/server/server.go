package main

import (
	"log"
	"net"

	pb "github.com/sourabhbiradar/gRPC2/factors/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFactorServiceServer
}

var addr = "0.0.0.0:1111"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	log.Printf("Listening on %v\n", addr)

	srv := grpc.NewServer()
	pb.RegisterFactorServiceServer(srv, &server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve : %v", err)
	}
}
