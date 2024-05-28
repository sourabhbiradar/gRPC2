package main

import (
	"log"
	"net"

	pb "github.com/sourabhbiradar/gRPC2/average/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAverageServiceServer
}

var addr = "0.0.0.0:5432"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Issue with Listener : %v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	srv := grpc.NewServer()
	pb.RegisterAverageServiceServer(srv, &server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Issue with serve :%v\n", err)
	}

}
