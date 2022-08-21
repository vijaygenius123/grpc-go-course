package main

import (
	pb "github.com/vijaygenius123/grpc-go-course/sum/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:8000"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to remote")
	}

	defer conn.Close()

	c := pb.NewSumServiceClient(conn)

	doSum(c)

}
