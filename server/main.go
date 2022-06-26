package main

import (
	"context"
	"net"
	"github.com/adammcclenaghan/go-grpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
)

// So that we can implement the ChatServiceServer interface in service.go
type server struct {}

func main() {
	fmt.Println("Server running")
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterChatServiceServer(srv, &server{})
	reflection.Register(srv)

	// Serve on the server
	if e:= srv.Serve(listener); e!= nil {
		panic(e)
	}
}

func (s *server) ExchangeMessage(ctx context.Context, message *proto.ChatMessage) (*proto.ServerResponse, error) {
	fmt.Println("Server received a message")
	msgContent := message.MessageContent
	msgClientName := message.ClientName
	response := fmt.Sprintf("[%s] %s", msgClientName, msgContent)
	return &proto.ServerResponse {FormattedMessage: response}, nil
}

