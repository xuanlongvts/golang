package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

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

func (ser *server) PrimeNumberDecomposition(req *pb.PndRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	K := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%K == 0 {
			N = N / K
			// send to client
			stream.Send(&pb.PndResponse{
				Result: K,
			})
			time.Sleep(time.Second)
		} else {
			K++
			log.Printf("K increase to %v", K)
		}
	}
	return nil
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
