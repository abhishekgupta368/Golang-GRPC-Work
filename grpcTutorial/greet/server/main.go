package main

import (
	"fmt"
	pb "grpcApi/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Printf("istening on Addr: %s\n", addr)
	// command to build and add ssl certifiv=cates to the system
	// cd ssl
	// chmod +x ssl.sh
	// ./ssl.sh
	tls := true
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalln(err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	// incase of no certificates
	// server := grpc.NewServer()
	server := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(server, &Server{})
	reflection.Register(server)
	if err = server.Serve(lis); err != nil {
		log.Println(err.Error())
	}
}
