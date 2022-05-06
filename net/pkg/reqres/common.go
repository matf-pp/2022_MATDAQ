package reqres

import (
	"fmt"
	"os"
)

type MessageType byte

const (
	MsgEven MessageType = 0
	MsgOdd  MessageType = 1
)

type Message struct {
	MsgType MessageType
	Data    string
}

func (msg *Message) ToBytes() []byte {
	return append([]byte(msg.Data), byte(msg.MsgType))
}

func (msg *Message) FromBytes(data []byte) {
	n := len(data)
	msg.MsgType = MessageType(data[n-1])
	msg.Data = string(data[:n-1])
}

type Serializable interface {
	ToBytes() []byte
	FromBytes(data []byte)
}

func die(format string, v ...any) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}
