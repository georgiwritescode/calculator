package main

import (
	calculatorpb "calculator/proto"
	"context"
	"log"

	"io"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello, I'm the calculate client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial with err: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doCalculateUnary(c)
	doServerStreamingDecomposition(c)
}

func doCalculateUnary(c calculatorpb.CalculatorServiceClient) {
	log.Println("Start an unary RPC...")

	numbers := calculatorpb.Numbers{
		First:  3,
		Second: 10,
	}

	req := &calculatorpb.CalculateRequest{
		Numbers: &numbers,
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send req with: %v", err)
	}
	log.Printf("Response: %v", res)
}

func doServerStreamingDecomposition(c calculatorpb.CalculatorServiceClient) {
	log.Println("[INFO] doServerStreamingDecomposition invoked ...")

	req := &calculatorpb.PrimeNumberRequest{
		PrimeNumber: &calculatorpb.PrimeNumber{
			PrimeNumber: 12,
		},
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] Failed to stream data: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Fatalln("[ERROR] End of stream reached")
			break
		}

		if err != nil {
			log.Fatalf("[ERROR] Failed to read stream: %v", err)
		}

		log.Printf("[INFO] Response from Decomposition: %v", msg.GetPrimeNumber())
	}
}
