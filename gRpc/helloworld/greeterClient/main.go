package main

import (
	"context"
	"log"

	pb "github.com/yudong2015/goLearn/gRpc/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:5000"
	defaultName = "Jack"
)

func main() {
	//set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting Response: %s", r.Message)
}
