package main

import pb "BlogGRPC/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
