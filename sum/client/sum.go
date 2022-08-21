package main

import (
	"context"
	"fmt"
	pb "github.com/vijaygenius123/grpc-go-course/sum/proto"
	"log"
)

func doSum(c pb.SumServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalln("Error Returning Response")
	}

	fmt.Println("Sum Is ", res.Result)
}
