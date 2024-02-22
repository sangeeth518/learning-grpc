package server

import (
	"context"
	"fmt"
	"net"

	proto "github.com/sangeeth518/learning-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {

	listner, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer() //Engine
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)
	if err := srv.Serve(listner); err != nil {
		panic(err)
	}

}
func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("Recieve request from client", req.SomeString)
	fmt.Println("Hello from server")
	return &proto.HelloResponse{Reply: "Hello"}, nil

}
