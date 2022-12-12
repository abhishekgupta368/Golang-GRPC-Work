package main

import (
	"fmt"
	pb "grpcApi/greet/proto"
	"io"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	res := ""
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.GreetResponse{
					Result: res,
				})
			} else {
				return err
			}
		}
		res += req.FirstName + " "
	}
	fmt.Println(res)
	return nil
}
