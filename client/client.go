package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gbarnovi/grpc-example/chat"
	"google.golang.org/grpc"
)

func main() {

	port := 9000

	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Got an error while tring to dial :%v - %v", port, err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	ctx := context.Background()

	msg, err := c.SayHello(ctx, &chat.Message{Body: "Message from client"})
	if err != nil {
		log.Fatalf("Got an error while trying to send a message to grpc server - %v", err)
	}

	fmt.Printf("msg.Body: %v\n", msg.Body)

}
