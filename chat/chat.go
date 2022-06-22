package chat

import (
	"context"
	"fmt"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s Server) SayHello(ctx context.Context, m *Message) (*Message, error) {
	fmt.Printf("We received a message: %v", m.GetBody())

	return &Message{Body: fmt.Sprintf("I am sending you an response on message: %v", m.GetBody())}, nil
}
