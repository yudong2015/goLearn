package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	constant "github.com/yudong2015/goLearn/gRpc/helloworld/common"
	pb "github.com/yudong2015/goLearn/gRpc/helloworld/proto"
)

type serverStruct struct{}

func (s *serverStruct) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *serverStruct) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

//server
func main() {
	cert, err := tls.LoadX509KeyPair(constant.WORKSPACE+"/tsl/certs/server/server.pem", constant.WORKSPACE+"/tsl/certs/server/server.key")
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(constant.WORKSPACE + "/tsl/certs/ca.pem")
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterGreeterServer(server, &serverStruct{})

	lis, err := net.Listen("tcp", ":"+constant.SERVER_PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server.Serve(lis)
}
