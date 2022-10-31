package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "mod-sum/sum_proto"
)

const (
	defaultNum1      = 1
	defaultNum2      = 2
	defaultNumPnd    = 120
	defaultNumSquare = 4
)

var (
	addr      = flag.String("addr", "localhost:50051", "the address to connnect to")
	num1      = flag.Int("num1", defaultNum1, "interger one")
	num2      = flag.Int("num2", defaultNum2, "interger two")
	numPnd    = flag.Int("numPnd", defaultNumPnd, "interger Prime Number")
	numSquare = flag.Int("numSquare", defaultNumSquare, "interger number square")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	cli1 := pb.NewCalculatorServiceClient(conn)
	total(cli1)

	fmt.Println("---------------------------")
	cli2 := pb.NewCalculatorServiceClient(conn)
	callPnd(cli2)

	fmt.Println("---------------------------")
	cli3 := pb.NewCalculatorServiceClient(conn)
	callAverage(cli3)

	fmt.Println("---------------------------")
	cli4 := pb.NewCalculatorServiceClient(conn)
	callFindMax(cli4)

	fmt.Println("---------------------------")
	cli5 := pb.NewCalculatorServiceClient(conn)
	callSquareRoot(cli5, int32(*numSquare))
}

func total(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.Sum(ctx, &pb.SumRequest{
		Num1: int32(*num1),
		Num2: int32((*num2)),
	})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}
	log.Printf("Total: %v", result.GetResult())
}

func callPnd(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.PrimeNumberDecomposition(ctx, &pb.PndRequest{
		Number: int32(*numPnd),
	})
	if err != nil {
		log.Fatalf("call Prime Number Decomposition %v", err)
	}

	for {
		result, err := stream.Recv()
		// EOF = End Of File (Server had sent done)
		if err == io.EOF {
			log.Println("Server finish streaming")
			return
		}

		log.Printf("Prime number %v", result.GetResult())
	}
}

func callAverage(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.Average(ctx)
	if err != nil {
		log.Fatalf("Call average err %v", err)
	}
	listReq := []pb.AverageRequest{
		pb.AverageRequest{
			Num: 5,
		},
		pb.AverageRequest{
			Num: 2,
		},
		pb.AverageRequest{
			Num: 3.5,
		},
		pb.AverageRequest{
			Num: 1.5,
		},
	}
	for _, val := range listReq {
		err := stream.Send(&val)
		if err != nil {
			log.Fatalf("send avarage request err %v", err)
		}
		time.Sleep(time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Receive average response err %v", err)
	}
	log.Printf("Average response total: %+v", resp)
}

func callFindMax(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.FindMax(ctx)
	if err != nil {
		log.Fatalf("Call FindMax err %v", err)
	}

	waitC := make(chan struct{})
	go func() {
		listReq := []pb.FindMaxRequest{
			pb.FindMaxRequest{
				Num: 5,
			},
			pb.FindMaxRequest{
				Num: 15,
			},
			pb.FindMaxRequest{
				Num: 12,
			},
			pb.FindMaxRequest{
				Num: 3,
			},
			pb.FindMaxRequest{
				Num: 35,
			},
		}
		for _, val := range listReq {
			err := stream.Send(&val)
			if err != nil {
				log.Fatalf("send find max request err %v", err)
				break
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("ending find max api ...")
				break
			}
			if err != nil {
				log.Fatalf("recv find max err %v", err)
				break
			}
			log.Printf("max: %v\n", resp.GetMax())
		}
		close(waitC)
	}()

	<-waitC
}

func callSquareRoot(client pb.CalculatorServiceClient, num int32) {
	resp, err := client.Square(context.Background(), &pb.SquareRequest{
		Num: num,
	})

	if err != nil {
		fmt.Println("Call square root api err: ", err)

		if errStatus, ok := status.FromError(err); ok {
			fmt.Println("err msg: ", errStatus.Message())
			fmt.Println("err code: ", errStatus.Code())

			if errStatus.Code() == codes.InvalidArgument {
				fmt.Println("InvalidArgument num ", num)
				return
			}
		}
	}
	log.Printf("Square root response: %v\n\n", resp.GetSquareRoot())
}
