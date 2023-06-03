package main

import (
	"log"
	"net"

	pb "github.com/Keania-Eric/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = ":50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	log.Printf("Listening on %s", addr)

	// create a new gRPC server
	s := grpc.NewServer()

	// register the server with the gRPC server
	pb.RegisterCalculatorServiceServer(s, &Server{})

	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
