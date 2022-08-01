package main

import (
	"context"
	"log"

	pb "github.com/nik-hil/go-grpc-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greent Server function was invoked with &v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
