package main

import (
	"context"
	"fmt"
	pb "grpcApi/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	reqs := []*pb.GreetRequest{
		&pb.GreetRequest{
			FirstName: "Abhishek",
		},
		&pb.GreetRequest{
			FirstName: "Vatsal",
		},
		&pb.GreetRequest{
			FirstName: "kartik",
		},
	}
	waitc := make(chan *struct{})

	go func() {
		for _, req := range reqs {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err.Error())
				break
			}
			fmt.Println(res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
