package main

import (
	"context"
	"log"

	pb "github.com/bloodgroup-cplusplus/grpc_go_course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Chad",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)

}
