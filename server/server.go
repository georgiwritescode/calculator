package main

import (
	calculatepb "calculator/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	log.Printf("Calculate func was invoked with %v", req)
	a := req.Numbers.GetFirst()
	b := req.Numbers.GetSecond()

	sum := a + b
	res := &calculatepb.CalculateResponse{
		Sum: sum,
	}

	return res, nil
}

func main() {
	log.Println("Hello, I'm the calculation server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
