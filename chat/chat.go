package chat

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) SayBye(ctx context.Context, message *Message) (*Message, error) {
	fmt.Println(message)
	return &Message{Body: "Bye From the Server!", Name: "server"}, nil
}

func (s *Server) SayName(ctx context.Context, message *Message) (*Message, error) {
	fmt.Println(message)
	return &Message{Name: "server"}, nil
}
