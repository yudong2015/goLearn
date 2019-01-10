package common

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const (
	SERVER_PORT  = ":5000"
	SERVER_HOST  = "jack.localhost.com"
	DEFAULT_NAME = "jack"
	WORKSPACE    = "/Users/yudong/workspace/go/src/github.com/yudong2015/goLearn/gRpc/helloworld"
	CA_PEM       = WORKSPACE + "/tsl/certs/ca.pem"
	CLIENT_KEY   = WORKSPACE + "/tsl/certs/client/client.key"
	CLIENT_PEM   = WORKSPACE + "/tsl/certs/client/client.pem"
	SERVER_PEM   = WORKSPACE + "/tsl/certs/server/server.pem"
	SERVER_KEY   = WORKSPACE + "/tsl/certs/server/server.key"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	return handler(ctx, req)
}
