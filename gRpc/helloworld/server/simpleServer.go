package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/yudong2015/goLearn/gRpc/helloworld/common"
	pb "github.com/yudong2015/goLearn/gRpc/helloworld/proto"
)

//simpleMainServer
func simpleMain() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeterServer{})

	lis, err := net.Listen("tcp", common.SERVER_PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server.Serve(lis)
}
