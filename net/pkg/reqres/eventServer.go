package reqres

import (
	"fmt"
	"os"
)

type messageTransform func(Message) Message

type MessageKey interface {
	~byte
}

type EventServer[K MessageKey] struct {
	addr     string
	handlers map[K]messageTransform
}

func NewEventServer[K MessageKey](addr string) *EventServer[K] {
	return &EventServer[K]{
		addr,
		make(map[K]messageTransform),
	}
}

func (sb *EventServer[K]) On(messageKey K, f messageTransform) {
	sb.handlers[messageKey] = f
}

func (sb *EventServer[K]) Start() {
	res := make(chan []byte, 1024)
	server := NewServer(sb.addr, res)
	go Listen(server)

	for data := range res {
		var msg Message
		msg.FromBytes(data)

		fmt.Printf("Got message %s !\n", msg)

		res := sb.handlers[K(msg.MsgType)](msg)
		err := server.Respond(res.ToBytes())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to send response")
		}
	}
}
