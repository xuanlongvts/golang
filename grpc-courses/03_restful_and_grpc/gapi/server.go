package gapi

import "grpc-combine-rest/pb"

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}
