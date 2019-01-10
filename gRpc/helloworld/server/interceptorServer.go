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

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

//interceptorMainServer
func interceptorMain() {
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

	// create new tls
	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	//grpc options
	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			common.RecoveryInterceptor,
			common.LoggingInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &greeterServer{})

	lis, err := net.Listen("tcp", common.SERVER_PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server.Serve(lis)
}
