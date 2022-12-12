package main

import (
	"context"
	"fmt"
	pb "grpcApi/greet/proto"
	"io"
	"log"
)

func doGreetManyTime(c pb.GreetServiceClient) {
	log.Println("DoGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Abhihek",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Result)

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Abhishek",
	})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return
			} else {
				log.Panicln(err)
			}
		}
		fmt.Println(msg.Result)
	}
}
