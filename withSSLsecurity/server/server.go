package main

import (
	"log"
	"net"

	pb "github.com/gRPC/greetGRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct { // to access all service procedures
	pb.UnimplementedGreetServiceServer
}

var addr = "localhost:9321"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on :%v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	// SSL related
	opts := []grpc.ServerOption{}
	tls := true // change as reqd.

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"                                       // NOT server.key
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile) // grpc.credentials

		if err != nil {
			log.Fatalf("failed to load SSL certificate & key : %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...) // opts := []grpc.ServerOption{}
	pb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v\n", err)
	}

}
