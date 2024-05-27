package main

import (
	"log"
	"net"

	pb "github.com/sourabhbiradar/gRPC2/sumURPC/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSumServiceServer
}

var addr = "localhost:9898"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	srv := grpc.NewServer()
	pb.RegisterSumServiceServer(srv, &server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on : %v\n", err)
	}
}
