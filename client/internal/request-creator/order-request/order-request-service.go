package order_request

import (
	"fmt"
	nos "github.com/matf-pp/2022_MATDAQ/internal/request-creator/new-order-single"
	"io"
	"os"
	"strconv"
)

func parseOrderType(orderType string) nos.OrderTypeReqEnum {
	if orderType == "Market Order" {
		return nos.OrderTypeReq.MarketOrder
	}
	return nos.OrderTypeReq.LimitOrder
}

func parseOrderSide(side string) nos.SideEnum {
	if side == "Buy" {
		return nos.Side.Buy
	}
	return nos.Side.Sell
}

// FIX: should not return 0 when the order type is Market Order
func parseOrderPrice(orderType nos.OrderTypeReqEnum, price string) float64 {
	if orderType == nos.OrderTypeReq.MarketOrder {
		return 0
	}
	priceVal, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return priceVal
}

func parseOrderAmount(amount string) uint32 {
	amountVal, err := strconv.ParseUint(amount, 10, 32)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return uint32(amountVal)
}

func parseOrder(orderType string, side string, price string, amount string) nos.NewOrderSingle {
	var timeInForce nos.TimeInForceEnum = nos.TimeInForce.GTC
	var ordTypeVal nos.OrderTypeReqEnum = parseOrderType(orderType)
	var sideVal nos.SideEnum = parseOrderSide(side)
	var amountVal uint32 = parseOrderAmount(amount)
	var priceVal float64 = parseOrderPrice(ordTypeVal, price)

	return nos.NewOrderSingle{
		Price:                priceVal,
		OrderQty:             amountVal,
		SecurityID:           1,
		Side:                 sideVal,
		SeqNum:               0,
		SenderID:             [20]byte{},
		ClOrdID:              [20]byte{},
		OrderRequestID:       0,
		SendingTimeEpoch:     0,
		OrdType:              ordTypeVal,
		TimeInForce:          timeInForce,
		ManualOrderIndicator: 0,
	}
}

func SendOrder(conn io.Writer, orderType string, side string, price string, amount string) {
	newOrderData := parseOrder(orderType, side, price, amount)

	// this can maybe be inside of model
	m := nos.NewSbeGoMarshaller()

	header := nos.SbeGoMessageHeader{newOrderData.SbeBlockLength(), newOrderData.SbeTemplateId(), newOrderData.SbeSchemaId(), newOrderData.SbeSchemaVersion()}
	header.Encode(m, conn)

	err := newOrderData.Encode(m, conn, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
