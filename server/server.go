package main

import (
	calculatorpb "calculator/proto"
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	log.Printf("[INFO] Calculate func was invoked with %v", req)
	a := req.Numbers.GetFirst()
	b := req.Numbers.GetSecond()

	sum := a + b
	res := &calculatorpb.CalculateResponse{
		Sum: sum,
	}

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("[INFO] PrimeNumberDecomposition function invoked")
	number := req.PrimeNumber.GetPrimeNumber()
	var divisor = int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberResponse{
				PrimeNumber: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			log.Printf("[INFO] Divisor has increased to: %v", divisor)
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Println("[INFO] ComputeAverage function invoked")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := float32(sum) / float32(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{Average: avg})
		}
		if err != nil {
			log.Fatalf("[ERROR] Cannot read client stream: %v", err)
		}

		sum += int(req.GetNumber())
		count++
	}
}

func main() {
	log.Println("[INFO] Calculator server has started ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("[ERROR] Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[ERROR] Failed to serve: %v", err)
	}
}
