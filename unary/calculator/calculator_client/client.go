package main

import (
	"context"
	"fmt"
	"log"

	"github.com/maneeshchaturvedi/grpc-go-playground/unary/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	client := calculatorpb.NewCalculatorServiceClient(connection)
	doUnary(client)

}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Making a unary RPC call")
	request := &calculatorpb.SumRequest{
		Nums: &calculatorpb.Nums{
			Param1: 122,
			Param2: 42,
		},
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling greet rpc : %v", err)
	}

	log.Printf("Response from server: %v", response.Result)
}
