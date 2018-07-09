package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/nyogjtrc/grpc-example/proto"
	"google.golang.org/grpc"
)

type echoServer struct{}

// Echo
func (s *echoServer) Echo(ctx context.Context, in *pb.EchoMessage) (*pb.EchoMessage, error) {
	reply := new(pb.EchoMessage)
	reply.Value = "echo:" + in.Value
	return reply, nil
}

func main() {
	go grpcServer()
	gatewayServer()
}

func grpcServer() {
	fmt.Println("grpc echo server :8888")

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("can not listen %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterEchoServiceServer(s, &echoServer{})
	s.Serve(lis)
}

func gatewayServer() {
	fmt.Println("RESTful echo server :9999")
	grpcAddr := "localhost:8888"
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("can not register endpoint %v", err)
	}

	http.ListenAndServe(":9999", mux)
}
