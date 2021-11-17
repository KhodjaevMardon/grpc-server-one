package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"grpc-server-one/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body:     "dummy text",
		Priority: "dummy status",
	}

	for message.Body != "stop" {

		fmt.Print("Enter the text of your message :")
		message.Body, _ = reader.ReadString('\n')
		fmt.Print("Enter the priority of your message:")
		fmt.Scan(&message.Priority)
		_, err := c.PostMessageToSend(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SendMessage: %s", err)
		}

	}
}
