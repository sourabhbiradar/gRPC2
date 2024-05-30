package main

import (
	"log"

	pb "github.com/gRPC/greetGRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr = "localhost:9321"

func main() {
	// SSL
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		cretFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(cretFile, "")

		if err != nil {
			log.Fatalf("Error loading CA cerificate :%v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
}
