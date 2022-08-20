package main

import (
	"context"
	"fmt"
	pb "github.com/vijaygenius123/grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Vijay",
	})

	if err != nil {
		log.Fatalln("Could not run greet")
	}
	fmt.Println("Greet returned " + res.Result)
}
