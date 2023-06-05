package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Keania-Eric/grpc-go-course/calculator/proto"
)

var addr string = ":50052"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//doSum(c)

	primeNumberDecomposition(c)

}
