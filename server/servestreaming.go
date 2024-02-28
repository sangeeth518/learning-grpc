package main

import (
	"fmt"

	"github.com/sangeeth518/learning-grpc/proto"
)

func (s *Server) SereverStreaming(req *proto.HelloRequest, stream proto.Example_SereverStreamingServer) error {
	fmt.Println(req.SomeString)
	Message := []*proto.HelloResponse{
		{Reply: "hello"},
		{Reply: "where are you"},
		{Reply: "are you ok"},
		{Reply: "glad to hear"},
		{Reply: "okk bye"},
	}
	for _, msg := range Message {
		if err := stream.Send(msg); err != nil {
			return err
		}

	}
	return nil

}
