package main

import (
	"context"
	"fmt"
	"github.com/dumunari/go-grpc-snippets/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddServerStream(client)
}

func AddServerStream(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Spike",
		Email: "spikethepirate@instagram.com",
	}
	responseStream, err := client.AddUserServerStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status: ", stream.Status)
	}
}
