package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "John",
		Email: "john@snow.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not send gRPC message: %v", err)
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
		log.Fatalf("Could not send gRPC message: %v", err)
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

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "John",
			Email: "john@snow.com",
		},
		{
			Id:    "2",
			Name:  "John2",
			Email: "john2@snow.com",
		},
		{
			Id:    "3",
			Name:  "John3",
			Email: "john3@snow.com",
		},
		{
			Id:    "4",
			Name:  "John4",
			Email: "john4@snow.com",
		},
		{
			Id:    "5",
			Name:  "John5",
			Email: "john5@snow.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Could not send gRPC message stream: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Could not receive message: %v", err)
	}

	fmt.Println(res)
}
