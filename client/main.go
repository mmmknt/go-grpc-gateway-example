package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/mmmknt/go-grpc-gateway-example/service"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	ctx, cancel:= context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Echo(ctx, &pb.StringMessage{Value: "World"})
	if err != nil {
		log.Fatalf("failed to echo: %v", err)
		return
	}
	log.Printf("Echo: %v", res)
}
