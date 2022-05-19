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

const HostName string = "request-creator-server:8081"

func handleConnection(marshaller *new_order_single.SbeGoMarshaller, conn net.Conn, client api.MatchingEngineClient) {
	log.Printf("New connection: %s\n", conn.RemoteAddr().String())
	for {
		var hdr new_order_single.SbeGoMessageHeader
		if err := hdr.Decode(marshaller, conn); err != nil {
			log.Println("Failed decoding")
			break
		}

		var newOrderData new_order_single.NewOrderSingle
		if err := newOrderData.Decode(marshaller, conn, hdr.Version, hdr.BlockLength, false); err != nil {
			fmt.Println("Reading NewOrderSingle failed.")
			break
		}
		log.Println("Reading NewOrderSingle successful")

		// Prepare data to be sent to the Matching Engine
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
			log.Printf("Error when calling CreateOrder: %s\n", err)
		}
		log.Println("Order created with data: ", newOrderData)
	}
	conn.Close()
}

func createMatchingEngineClient() *api.MatchingEngineClient {
	const GRPC_PORT = 10000
	var conn *grpc.ClientConn
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("matching-engine:%d", GRPC_PORT), opts...)
	if err != nil {
		log.Fatalf("Did not connect to matching-engine: %s", err)
	}
	matchingEngineClient := api.NewMatchingEngineClient(conn)

	return &matchingEngineClient
}

func main() {
	matchingEngineClient := createMatchingEngineClient()
	listener, err := net.Listen("tcp", HostName)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	defer listener.Close()

	marshaller := new_order_single.NewSbeGoMarshaller()

	log.Printf("Listening on %s\n", HostName)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(marshaller, conn, *matchingEngineClient)
	}
}
