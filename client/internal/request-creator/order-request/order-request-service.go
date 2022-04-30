package order_request

import (
	"io"
	"strconv"
	"time"

	nos "github.com/matf-pp/2022_MATDAQ/client/pkg/new-order-single"
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
func parseOrderPrice(orderType nos.OrderTypeReqEnum, price string) (float64, error) {
	if orderType == nos.OrderTypeReq.MarketOrder {
		return 0, nil
	}
	priceVal, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, err
	}
	return priceVal, nil
}

func parseOrderAmount(amount string) (uint32, error) {
	amountVal, err := strconv.ParseUint(amount, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(amountVal), nil
}

func parseOrder(securityId int32, orderType string, side string, price string, amount string) (nos.NewOrderSingle, error) {
	var timeInForce nos.TimeInForceEnum = nos.TimeInForce.GTC
	var ordTypeVal nos.OrderTypeReqEnum = parseOrderType(orderType)
	var sideVal nos.SideEnum = parseOrderSide(side)
	var sendingTime uint64 = uint64(time.Now().UnixNano())
	var amountVal uint32
	var priceVal float64
	var err error
	if amountVal, err = parseOrderAmount(amount); err != nil {
		return nos.NewOrderSingle{}, err
	}
	if priceVal, err = parseOrderPrice(ordTypeVal, price); err != nil {
		return nos.NewOrderSingle{}, err
	}

	return nos.NewOrderSingle{
		Price:                priceVal,
		OrderQty:             amountVal,
		SecurityID:           securityId,
		Side:                 sideVal,
		SeqNum:               0,
		SenderID:             [20]byte{},
		ClOrdID:              [20]byte{},
		OrderRequestID:       0,
		SendingTimeEpoch:     sendingTime,
		OrdType:              ordTypeVal,
		TimeInForce:          timeInForce,
		ManualOrderIndicator: 1,
	}, nil
}

func SendOrder(conn io.Writer, securityId int32, orderType string, side string, price string, amount string) error {
	newOrderData, err := parseOrder(securityId, orderType, side, price, amount)
	if err != nil {
		return err
	}

	m := nos.NewSbeGoMarshaller()

	header := nos.SbeGoMessageHeader{
		BlockLength: newOrderData.SbeBlockLength(),
		TemplateId:  newOrderData.SbeTemplateId(),
		SchemaId:    newOrderData.SbeSchemaId(),
		Version:     newOrderData.SbeSchemaVersion()}
	header.Encode(m, conn)

	if err = newOrderData.Encode(m, conn, false); err != nil {
		return err
	}

	return nil
}
