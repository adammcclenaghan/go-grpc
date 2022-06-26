package main

import (
	"time"
	"fmt"
	"google.golang.org/grpc"
	"github.com/adammcclenaghan/go-grpc/proto"
	"context"
	"bufio"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter your name and press enter")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("Enter a message and press enter")
	message, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	message = strings.TrimSuffix(message, "\n")

	
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := proto.NewChatServiceClient(conn)
	request := &proto.ChatMessage {
		MessageContent: message,
		ClientName: name,
	}

	// Send message to the server and print the response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if response, err := client.ExchangeMessage(ctx, request); err == nil {
		fmt.Println(response.FormattedMessage)
	} else {
		panic(err)
	}
	
}
