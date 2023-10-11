package main

import (
	"log"
	"net"

	pb "github.com/bloodgroup-cplusplus/grpc_go_course/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {

	// create a listener in this addrss
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Faild to listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	// create a grpc server object

	// create a server instance
	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v\n", err)
	}
}
