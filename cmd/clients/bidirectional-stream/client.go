package main

import (
	"context"
	"fmt"
	"github.com/dumunari/go-grpc-snippets/pb"
	"google.golang.org/grpc"
	"io"
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
	AddUserStream(client)
}

func AddUserStream(client pb.UserServiceClient) {
	stream, err := client.AddUsersBidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

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

	wait := make(chan int)

	go func(){
		for _, req := range reqs{
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func(){
		for{
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}
			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
