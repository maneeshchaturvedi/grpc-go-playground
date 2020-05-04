package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/maneeshchaturvedi/grpc-go-playground/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Server invoke with : %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	greeting := "Hello " + firstName + " " + lastName

	res := &greetpb.GreetResponse{
		Result: greeting,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")
	//create listener on port
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//create new grpc server
	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &server{})
	//serve requests on listener
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
