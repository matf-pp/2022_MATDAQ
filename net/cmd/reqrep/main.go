package main

import (
	"fmt"
	"github.com/matf-pp/2022_MATDAQ/net/pkg/reqrep"
	"os"
)

type MessageType byte

const (
	MsgEven MessageType = 0
	MsgOdd  MessageType = 1
)

type Message struct {
	msgType MessageType
	data    string
}

func (msg *Message) ToBytes() []byte {
	return append([]byte(msg.data), byte(msg.msgType))
}

func (msg *Message) FromBytes(data []byte) {
	n := len(data)
	msg.msgType = MessageType(data[n-1])
	msg.data = string(data[:n-1])
}

func client(addr string) {
	msgs := [4]Message{
		{MsgEven, "Message 1"},
		{MsgOdd, "Message 2"},
		{MsgEven, "Message 3"},
		{MsgOdd, "Message 4"},
	}

	client := reqrep.NewClient(addr)
	for _, msg := range msgs {
		resp, err := reqrep.SendRequest(client, msg.ToBytes())
		if err != nil {
			panic(err)
		}

		var msg Message
		msg.FromBytes(resp)

		fmt.Println(msg)
	}

	client.Close()
}

func server(addr string) {
	res := make(chan []byte, 1024)
	server := reqrep.NewServer(addr, res)
	go reqrep.Listen(server)

	for data := range res {
		var msg Message
		msg.FromBytes(data)

		fmt.Printf("Got message %s !\n", msg)

		var err error
		switch msg.msgType {
		case MsgEven:
			res := Message{msgType: MsgEven, data: "Okk"}
			err = server.Respond(res.ToBytes())
		case MsgOdd:
			res := Message{msgType: MsgOdd, data: "Okk"}
			err = server.Respond(res.ToBytes())
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to send response")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: reqrep client|server <URL>?\n")
		os.Exit(1)
	}
	fmt.Println("Request response example")

	addr := "tcp://127.0.0.1:6000"
	if len(os.Args) == 3 {
		addr = os.Args[2]
	}

	if os.Args[1] == "client" {
		client(addr)
	} else {
		server(addr)
	}
}
