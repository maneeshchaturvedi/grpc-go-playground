package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/maneeshchaturvedi/grpc-go-playground/bidirectional_streaming/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindMaximun(stream calculatorpb.FindMaximumService_FindMaximunServer) error {
	fmt.Println("Find Maximum Server invoked by streaming client")
	var currentMax int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving client request : %v", err)
			return err
		}
		num := req.GetNumber()
		if num > currentMax {
			currentMax = num
		}
		err = stream.Send(&calculatorpb.FindMaximumResponse{
			Maximum: currentMax,
		})
		if err != nil {
			log.Fatalf("error sending response to client : %v", err)
			return err
		}
	}

}

func main() {
	fmt.Println("Starting Server for Client streaming")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen on address : %v", err)
	}

	grpcServer := grpc.NewServer()
	calculatorpb.RegisterFindMaximumServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server:%v", err)
	}

}
