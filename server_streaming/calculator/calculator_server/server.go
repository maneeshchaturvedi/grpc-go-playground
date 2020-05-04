package main

import (
	"fmt"
	"log"
	"net"

	"github.com/maneeshchaturvedi/grpc-go-playground/server_streaming/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeFactrors(req *calculatorpb.PrimeFactorsRequest, stream calculatorpb.CalculatorService_PrimeFactrorsServer) error {
	fmt.Printf("Calculator PrimeFactors invoked with : %v", req)
	results := make([]int32, 0)
	var k int32
	k = 2
	num := req.GetNum()
	for num > 1 {
		if num%k == 0 {
			results = append(results, k)
			num = num / k
		} else {
			k = k + 1
		}
	}

	for _, num := range results {
		res := &calculatorpb.PrimeFactorsResponse{
			Num: num,
		}
		stream.Send(res)
	}

	return nil
}

func main() {
	fmt.Println("Started Calculator service")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//create new grpc server
	grpcServer := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error starting server : %v", err)
	}

}
