package main

import (
	"log"
	"net"

	pb "github.com/gRPC/greetGRPC/proto"
	"google.golang.org/grpc"
)

type server struct { // to access all service procedures
	pb.UnimplementedGreetServiceServer
}

var addr = "0.0.0.0:9321" // localhost:9321

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on :%v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v\n", err)
	}

}
