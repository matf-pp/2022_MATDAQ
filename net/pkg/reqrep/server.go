package reqrep

import (
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/rep"
)

type ConvertableFromBytes[C any] interface {
	FromBytes(data []byte) C
}

type Server[C ConvertableFromBytes[C]] struct {
	sock mangos.Socket
	addr string
	out  chan C
}

func newServer[C ConvertableFromBytes[C]](addr string, out chan C) *Server[C] {
	sock, err := rep.NewSocket()
	if err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	return &Server[C]{sock, addr, out}
}

func (sr *Server[C]) Listen() error {
	for {
		msg, err := sr.sock.Recv()
		// TODO: probably shouldn't die if can't read a socket
		if err != nil {
			die("cannot receive on rep socket", err.Error())
		}
		sr.out <- C.FromBytes(msg)
		// TODO: define a protocol
		err = sr.sock.Send([]byte("Ok"))
		if err != nil {
			die("cannot send a reply", err.Error())
		}
	}
}
