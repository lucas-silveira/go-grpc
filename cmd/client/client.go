package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/lucas-silveira/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "John",
		Email: "john@snow.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not send gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "John",
		Email: "john@snow.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not send gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()

		if err == io.EOF {
			break // finish the loop
		}

		if err != nil {
			log.Fatalf("Could not receive gRPC message: %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.User)
	}
}
