package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/mmmknt/go-grpc-gateway-example/service"
)

type server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	input := in.GetValue()
	log.Printf("Received: %v", input)

	if input == "error" {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &pb.StringMessage{
		Value:                fmt.Sprintf("Hello, %v", input),
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}