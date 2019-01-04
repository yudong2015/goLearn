package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yudong2015/goLearn/gRpc/helloworld/helloworld"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

//server

const (
	port = ": 5000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
