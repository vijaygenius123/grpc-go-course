package main

import (
	"fmt"
	pb "github.com/vijaygenius123/grpc-go-course/sum/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.SumServiceServer
}

var addr string = "0.0.0.0:8000"

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {

	}

	fmt.Printf("Listening on %v", addr)

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
