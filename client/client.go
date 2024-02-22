package main

import (
	"context"

	"github.com/sangeeth518/learning-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewExampleClient(conn)
	req := &proto.HelloRequest{SomeString: "Hello from client"}
	client.ServerReply(context.TODO(), req)

}
