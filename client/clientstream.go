package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/learning-grpc/proto"
)

func ClientConServer(c *gin.Context) {
	req := []*proto.HelloRequest{
		{SomeString: "req 1"},
		{SomeString: "req 2"},
		{SomeString: "req 3"},
		{SomeString: "req 4"},
		{SomeString: "req 5"},
	}
	stream, err := client.ServerClientStreaming(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	for _, re := range req {
		if err := stream.Send(re); err != nil {
			fmt.Printf("error while sending req %v", err)
		}
		log.Printf("sent the reeq %v", re)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("there is some error", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": res,
	})
}
