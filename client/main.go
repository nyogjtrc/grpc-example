package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/nyogjtrc/grpc-example/proto"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":8888", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		r, err := client.Echo(context.Background(), &pb.EchoMessage{Value: text})
		if err != nil {
			log.Fatalf("%v.Echo error: %v", client, err)
		}
		fmt.Println(r.Value)
	}
}
