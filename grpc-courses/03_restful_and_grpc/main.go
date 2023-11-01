package main

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"grpc-combine-rest/gapi"
	"grpc-combine-rest/pb"
	"net"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func main() {
	server, _ := gapi.NewServer()

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterAuthServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
