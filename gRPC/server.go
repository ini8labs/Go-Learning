package main

import (
	"client-server/chat/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090 %s", err)
	}

	s := chat.Server{}
	grpcSever := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcSever, &s)

	log.Println("Running Chat server on port :9090")
	if err := grpcSever.Serve(list); err != nil {
		log.Fatal("Error Running gRPC Server %s", err)
	}
}
