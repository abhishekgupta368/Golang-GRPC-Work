package main

import (
	"context"
	"fmt"
	pb "grpcApi/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("DoGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Abhihek",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Result)
}
