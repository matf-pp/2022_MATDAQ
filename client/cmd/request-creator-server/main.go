package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	api "github.com/matf-pp/2022_MATDAQ/api/matching-engine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	new_order_single "github.com/matf-pp/2022_MATDAQ/client/pkg/new-order-single"
)

const HOST_NAME string = "request-creator-server:8081"

func handleConnection(m *new_order_single.SbeGoMarshaller, conn net.Conn, client api.MatchingEngineClient) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	for {
		var hdr new_order_single.SbeGoMessageHeader
		hdr.Decode(m, conn)

		fmt.Println("reading new order single")
		var newOrderData new_order_single.NewOrderSingle
		if err := newOrderData.Decode(m, conn, hdr.Version, hdr.BlockLength, false); err != nil {
			fmt.Println("Order for NewOrderSingle failed.")
			break
		}
		//	  Send request to Matching Engine

		orderTypeId := func() int {
			if newOrderData.OrdType == new_order_single.OrderTypeReq.LimitOrder {
				return 0
			} else {
				return 1
			}
		}()

		orderSideId := func() int {
			if newOrderData.Side == new_order_single.Side.Buy {
				return 0
			} else {
				return 1
			}
		}()

		loginUserReq := &api.CreateOrderRequest{
			SecurityOrder: &api.SecurityOrder{
				Price:         newOrderData.Price,
				SecurityId:    uint32(newOrderData.SecurityID),
				OrderQuantity: newOrderData.OrderQty,
				OrderSide:     api.SecurityOrder_OrderSide(orderSideId),
			},
			OrderId:   rand.Uint64(),
			OrderType: api.CreateOrderRequest_OrderType(orderTypeId),
			SenderId:  newOrderData.SenderID[:],
		}

		_, err := client.CreateOrder(context.Background(), loginUserReq)
		if err != nil {
			log.Fatalf("Error when calling CreateOrder: %s", err)
		}

		fmt.Println(newOrderData)
	}
	conn.Close()
}

func createMatchingEngineClient() *api.MatchingEngineClient {
	const GRPC_PORT = 10000
	var conn *grpc.ClientConn
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf(":%d", GRPC_PORT), opts...)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	matchingEngineClient := api.NewMatchingEngineClient(conn)

	return &matchingEngineClient
}

// this is only example server, it will be removed later
func main() {
	matchingEngineClient := createMatchingEngineClient()
	listener, err := net.Listen("tcp", HOST_NAME)
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
		go handleConnection(m, conn, *matchingEngineClient)
	}
}
