package reqres

import (
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/rep"
)

type Server struct {
	sock mangos.Socket
	addr string
	out  chan []byte // TODO: checkout whether fixed size is better
}

func NewServer(addr string, out chan []byte) *Server {
	sock, err := rep.NewSocket()
	if err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	return &Server{sock, addr, out}
}

func (sr *Server) Respond(payload []byte) error {
	return sr.sock.Send(payload)
}

func Listen(sr *Server) {
	if err := sr.sock.Listen(sr.addr); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}
	for {
		msg, err := sr.sock.Recv()
		// TODO: probably shouldn't die if can't read a socket
		if err != nil {
			die("cannot receive on rep socket", err.Error())
		}

		sr.out <- msg
	}
}
