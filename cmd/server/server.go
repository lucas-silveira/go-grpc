package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not server: %v", err)
	}
}
