package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(c context.Context, message *Message) (*Message, error) {
	log.Printf("Recieved message %s", message.Body)
	return &Message{Body: "Hello from Server"}, nil
}
