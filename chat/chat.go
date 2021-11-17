package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
)

var (
	botToken = "2108403709:AAHUN0sg0POV1Vpsc3x7ZFCG1lSsbheB0Ow"
	//channelname        = "@test_channel_for_rest_server"
	groupId      int64 = -597182447 // idk, some private group
	floodGroupId int64 = -1001339027507
	channelId    int64 = -1001580495914 // bot's channel
	reader             = bufio.NewReader(os.Stdin)
)

type TgMessage struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type Server struct {
	highMsgs []Message
	medMsgs  []Message
	lowMsgs  []Message
}

func SendMessage(message Message) (Response, error) {
	msg := TgMessage{
		ChatId: channelId,
		Text:   message.Body,
	}

	reqBytes, _ := json.Marshal(msg)
	url := "https://api.telegram.org/bot" + botToken + "/sendMessage"

	res, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

	fmt.Println("post message result: " + res.Status)

	return Response{Status: res.Status}, nil
}

func (s *Server) InitiatePerpetualMessageSender(ctx context.Context, empty *Empty) (*Empty, error) {

	duration := 5000 * time.Millisecond
	for count := 0; ; time.Sleep(duration) {
		msgFound := true
		var response Response
		var msg Message
		if len(s.highMsgs) > 0 {
			response, _ = SendMessage(s.highMsgs[0])
			msg = s.highMsgs[0]
			s.highMsgs = s.highMsgs[1:]
		} else if len(s.medMsgs) > 0 {
			response, _ = SendMessage(s.medMsgs[0])
			msg = s.medMsgs[0]
			s.medMsgs = s.medMsgs[1:]
		} else if len(s.lowMsgs) > 0 {
			response, _ = SendMessage(s.lowMsgs[0])
			msg = s.lowMsgs[0]
			s.lowMsgs = s.lowMsgs[1:]
		} else {
			msgFound = false
		}

		if msgFound {
			fmt.Println("Message found and sent. Body:" + msg.Body + "Result status: " + response.Status)
			count++
		} else {
			fmt.Println("Message not found.")
		}
		fmt.Println("Time:", time.Now(), "Count: ", count)
	}
}

func (s *Server) PostMessageToSend(ctx context.Context, message *Message) (*Empty, error) {
	log.Printf("Received message body from client: %s", message.Body)

	// just add message to corresponding slice
	switch message.Priority {
	case "high":
		s.highMsgs = append(s.highMsgs, *message)
	case "med":
		s.medMsgs = append(s.medMsgs, *message)
	case "low":
		s.lowMsgs = append(s.lowMsgs, *message)
	default:
		fmt.Println("unexpected priority: ", message.Priority)
	}

	return &Empty{}, nil
}
