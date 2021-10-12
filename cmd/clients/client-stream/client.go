package main

import (
	"context"
	"fmt"
	"github.com/dumunari/go-grpc-snippets/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUsersClientStream(client)
}

func AddUsersClientStream(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "0",
			Name:  "Spike",
			Email: "spikethepirate@instagram.com",
		},
		{
			Id:    "1",
			Name:  "Spike 1",
			Email: "spikethepirate1@instagram.com",
		},
		{
			Id:    "2",
			Name:  "Spike 2",
			Email: "spikethepirate2@instagram.com",
		},
	}

	stream, err := client.AddUsersClientStream(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}
