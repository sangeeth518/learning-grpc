package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	proto "github.com/sangeeth518/learning-grpc/proto"
)

func ClientServerStream(c *gin.Context) {
	stream, err := client.SereverStreaming(context.Background(), &proto.HelloRequest{SomeString: "hello"})
	if err != nil {
		fmt.Println("something error")
		return
	}
	count := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error in streaming %v", err)
		}
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
		count++
	}
	c.JSON(200, gin.H{
		"message": count,
	})

}
