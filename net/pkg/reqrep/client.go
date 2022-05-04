package reqrep

import (
	"fmt"
	"os"

	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/req"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

type Client struct {
	sock mangos.Socket
	addr string
}

func newClient(addr string) *Client {
	sock, err := req.NewSocket()
	if err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	return &Client{sock, addr}
}

func (cl *Client) SendRequest(data []byte) error {
	return cl.sock.Send(data)
}

func (cl *Client) Close() {
	cl.sock.Close()
}

func die(format string, v ...any) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}
