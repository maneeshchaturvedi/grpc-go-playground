package main

import (
	"fmt"
	"log"
	"net"

	"github.com/maneeshchaturvedi/grpc-go-playground/server_streaming/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) PrimeFactrors(req *calculatorpb.PrimeFactorsRequest, stream calculatorpb.CalculatorService_PrimeFactrorsServer) error {
	fmt.Printf("Calculator PrimeFactors invoked with : %v\n", req)
	number := req.GetNum()
	if number < 0 {
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %v", number),
		)
	}
	k := int32(2)
	num := req.GetNum()
	for num > 1 {
		if num%k == 0 {
			res := &calculatorpb.PrimeFactorsResponse{
				Num: k,
			}
			stream.Send(res)
			num = num / k
		} else {
			k = k + 1
		}
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
