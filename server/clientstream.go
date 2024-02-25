package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/sangeeth518/learning-grpc/proto"
)

func (s *Server) ServerClientStreaming(stream proto.Example_ServerClientStreamingServer) error {
	total := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa(total),
			})
		}
		if err != nil {
			return err
		}
		total++
		fmt.Println(req)
	}

}
