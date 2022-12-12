package main

import (
	"context"
	"fmt"
	pb "grpcApi/greet/proto"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	reqs := []pb.GreetRequest{
		pb.GreetRequest{
			FirstName: "abhishek",
		},
		pb.GreetRequest{
			FirstName: "kartik",
		},
		pb.GreetRequest{
			FirstName: "vatsal",
		},
	}

	stream, _ := c.LongGreet(context.Background())

	for _, req := range reqs {
		stream.Send(&req)
		time.Sleep(1 * time.Second)
	}

	res, _ := stream.CloseAndRecv()
	fmt.Println(res.Result)

}
