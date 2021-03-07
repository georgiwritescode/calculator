package main

import (
	calculatorpb "calculator/proto"
	"context"
	"log"

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
