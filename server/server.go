package main

import (
	calculatorpb "calculator/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	log.Printf("Calculate func was invoked with %v", req)
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

func main() {
	log.Println("Hello, I'm the calculation server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
