package main

import (
	"context"
	"fmt"
	pb "github.com/vijaygenius123/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("Greet Was Invoked")
	return &pb.GreetResponse{
		Result: "Hello " + request.FirstName,
	}, nil
}
