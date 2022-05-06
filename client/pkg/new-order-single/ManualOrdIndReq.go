// Generated SBE (Simple Binary Encoding) message codec

package baseline

import (
	"fmt"
	"io"
	"reflect"
)

type ManualOrdIndReqEnum uint8
type ManualOrdIndReqValues struct {
	Automated ManualOrdIndReqEnum
	Manual    ManualOrdIndReqEnum
	NullValue ManualOrdIndReqEnum
}

var ManualOrdIndReq = ManualOrdIndReqValues{0, 1, 255}

func (m ManualOrdIndReqEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *ManualOrdIndReqEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m ManualOrdIndReqEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ManualOrdIndReq)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ManualOrdIndReq, unknown enumeration value %d", m)
}

func (*ManualOrdIndReqEnum) EncodedLength() int64 {
	return 1
}

func (*ManualOrdIndReqEnum) AutomatedSinceVersion() uint16 {
	return 0
}

func (m *ManualOrdIndReqEnum) AutomatedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AutomatedSinceVersion()
}

func (*ManualOrdIndReqEnum) AutomatedDeprecated() uint16 {
	return 0
}

func (*ManualOrdIndReqEnum) ManualSinceVersion() uint16 {
	return 0
}

func (m *ManualOrdIndReqEnum) ManualInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ManualSinceVersion()
}

func (*ManualOrdIndReqEnum) ManualDeprecated() uint16 {
	return 0
}
