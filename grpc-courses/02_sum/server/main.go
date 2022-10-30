package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "mod-sum/sum_proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (ser *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Received, num1: %v, num2: %v", in.GetNum1(), in.GetNum2())
	return &pb.SumResponse{
		Result: in.GetNum1() + in.GetNum2(),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	ser := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(ser, &server{})
	log.Println("server listening at %v", lis.Addr())

	if err := ser.Serve(lis); err != nil {
		log.Fatalf("Faile to serve: %v", err)
	}
}
