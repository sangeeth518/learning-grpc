package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	proto "github.com/sangeeth518/learning-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)
	r := gin.Default()
	r.GET("/sentmsg-to-server/:message", clientConnectionServer)
	r.GET("/sent", ClientConServer)
	r.GET("/server_stream", ClientServerStream)
	r.Run(":8000")

}
func clientConnectionServer(c *gin.Context) {
	message := c.Param("message")
	req := &proto.HelloRequest{SomeString: message}
	client.ServerReply(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent succesfully to the server" + message,
	})
}
