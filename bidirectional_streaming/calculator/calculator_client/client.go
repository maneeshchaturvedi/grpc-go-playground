package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/maneeshchaturvedi/grpc-go-playground/bidirectional_streaming/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator streaming client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	client := calculatorpb.NewFindMaximumServiceClient(connection)
	doBiDirectionalStreaming(client)

}
func doBiDirectionalStreaming(client calculatorpb.FindMaximumServiceClient) {
	fmt.Println("Making a client streaming RPC call")

	requests := []*calculatorpb.FindMaximumRequest{
		&calculatorpb.FindMaximumRequest{
			Number: 100,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 13,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 21,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 419,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 3132,
		},
	}

	stream, err := client.FindMaximun(context.Background())
	if err != nil {
		log.Fatalf("Failed to send request : %v", err)
		return
	}
	waitChan := make(chan struct{})
	go func() {
		for _, req := range requests {
			log.Printf("Sending req :%v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving :%v", err)
				break
			}
			log.Printf("recieved result from server : %v\n", response.GetMaximum())
		}
		close(waitChan)
	}()

	<-waitChan
}
