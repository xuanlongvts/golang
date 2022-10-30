package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "mod-sum/sum_proto"
)

const (
	defaultNum1 = 1
	defaultNum2 = 2
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connnect to")
	num1 = flag.Int("num1", defaultNum1, "interger one")
	num2 = flag.Int("num2", defaultNum2, "interger two")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	defer conn.Close()

	cli := pb.NewCalculatorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := cli.Sum(ctx, &pb.SumRequest{
		Num1: int32(*num1),
		Num2: int32((*num2)),
	})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}
	log.Printf("Total: %v", result.GetResult())
}
