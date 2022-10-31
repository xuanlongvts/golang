package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	fmt.Println("Unary ---------------------------")
	log.Printf("Received, num1: %v, num2: %v", in.GetNum1(), in.GetNum2())
	return &pb.SumResponse{
		Result: in.GetNum1() + in.GetNum2(),
	}, nil
}

func (ser *server) PrimeNumberDecomposition(req *pb.PndRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Println("Server streaming ---------------------------")
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

func (ser *server) Average(stream pb.CalculatorService_AverageServer) error {
	fmt.Println("Client streaming ---------------------------")
	var total float32
	var count int

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &pb.AverageResponse{
				Result: total / float32(count),
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatalf("Erro while Recv Average %v", err)
			return err
		}
		log.Printf("receive request %+v \n", req)
		total += req.GetNum()
		count++
	}

	return nil
}

func (ser *server) FindMax(stream pb.CalculatorService_FindMaxServer) error {
	fmt.Println("Bi Directional streaming ---------------------------")
	var max int32 = int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF...")
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receive FindMax %v", err)
			return err
		}
		num := req.GetNum()
		log.Printf("Recv num %v\n", num)
		if num > max {
			max = num
		}
		err = stream.Send(&pb.FindMaxResponse{
			Max: max,
		})
		if err != nil {
			log.Fatalf("send max err %v", err)
			return err
		}
	}
	return nil
}

func (ser *server) Square(ctx context.Context, req *pb.SquareRequest) (*pb.SquareResponse, error) {
	fmt.Println("Square called ---------------------------")
	num := req.GetNum()

	if num < 0 {
		log.Printf("number <= 0, num = %v ---> InvalidArgument ", num)
		return nil, status.Errorf(codes.InvalidArgument, "Expect number > 0, request number was: %v", num)
	}

	return &pb.SquareResponse{
		SquareRoot: math.Sqrt(float64(num)),
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
