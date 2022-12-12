package main

import (
	pb "grpcApi/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	for {
		req,err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil{
			log.Fatalln(err)
		}

		res := "hello: "+ req.FirstName
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil{
			log.Fatal(err)
		}
	}
	return nil
}
