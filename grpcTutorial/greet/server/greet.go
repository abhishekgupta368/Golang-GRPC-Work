package main

import (
	"context"
	pb "grpcApi/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
