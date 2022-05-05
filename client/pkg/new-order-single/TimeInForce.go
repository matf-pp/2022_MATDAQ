// Generated SBE (Simple Binary Encoding) message codec

package baseline

import (
	"fmt"
	"io"
	"reflect"
)

type TimeInForceEnum uint8
type TimeInForceValues struct {
	Day       TimeInForceEnum
	GTC       TimeInForceEnum
	FAK       TimeInForceEnum
	FOK       TimeInForceEnum
	GTD       TimeInForceEnum
	GFS       TimeInForceEnum
	NullValue TimeInForceEnum
}

var TimeInForce = TimeInForceValues{0, 1, 3, 4, 6, 99, 255}

func (t TimeInForceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(t)); err != nil {
		return err
	}
	return nil
}

func (t *TimeInForceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(t)); err != nil {
		return err
	}
	return nil
}

func (t TimeInForceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TimeInForce)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TimeInForce, unknown enumeration value %d", t)
}

func (*TimeInForceEnum) EncodedLength() int64 {
	return 1
}

func (*TimeInForceEnum) DaySinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.DaySinceVersion()
}

func (*TimeInForceEnum) DayDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GTCSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GTCInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GTCSinceVersion()
}

func (*TimeInForceEnum) GTCDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FAKSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FAKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FAKSinceVersion()
}

func (*TimeInForceEnum) FAKDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FOKSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FOKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FOKSinceVersion()
}

func (*TimeInForceEnum) FOKDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GTDSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GTDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GTDSinceVersion()
}

func (*TimeInForceEnum) GTDDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GFSSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GFSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GFSSinceVersion()
}

func (*TimeInForceEnum) GFSDeprecated() uint16 {
	return 0
}
