package main

import "fmt"

import "google.golang.org/grpc"

import "log"

import "github.com/maneeshchaturvedi/grpc-go-playground/server_streaming/greet/greetpb"

import "context"

import "io"

func main() {
	fmt.Println("Hello from client")

	connection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := greetpb.NewGreetServiceClient(connection)

	doServerStreaming(client)

}

func doServerStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Making a call to a streaming server")

	req := &greetpb.GreetRequest{
		Name: &greetpb.Name{
			Name: "Maneesh",
		},
	}
	resStream, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to invoke Server Streaming RPC : %v ", err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			//end of stream, so break
			break
		}
		if err != nil {
			log.Fatalf("failed to read response: %v", err)
		}
		fmt.Printf("Response Received : %v\n ", res.GetResponse())
	}

}
