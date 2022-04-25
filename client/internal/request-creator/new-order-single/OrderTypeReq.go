// Generated SBE (Simple Binary Encoding) message codec

package new_order_single

import (
	"fmt"
	"io"
	"reflect"
)

type OrderTypeReqEnum uint8
type OrderTypeReqValues struct {
	MarketOrder    OrderTypeReqEnum
	LimitOrder     OrderTypeReqEnum
	StopOrder      OrderTypeReqEnum
	StopLimitOrder OrderTypeReqEnum
	NullValue      OrderTypeReqEnum
}

var OrderTypeReq = OrderTypeReqValues{1, 2, 3, 4, 255}

func (o OrderTypeReqEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderTypeReqEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderTypeReqEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderTypeReq)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderTypeReq, unknown enumeration value %d", o)
}

func (*OrderTypeReqEnum) EncodedLength() int64 {
	return 1
}

func (*OrderTypeReqEnum) MarketOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) MarketOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketOrderSinceVersion()
}

func (*OrderTypeReqEnum) MarketOrderDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) LimitOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) LimitOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitOrderSinceVersion()
}

func (*OrderTypeReqEnum) LimitOrderDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) StopOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) StopOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopOrderSinceVersion()
}

func (*OrderTypeReqEnum) StopOrderDeprecated() uint16 {
	return 0
}

func (*OrderTypeReqEnum) StopLimitOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderTypeReqEnum) StopLimitOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopLimitOrderSinceVersion()
}

func (*OrderTypeReqEnum) StopLimitOrderDeprecated() uint16 {
	return 0
}
