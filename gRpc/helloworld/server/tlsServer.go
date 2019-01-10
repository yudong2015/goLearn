package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/yudong2015/goLearn/gRpc/helloworld/common"
	pb "github.com/yudong2015/goLearn/gRpc/helloworld/proto"
)

//tlsMainServer
func tlsMain() {
	// create cert
	cert, err := tls.LoadX509KeyPair(common.SERVER_PEM, common.SERVER_KEY)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(common.CA_PEM)
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	//create tls
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	//create server
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterGreeterServer(server, &greeterServer{})

	lis, err := net.Listen("tcp", common.SERVER_PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server.Serve(lis)
}
