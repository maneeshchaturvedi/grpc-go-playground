package main

import (
	"context"
	"fmt"
	"log"

	"github.com/maneeshchaturvedi/grpc-go-playground/client_streaming/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator streaming client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	client := calculatorpb.NewAverageServiceClient(connection)
	doClientStreaming(client)

}
func doClientStreaming(client calculatorpb.AverageServiceClient) {
	fmt.Println("Making a client streaming RPC call")

	requests := []*calculatorpb.AverageRequest{
		&calculatorpb.AverageRequest{
			Num: 100,
		},
		&calculatorpb.AverageRequest{
			Num: 13,
		},
		&calculatorpb.AverageRequest{
			Num: 21,
		},
		&calculatorpb.AverageRequest{
			Num: 419,
		},
		&calculatorpb.AverageRequest{
			Num: 3132,
		},
	}

	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalf("Failed to send request : %v", err)
	}
	for _, req := range requests {
		stream.Send(req)
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to get response :%v", err)
	}

	log.Printf("recieved result from server : %v", result.GetAverage())

}
