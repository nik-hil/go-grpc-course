package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nik-hil/go-grpc-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greent Server function was invoked with %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s count %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
