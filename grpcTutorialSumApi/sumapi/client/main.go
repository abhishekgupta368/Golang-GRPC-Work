package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "sumApi/sumapi/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var addr string = "localhost:50051"

func getSum(c pb.SumServiceClient) {
	//sen the cancellation signal to the system
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.GetSum(ctx, &pb.SumRequest{
		Num1: 10,
		Num2: 100,
	})

	// handling gprc baed error
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println(e.Message())
			log.Println(e.Code())
		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(res.Sum)
}

func getPrimeNumber(c pb.PrimeServiceClient) {
	stream, err := c.GetPrime(context.Background(), &pb.PrimeRequest{
		Num: 120,
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
		fmt.Println(msg.Num)
	}
}

func getAvgNumber(c pb.AvgServiceClient) {
	num := []int{1, 2, 3, 4, 5}
	stream, _ := c.GetAvg(context.Background())
	for _, req := range num {
		stream.Send(&pb.AvgRequest{
			Num: int32(req),
		})
		time.Sleep(1 * time.Second)
	}
	res, _ := stream.CloseAndRecv()
	fmt.Println(res.Num)
}

func getMaxNumber(c pb.MaxServiceClient) {
	stream, err := c.GetMax(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	nums := []*pb.MaxRequest{
		&pb.MaxRequest{
			Num: 1,
		},
		&pb.MaxRequest{
			Num: 3,
		},
		&pb.MaxRequest{
			Num: 2,
		},
		&pb.MaxRequest{
			Num: 5,
		},
		&pb.MaxRequest{
			Num: 4,
		},
	}
	waitc := make(chan *struct{})
	go func() {
		for _, num := range nums {
			stream.Send(num)
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
			fmt.Println(res.Num)
		}
		close(waitc)
	}()
	<-waitc
}

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()
	// fmt.Printf("listening on Addr: %s\n", addr)
	fmt.Println("-------unary--------------------------")
	c := pb.NewSumServiceClient(conn)
	getSum(c)

	fmt.Println("-------server streaming---------------")
	c1 := pb.NewPrimeServiceClient(conn)
	getPrimeNumber(c1)

	fmt.Println("-------clinet streaming---------------")
	c2 := pb.NewAvgServiceClient(conn)
	getAvgNumber(c2)

	fmt.Println("-------bidirectinal streaming---------")
	c3 := pb.NewMaxServiceClient(conn)
	getMaxNumber(c3)
}
