package main

import (
	"log"
	"net"

	pb "github.com/sourabhbiradar/gRPC2/clientStream/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

var addr = "0.0.0.0:2222"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen : %v\n", err)
	}
	log.Printf("Listening on %v\n", addr)

	srv := grpc.NewServer()
	pb.RegisterGreetServiceServer(srv, &server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}
}
