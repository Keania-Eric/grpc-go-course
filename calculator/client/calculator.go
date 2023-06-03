package main

import (
	"context"
	"log"

	pb "github.com/Keania-Eric/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {

	log.Println("Starting to do a Sum Unary RPC...")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Failed to Sum: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
}
