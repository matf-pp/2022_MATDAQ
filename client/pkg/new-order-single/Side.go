// Generated SBE (Simple Binary Encoding) message codec

package baseline

import (
	"fmt"
	"io"
	"reflect"
)

type SideEnum uint8
type SideValues struct {
	Buy       SideEnum
	Sell      SideEnum
	NullValue SideEnum
}

var Side = SideValues{1, 2, 255}

func (s SideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(Side)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on Side, unknown enumeration value %d", s)
}

func (*SideEnum) EncodedLength() int64 {
	return 1
}

func (*SideEnum) BuySinceVersion() uint16 {
	return 0
}

func (s *SideEnum) BuyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BuySinceVersion()
}

func (*SideEnum) BuyDeprecated() uint16 {
	return 0
}

func (*SideEnum) SellSinceVersion() uint16 {
	return 0
}

func (s *SideEnum) SellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SellSinceVersion()
}

func (*SideEnum) SellDeprecated() uint16 {
	return 0
}
