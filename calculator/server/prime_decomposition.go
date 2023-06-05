package main

import (
	"log"

	pb "github.com/Keania-Eric/grpc-go-course/calculator/proto"
)

func (s *Server) PrimeNumberDecomposition(req *pb.PrimeNumberDecompositionRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {

	log.Println("we are about decomposing this piece of shit")

	var k int32

	N := req.Number
	k = 2

	for {

		if N <= 1 {
			break
		}

		if N%k == 0 {

			stream.Send(&pb.PrimeNumberDecompositionResponse{Result: k})
			N = N / k

		} else {

			k = k + 1
		}
	}

	return nil
}
