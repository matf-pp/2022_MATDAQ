package main

import (
	"fmt"
	new_order_single "github.com/matf-pp/2022_MATDAQ/pkg/new-order-single"
	"net"
)

const PORT string = "127.0.0.1:8081"

func handleConnection(m *new_order_single.SbeGoMarshaller, conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		fmt.Println("reading message header")
		var hdr new_order_single.SbeGoMessageHeader
		hdr.Decode(m, conn)
		fmt.Println("message header read")

		fmt.Println("reading new order single")
		var newOrderData new_order_single.NewOrderSingle
		if err := newOrderData.Decode(m, conn, hdr.Version, hdr.BlockLength, false); err != nil {
			fmt.Println("Order for NewOrderSingle failed.")
			continue
		}

		fmt.Println(newOrderData)
	}
	conn.Close()
}

// this is only example server, it will be removed later
func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	m := new_order_single.NewSbeGoMarshaller()

	fmt.Println("listen")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(m, conn)
	}
}
