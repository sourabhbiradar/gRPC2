package main

import (
	"log"
	"net"

	pb "github.com/sourabhbiradar/gRPC2/maxNum/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMaxNumServiceServer
}

var addr = "0.0.0.0:6666"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to create listener : %v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	srv := grpc.NewServer()
	pb.RegisterMaxNumServiceServer(srv, &server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v\n", err)
	}
}
