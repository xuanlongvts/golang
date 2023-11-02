package main

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"grpc-combine-rest/gapi"
	"grpc-combine-rest/pb"
	"net"
)

func main() {
	server, _ := gapi.NewServer()

	//grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	//grpcServer := grpc.NewServer(grpcLogger)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(err)
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
