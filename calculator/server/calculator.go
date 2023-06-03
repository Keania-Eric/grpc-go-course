package main

import (
	"context"
	"log"

	pb "github.com/Keania-Eric/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	firstNumber := in.GetFirstNumber()
	secondNumber := in.GetSecondNumber()
	sum := firstNumber + secondNumber
	return &pb.SumResponse{Result: sum}, nil
}
