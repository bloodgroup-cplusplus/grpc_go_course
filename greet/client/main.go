package main

import (
	"log"

	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {

	// create a fucntion called grpc dial

	conn, err := grpc.Dial(addr)

	if err != nil {
		log.Fatalf("Failed to connect to :%v\n", err)
	}
	defer conn.Close()
}
