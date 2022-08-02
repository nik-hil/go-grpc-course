package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nik-hil/go-grpc-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Nikhil",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)

}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Nikhil",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling doGreetManyTimes")
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("doGreetManyTimes %s\n", msg.Result)
	}

}
