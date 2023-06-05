package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Keania-Eric/grpc-go-course/calculator/proto"
)

func primeNumberDecomposition(c pb.CalculatorServiceClient) {

	log.Println("prime number decomposition invoked")

	req := &pb.PrimeNumberDecompositionRequest{
		Number: 120,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error occureed while calling prime number decomposition %v \n", err)
	}

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream : %v", err)
		}

		log.Printf("%d is a factor \n", msg.Result)
	}
}
