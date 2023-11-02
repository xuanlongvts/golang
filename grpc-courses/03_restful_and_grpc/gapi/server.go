package gapi

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-combine-rest/pb"
	"time"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config string
}

func NewServer() (*Server, error) {
	ser := Server{
		UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
		config:                         "config test",
	}
	return &ser, nil
}

func (ser Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("---> CreateUserRequest: ", req.Email)

	fmt.Println("ser ", ser.config)

	return &pb.CreateUserResponse{User: &pb.User{
		UserName:          req.UserName,
		FullName:          req.FullName,
		Email:             req.Email,
		PasswordChangedAt: timestamppb.New(time.Now()),
		CreatedAt:         timestamppb.New(time.Now()),
	}}, nil
}

func (ser Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	fmt.Println("---> LoginUserRequest: ", req.UserName)
	fmt.Println("ser ", ser.config)

	return &pb.LoginUserResponse{
		User:                  nil,
		SessionId:             "SessionId ",
		AccessToken:           "",
		RefreshToken:          "",
		AccessTokenExpiresAt:  nil,
		RefreshTokenExpiresAt: nil,
	}, nil
}
