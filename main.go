package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/gbarnovi/grpc-example/chat"
)

func main() {

	port := "9000"

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		log.Fatalf("Got an error while trying to listen on port :%v - %v", port, err)
	}

	grpcServer := grpc.NewServer()

	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, s)

	fmt.Printf("Staring grpc on port %v", port)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Got an error while trying to serve grpc %v", err)
	}

}
