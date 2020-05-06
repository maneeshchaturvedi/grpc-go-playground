package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/maneeshchaturvedi/grpc-go-playground/unary/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello from client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := greetpb.NewGreetServiceClient(connection)

	//doUnary(client)

	doUnaryWithDeadline(client, 5*time.Second)
	doUnaryWithDeadline(client, 1*time.Second)

}

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Println("Making a unary RPC call")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Maneesh",
			LastName:  "Chaturvedi",
		},
	}

	response, err := client.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling greet rpc : %v", err)
	}

	log.Printf("Response from server: %v", response.Result)
}

func doUnaryWithDeadline(client greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Making a unary with deadline RPC call")
	request := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Maneesh",
			LastName:  "Chaturvedi",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	response, err := client.GreetWithDeadline(ctx, request)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
			} else {
				log.Fatalf("something went wrong : %v", statusErr)
			}
		} else {
			log.Fatalf("Error while calling greet rpc : %v", err)
		}
		return
	}

	log.Printf("Response from server: %v", response.GetResponse())
}
