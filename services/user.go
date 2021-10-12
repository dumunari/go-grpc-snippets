package services

import (
	"context"
	"fmt"
	"github.com/dumunari/go-grpc-snippets/pb"
	"io"
	"log"
)

//type UserServiceServer interface {
//  AddUserUnary(context.Context, *User) (*User, error)
//  AddUserServerStream(*User, UserService_AddUserServerStreamServer) error
//  AddUsersClientStream(UserService_AddUsersClientStreamServer) error
//  AddUsersBidirectionalStream(UserService_AddUsersBidirectionalStreamServer) error
//  mustEmbedUnimplementedUserServiceServer()
//}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUserUnary(ctx context.Context, req *pb.User) (*pb.User, error) {
	// insert

	fmt.Println(req.GetName())

	return &pb.User{
		Id:    "1705",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserServerStream(req *pb.User, stream pb.UserService_AddUserServerStreamServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "1705",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "1705",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	return nil
}

func (*UserService) AddUsersClientStream(stream pb.UserService_AddUsersClientStreamServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding: ", req.GetName())
	}
}

func (*UserService) AddUsersBidirectionalStream(stream pb.UserService_AddUsersBidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})
		if err != nil {
			log.Fatalf("Error sending stream to the client: %v", err)
		}
	}
}
