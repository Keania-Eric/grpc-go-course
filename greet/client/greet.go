package main

import (
	"context"
	"log"

	pb "github.com/Keania-Eric/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {

	log.Println("Starting to do a Unary RPC...")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Eric",
	})

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
