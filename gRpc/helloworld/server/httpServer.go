package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/yudong2015/goLearn/gRpc/helloworld/common"
	pb "github.com/yudong2015/goLearn/gRpc/helloworld/proto"
)

func getHTTPServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("jack: go-grpc http server.."))
	})
	return mux
}

//httpMainServer
func httpMain() {
	// setup cert
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

	//create grpc server
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterGreeterServer(server, &greeterServer{})

	//create http server
	mux := getHTTPServerMux()
	http.ListenAndServeTLS(
		common.SERVER_PORT,
		common.SERVER_PEM,
		common.SERVER_KEY,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
			return
		}),
	)

}
