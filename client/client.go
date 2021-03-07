package main

import (
	calculatorpb "calculator/proto"
	"context"
	"log"

	"io"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.Println("[INFO] Calculator client has started ...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[ERROR] Failed to dial with err: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doCalculateUnary(c)
	// doClientStreamingComputeAverage(c)
	// doServerStreamingDecomposition(c)
	doBiDiStreamingFindMaximum(c)
}

func doCalculateUnary(c calculatorpb.CalculatorServiceClient) {
	log.Println("[INFO] doCalculateUnary was invoked")

	numbers := calculatorpb.Numbers{
		First:  3,
		Second: 10,
	}

	req := &calculatorpb.CalculateRequest{
		Numbers: &numbers,
	}

	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("[ERROR] Failed to send req with: %v", err)
	}
	log.Printf("[INFO] Response: %v", res)
}

func doServerStreamingDecomposition(c calculatorpb.CalculatorServiceClient) {
	log.Println("[INFO] doServerStreamingDecomposition was invoked ...")

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

func doClientStreamingComputeAverage(c calculatorpb.CalculatorServiceClient) {
	log.Println("[INFO] doClientStreamingComputeAverage was invoked ...")

	requests := []*calculatorpb.ComputeAverageRequest{
		{Number: 5},
		{Number: 3},
		{Number: 12},
		{Number: 97},
	}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("[ERROR] Failed call Compute Average: %v", err)
	}

	for _, req := range requests {
		log.Printf("[INFO] Sending request: %v", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[ERROR] Could not receive response from Compute Average: %v", err)
	}
	log.Printf("[INFO] Response: %v", res)
}

func doBiDiStreamingFindMaximum(c calculatorpb.CalculatorServiceClient) {
	log.Println("[INFO] doBiDiStreamingFindMaximum was invoked ...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("[ERROR] Failed to create stream: %v", err)
		return
	}
	requests := []*calculatorpb.FindMaximumRequest{
		{Number: 5},
		{Number: 3},
		{Number: 12},
		{Number: 97},
	}

	//create a channel with empty struct to block
	waitc := make(chan struct{})

	//send data to the server
	go func() {
		for _, req := range requests {
			log.Printf("[INFO] Outgoing message: %v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	//read messages from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
			}

			if err != nil {
				log.Fatalf("[ERROR] Failed to read data from server: %v", err)
			}

			log.Printf("[INFO] Response from FindMaximum server: %v", res)
		}
	}()

	//wait for the channel to be blocked
	<-waitc
}
