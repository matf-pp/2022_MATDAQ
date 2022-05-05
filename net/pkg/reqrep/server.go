package reqrep

import (
	"fmt"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/rep"
)

type Server struct {
	sock mangos.Socket
	addr string
	out  chan Serializable
}

func NewServer(addr string, out chan Serializable) *Server {
	sock, err := rep.NewSocket()
	if err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	return &Server{sock, addr, out}
}

func (sr *Server) Respond(s Serializable) error {
	return sr.sock.Send(s.ToBytes())
}

func Listen[S Serializable](sr *Server) {
	if err := sr.sock.Listen(sr.addr); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}
	for {
		msg, err := sr.sock.Recv()
		// TODO: probably shouldn't die if can't read a socket
		if err != nil {
			die("cannot receive on rep socket", err.Error())
		}

		// TODO: check if the interface should be implemented via pointers
		s := new(S)
		(*s).FromBytes(msg)

		fmt.Println("servermsg", *s)

		sr.out <- *s
	}
}
