package main

import (
	pb "github.com/vijaygenius123/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:8000"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
