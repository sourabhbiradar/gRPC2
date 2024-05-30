package main

import (
	"log"
	"net"

	pb "github.com/sourabhbiradar/gRPC2/withDeadlines/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

var addr = "0.0.0.0:1111"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to create listener :%v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	srv := grpc.NewServer()
	pb.RegisterGreetServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v\n", err)
	}

}
