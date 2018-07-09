package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
	fmt.Println("grpc example echo server :8888")

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("can not listen %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterEchoServiceServer(s, &echoServer{})
	s.Serve(lis)
}
