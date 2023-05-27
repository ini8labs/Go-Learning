package main

import (
	"client-server/chat/chat"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error connecting the server")
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client",
	}
	response, err := c.SayHello(context.TODO(), &message)
	if err != nil {
		log.Fatalf("error connecting with server %s", err)
	}

	log.Println("Message from server %s", response.Body)
}
