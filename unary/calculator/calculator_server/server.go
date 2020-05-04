package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/maneeshchaturvedi/grpc-go-playground/unary/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Calculator Sum invoked with : %v", req)

	param1 := req.GetNums().GetParam1()
	param2 := req.GetNums().GetParam2()
	sum := param1 + param2
	res := &calculatorpb.SumResponse{
		Result: sum,
	}
	return res, nil
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
