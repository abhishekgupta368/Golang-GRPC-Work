package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	pb "sumApi/sumapi/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

type ServerPrime struct {
	pb.PrimeServiceServer
}

type ServerAvg struct {
	pb.AvgServiceServer
}
type ServerMax struct {
	pb.MaxServiceServer
}

// uninary base request
func (s *Server) GetSum(ctx context.Context, sumRequest *pb.SumRequest) (*pb.SumResponse, error) {
	// handling deadline error
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.SumResponse{
		Sum: sumRequest.Num1 + sumRequest.Num2,
	}, nil
}

// server side rendoring
func (s *ServerPrime) GetPrime(in *pb.PrimeRequest, stream pb.PrimeService_GetPrimeServer) error {
	val := in.Num
	for i := 2; i <= int(math.Sqrt(float64(val))); i++ {
		for val%int32(i) == 0 && val > 0 {
			stream.Send(&pb.PrimeResponse{
				Num: int32(i),
			})
			time.Sleep(2 * time.Second)
			val /= int32(i)
		}
	}
	return nil
}
func (s *ServerAvg) GetAvg(stream pb.AvgService_GetAvgServer) error {
	sum := 0
	cnt := 0
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				avg := (sum / cnt)
				return stream.SendAndClose(&pb.AvgResponse{
					Num: int32(avg),
				})
			} else {
				return err
			}
		}
		sum += int(req.Num)
		cnt++
	}
	return nil
}

func (s *ServerMax) GetMax(stream pb.MaxService_GetMaxServer) error {
	var max int32 = 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		if max < (res.Num) {
			max = (res.Num)
		}
		err = stream.Send(&pb.MaxResponse{
			Num: max,
		})
	}
	return nil
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

	server := grpc.NewServer()
	pb.RegisterSumServiceServer(server, &Server{})
	pb.RegisterPrimeServiceServer(server, &ServerPrime{})
	pb.RegisterAvgServiceServer(server, &ServerAvg{})
	pb.RegisterMaxServiceServer(server, &ServerMax{})
	if err = server.Serve(lis); err != nil {
		log.Println(err.Error())
	}
}
