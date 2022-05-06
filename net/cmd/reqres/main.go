package main

import (
	"fmt"
	"github.com/matf-pp/2022_MATDAQ/net/pkg/reqres"
	"os"
)

func client(addr string) {
	msgs := [4]reqres.Message{
		{reqres.MsgEven, "Message 1"},
		{reqres.MsgOdd, "Message 2"},
		{reqres.MsgEven, "Message 3"},
		{reqres.MsgOdd, "Message 4"},
	}

	client := reqres.NewClient(addr)
	for _, msg := range msgs {
		resp, err := reqres.SendRequest(client, msg.ToBytes())
		if err != nil {
			panic(err)
		}

		var msg reqres.Message
		msg.FromBytes(resp)

		fmt.Println(msg)
	}
	client.Close()
}

func server(addr string) {
	eventServer := reqres.NewEventServer[reqres.MessageType](addr)
	eventServer.On(reqres.MsgEven, func(msg reqres.Message) reqres.Message {
		return reqres.Message{MsgType: reqres.MsgOdd, Data: "uspeo even"}
	})
	eventServer.On(reqres.MsgOdd, func(msg reqres.Message) reqres.Message {
		return reqres.Message{MsgType: reqres.MsgEven, Data: "uspeo odd"}
	})
	eventServer.Start()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: reqres client|server <URL>?\n")
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
