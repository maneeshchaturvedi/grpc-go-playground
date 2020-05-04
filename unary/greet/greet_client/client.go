package main

import (
	"context"
	"fmt"
	"log"

	"github.com/maneeshchaturvedi/grpc-go-playground/unary/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := greetpb.NewGreetServiceClient(connection)

	unaryService(client)

}

func unaryService(client greetpb.GreetServiceClient) {
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
