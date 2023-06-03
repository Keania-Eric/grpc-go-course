package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Keania-Eric/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	log.Println("do greet many times was invoked")

	req := &pb.GreetRequest{
		FirstName: "Keania",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greetmanytimes %v \n", err)
	}

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the steam : %v", err)
		}

		log.Printf("Greet many times %s \n", msg.Result)
	}
}
