package matching_engine

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	api "github.com/matf-pp/2022_MATDAQ/api/matching-engine"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

const PORT = 10000

type OrderResponse = api.PublishOrderResponse
type TradeResponse = api.PublishTradeResponse

func StartMatchingEngine(orderResponses chan *OrderResponse, tradeResponses chan *TradeResponse) {
	var conn *grpc.ClientConn
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// conn, err := grpc.Dial(fmt.Sprintf("matching-engine:%d", PORT), opts...)
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", PORT), opts...)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	matchingEngineClient := api.NewMatchingEngineClient(conn)

	publishOrderRequest := &api.PublishOrderRequest{}
	orderResponseStream, err := matchingEngineClient.PublishOrderCreation(context.Background(), publishOrderRequest)
	if err != nil {
		log.Fatalf("Error when calling PublishOrderCreation: %s", err)
	}
	go func() {
		for {
			orderResponse, err := orderResponseStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.ListFeatures(_) = _, %v", matchingEngineClient, err)
			}
			log.Println(orderResponse)
			orderResponses <- orderResponse
		}
	}()

	publishTradeRequest := &api.PublishTradeRequest{}
	tradeResponseStream, err := matchingEngineClient.PublishTrade(context.Background(), publishTradeRequest)
	if err != nil {
		log.Fatalf("Error when calling PublishTrade: %s", err)
	}
	go func() {
		for {
			tradeResponse, err := tradeResponseStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.ListFeatures(_) = _, %v", matchingEngineClient, err)
			}
			log.Println(tradeResponse)
			tradeResponses <- tradeResponse
		}
	}()
}

func HandleBubbleTea(p *tea.Program, orderResponses chan *OrderResponse, tradeResponses chan *TradeResponse) {
	for {
		select {
		case orderResponse := <-orderResponses:
			fmt.Println("handle", orderResponse)
			p.Send(orderResponse)
		case tradeResponse := <-tradeResponses:
			fmt.Println("handle", tradeResponse)
			p.Send(tradeResponse)
		}
	}
}
