package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/maneeshchaturvedi/grpc-go-playground/server_streaming/calculator/calculatorpb"
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
	doServerStreaming(client)

}

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Making a server streaming RPC call")
	request := &calculatorpb.PrimeFactorsRequest{
		Num: 13342256,
	}

	resStream, err := client.PrimeFactrors(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling greet rpc : %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			//reached end of stream
			break
		}
		if err != nil {
			log.Fatalf("error receiving response : %v", err)
		}
		log.Printf("received response : %v", res.GetNum())

	}

}
