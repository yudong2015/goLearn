package main

import (
	"context"
	"log"

	"github.com/yudong2015/goLearn/gRpc/helloworld/common"
	pb "github.com/yudong2015/goLearn/gRpc/helloworld/proto"
	"google.golang.org/grpc"
)

func simpleMain() {

	//set up a connection to the server
	conn, err := grpc.Dial(common.SERVER_HOST+common.SERVER_PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: common.DEFAULT_NAME})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting Response: %s", response.Message)

	responseAgain, err := client.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: common.DEFAULT_NAME})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting Response: %s", responseAgain.Message)
}
