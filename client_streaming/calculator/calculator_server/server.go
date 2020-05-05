package main

import "github.com/maneeshchaturvedi/grpc-go-playground/client_streaming/calculator/calculatorpb"

import "io"

import "log"

import "net"

import "google.golang.org/grpc"

import "fmt"

type server struct{}

func (*server) Average(stream calculatorpb.AverageService_AverageServer) error {
	fmt.Println("Average Server invoked by streaming client")
	var average float32 = 0
	var sum int64 = 0
	var count int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average = (float32)(sum / count)
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while calculating average : %v", err)
		}
		fmt.Printf("Got request : %v\n", req.GetNum())
		sum += req.GetNum()
		count++
	}

}

func main() {
	fmt.Println("Starting Server for Client streaming")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen on address : %v", err)
	}

	grpcServer := grpc.NewServer()
	calculatorpb.RegisterAverageServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server:%v", err)
	}

}
