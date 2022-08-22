package main

import (
	"fmt"
	pb "github.com/vijaygenius123/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
