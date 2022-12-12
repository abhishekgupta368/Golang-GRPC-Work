package main

import (
	"fmt"
	"log"

	pb "grpcApi/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewServerTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalln(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	// in case of the no ssl certifcates
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()
	// fmt.Printf("listening on Addr: %s\n", addr)
	c := pb.NewGreetServiceClient(conn)
	fmt.Println("----------- unary -------------------")
	doGreet(c)
	fmt.Println("---------server to client------------")
	doGreetManyTime(c)
	fmt.Println("---------client to server------------")
	doLongGreet(c)
	fmt.Println("---------bidirectional---------------")
	doGreetEveryone(c)
	fmt.Println("-------------------------------------")

}
