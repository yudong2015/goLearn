package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	constant "github.com/yudong2015/goLearn/gRpc/helloworld/common"
	pb "github.com/yudong2015/goLearn/gRpc/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultName = "Jack"
)

func main() {

	//setup tsl certs
	cert, err := tls.LoadX509KeyPair(constant.WORKSPACE+"/tsl/certs/client/client.pem", constant.WORKSPACE+"/tsl/certs/client/client.key")
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
		ServerName:   constant.SERVER_HOST,
		RootCAs:      certPool,
	})

	//set up a connection to the server
	conn, err := grpc.Dial(constant.SERVER_HOST+":"+constant.SERVER_PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting Response: %s", response.Message)

	responseAgain, err := client.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting Response: %s", responseAgain.Message)
}
