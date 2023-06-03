package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Keania-Eric/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {

	// we create a connection to the server with grpc dial
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// we check for null safety
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	// now we know connection was successful we defer closing it to the end of function
	defer conn.Close()

	// we use the client to create a connection
	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)

	doGreetManyTimes(c)
}
