package reqrep

import (
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/protocol/req"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

type Client struct {
	sock mangos.Socket
	addr string
}

func NewClient(addr string) *Client {
	sock, err := req.NewSocket()
	if err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	if err = sock.Dial(addr); err != nil {
		die("can't dial on req socket: %s", err.Error())
	}
	return &Client{sock, addr}
}

func SendRequest(cl *Client, data Serializable) (Serializable, error) {
	if err := cl.sock.Send(data.ToBytes()); err != nil {
		return nil, err
	}
	resp, err := cl.sock.Recv()
	if err != nil {
		return nil, err
	}
	// kinda hacky to use the same var, ali takva mi tura
	data.FromBytes(resp)
	res := data
	return res, nil
}

func (cl *Client) Close() {
	cl.sock.Close()
}
