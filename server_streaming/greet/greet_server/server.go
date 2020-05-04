package main

import "github.com/maneeshchaturvedi/grpc-go-playground/server_streaming/greet/greetpb"

import "strconv"

import "net"

import "log"

import "google.golang.org/grpc"

import "time"

import "fmt"

type server struct{}

func (*server) Greet(req *greetpb.GreetRequest, stream greetpb.GreetService_GreetServer) error {
	name := req.GetName().GetName()
	for i := 0; i < 10; i++ {
		greeting := "Hello " + name + " number " + strconv.Itoa(i)
		res := &greetpb.GreetResponse{
			Response: greeting,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Started Server Streaming")
	listener, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start grpc server : %v", err)
	}

}
