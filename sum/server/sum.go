package main

import (
	"context"
	pb "github.com/vijaygenius123/grpc-go-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: request.Num1 + request.Num2,
	}, nil
}
