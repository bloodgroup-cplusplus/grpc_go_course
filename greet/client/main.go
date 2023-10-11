package main

import (
	"log"

	pb "github.com/bloodgroup-cplusplus/grpc_go_course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// create an instance of greet

var addr string = "localhost:50051"

func main() {

	// create a fucntion called grpc dial

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to :%v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

}
