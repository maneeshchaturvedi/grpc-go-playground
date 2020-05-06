package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/maneeshchaturvedi/grpc-go-playground/unary/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline Server invoke with : %v\n", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("client cancelled the request")
			return nil, status.Error(codes.Canceled, "Client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	greeting := "Hello " + firstName + " " + lastName

	res := &greetpb.GreetWithDeadlineResponse{
		Response: greeting,
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
