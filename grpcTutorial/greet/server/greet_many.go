package main

import (
	pb "grpcApi/greet/proto"
	"strconv"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := strconv.Itoa(i) + "->" + in.FirstName
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
