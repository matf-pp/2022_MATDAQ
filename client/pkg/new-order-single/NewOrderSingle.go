// Generated SBE (Simple Binary Encoding) message codec

package baseline

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NewOrderSingle struct {
	Price                int32
	OrderQty             uint32
	SecurityID           int32
	Side                 SideEnum
	SeqNum               uint32
	SenderID             [20]byte
	ClOrdID              [20]byte
	OrderRequestID       uint64
	SendingTimeEpoch     uint64
	OrdType              OrderTypeReqEnum
	TimeInForce          TimeInForceEnum
	ManualOrderIndicator ManualOrdIndReqEnum
}

func (n *NewOrderSingle) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, n.Price); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, n.SecurityID); err != nil {
		return err
	}
	if err := n.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, n.SeqNum); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.SenderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.OrderRequestID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.SendingTimeEpoch); err != nil {
		return err
	}
	if err := n.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ManualOrderIndicator.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderSingle) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.PriceInActingVersion(actingVersion) {
		n.Price = n.PriceNullValue()
	} else {
		if err := _m.ReadInt32(_r, &n.Price); err != nil {
			return err
		}
	}
	if !n.OrderQtyInActingVersion(actingVersion) {
		n.OrderQty = n.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.OrderQty); err != nil {
			return err
		}
	}
	if !n.SecurityIDInActingVersion(actingVersion) {
		n.SecurityID = n.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &n.SecurityID); err != nil {
			return err
		}
	}
	if n.SideInActingVersion(actingVersion) {
		if err := n.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.SeqNumInActingVersion(actingVersion) {
		n.SeqNum = n.SeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &n.SeqNum); err != nil {
			return err
		}
	}
	if !n.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.SenderID[idx] = n.SenderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.SenderID[:]); err != nil {
			return err
		}
	}
	if !n.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.ClOrdID[idx] = n.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !n.OrderRequestIDInActingVersion(actingVersion) {
		n.OrderRequestID = n.OrderRequestIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.OrderRequestID); err != nil {
			return err
		}
	}
	if !n.SendingTimeEpochInActingVersion(actingVersion) {
		n.SendingTimeEpoch = n.SendingTimeEpochNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.SendingTimeEpoch); err != nil {
			return err
		}
	}
	if n.OrdTypeInActingVersion(actingVersion) {
		if err := n.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.TimeInForceInActingVersion(actingVersion) {
		if err := n.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ManualOrderIndicatorInActingVersion(actingVersion) {
		if err := n.ManualOrderIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NewOrderSingle) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.PriceInActingVersion(actingVersion) {
		if n.Price < n.PriceMinValue() || n.Price > n.PriceMaxValue() {
			return fmt.Errorf("Range check failed on n.Price (%v < %v > %v)", n.PriceMinValue(), n.Price, n.PriceMaxValue())
		}
	}
	if n.OrderQtyInActingVersion(actingVersion) {
		if n.OrderQty < n.OrderQtyMinValue() || n.OrderQty > n.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderQty (%v < %v > %v)", n.OrderQtyMinValue(), n.OrderQty, n.OrderQtyMaxValue())
		}
	}
	if n.SecurityIDInActingVersion(actingVersion) {
		if n.SecurityID < n.SecurityIDMinValue() || n.SecurityID > n.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on n.SecurityID (%v < %v > %v)", n.SecurityIDMinValue(), n.SecurityID, n.SecurityIDMaxValue())
		}
	}
	if err := n.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.SeqNumInActingVersion(actingVersion) {
		if n.SeqNum < n.SeqNumMinValue() || n.SeqNum > n.SeqNumMaxValue() {
			return fmt.Errorf("Range check failed on n.SeqNum (%v < %v > %v)", n.SeqNumMinValue(), n.SeqNum, n.SeqNumMaxValue())
		}
	}
	if n.SenderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.SenderID[idx] < n.SenderIDMinValue() || n.SenderID[idx] > n.SenderIDMaxValue() {
				return fmt.Errorf("Range check failed on n.SenderID[%d] (%v < %v > %v)", idx, n.SenderIDMinValue(), n.SenderID[idx], n.SenderIDMaxValue())
			}
		}
	}
	for idx, ch := range n.SenderID {
		if ch > 127 {
			return fmt.Errorf("n.SenderID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if n.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.ClOrdID[idx] < n.ClOrdIDMinValue() || n.ClOrdID[idx] > n.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on n.ClOrdID[%d] (%v < %v > %v)", idx, n.ClOrdIDMinValue(), n.ClOrdID[idx], n.ClOrdIDMaxValue())
			}
		}
	}
	for idx, ch := range n.ClOrdID {
		if ch > 127 {
			return fmt.Errorf("n.ClOrdID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if n.OrderRequestIDInActingVersion(actingVersion) {
		if n.OrderRequestID < n.OrderRequestIDMinValue() || n.OrderRequestID > n.OrderRequestIDMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderRequestID (%v < %v > %v)", n.OrderRequestIDMinValue(), n.OrderRequestID, n.OrderRequestIDMaxValue())
		}
	}
	if n.SendingTimeEpochInActingVersion(actingVersion) {
		if n.SendingTimeEpoch < n.SendingTimeEpochMinValue() || n.SendingTimeEpoch > n.SendingTimeEpochMaxValue() {
			return fmt.Errorf("Range check failed on n.SendingTimeEpoch (%v < %v > %v)", n.SendingTimeEpochMinValue(), n.SendingTimeEpoch, n.SendingTimeEpochMaxValue())
		}
	}
	if err := n.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.ManualOrderIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func NewOrderSingleInit(n *NewOrderSingle) {
	return
}

func (*NewOrderSingle) SbeBlockLength() (blockLength uint16) {
	return 76
}

func (*NewOrderSingle) SbeTemplateId() (templateId uint16) {
	return 514
}

func (*NewOrderSingle) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*NewOrderSingle) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*NewOrderSingle) SbeSemanticType() (semanticType []byte) {
	return []byte("D")
}

func (*NewOrderSingle) PriceId() uint16 {
	return 44
}

func (*NewOrderSingle) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrderSingle) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) PriceMinValue() int32 {
	return math.MinInt32 + 1
}

func (*NewOrderSingle) PriceMaxValue() int32 {
	return math.MaxInt32
}

func (*NewOrderSingle) PriceNullValue() int32 {
	return math.MinInt32
}

func (*NewOrderSingle) OrderQtyId() uint16 {
	return 38
}

func (*NewOrderSingle) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrderSingle) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) OrderQtyMinValue() uint32 {
	return 0
}

func (*NewOrderSingle) OrderQtyMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderSingle) OrderQtyNullValue() uint32 {
	return math.MaxUint32
}

func (*NewOrderSingle) SecurityIDId() uint16 {
	return 48
}

func (*NewOrderSingle) SecurityIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SecurityIDSinceVersion()
}

func (*NewOrderSingle) SecurityIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SecurityIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*NewOrderSingle) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*NewOrderSingle) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*NewOrderSingle) SideId() uint16 {
	return 54
}

func (*NewOrderSingle) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrderSingle) SideDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) SeqNumId() uint16 {
	return 9726
}

func (*NewOrderSingle) SeqNumSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SeqNumSinceVersion()
}

func (*NewOrderSingle) SeqNumDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SeqNumMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) SeqNumMinValue() uint32 {
	return 0
}

func (*NewOrderSingle) SeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*NewOrderSingle) SeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*NewOrderSingle) SenderIDId() uint16 {
	return 5392
}

func (*NewOrderSingle) SenderIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SenderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SenderIDSinceVersion()
}

func (*NewOrderSingle) SenderIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SenderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) SenderIDMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle) SenderIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle) SenderIDNullValue() byte {
	return 0
}

func (n *NewOrderSingle) SenderIDCharacterEncoding() string {
	return "ASCII"
}

func (*NewOrderSingle) ClOrdIDId() uint16 {
	return 11
}

func (*NewOrderSingle) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrderSingle) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) ClOrdIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle) ClOrdIDNullValue() byte {
	return 0
}

func (n *NewOrderSingle) ClOrdIDCharacterEncoding() string {
	return "ASCII"
}

func (*NewOrderSingle) OrderRequestIDId() uint16 {
	return 2422
}

func (*NewOrderSingle) OrderRequestIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrderRequestIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderRequestIDSinceVersion()
}

func (*NewOrderSingle) OrderRequestIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrderRequestIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) OrderRequestIDMinValue() uint64 {
	return 0
}

func (*NewOrderSingle) OrderRequestIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle) OrderRequestIDNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle) SendingTimeEpochId() uint16 {
	return 5297
}

func (*NewOrderSingle) SendingTimeEpochSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SendingTimeEpochInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SendingTimeEpochSinceVersion()
}

func (*NewOrderSingle) SendingTimeEpochDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SendingTimeEpochMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) SendingTimeEpochMinValue() uint64 {
	return 0
}

func (*NewOrderSingle) SendingTimeEpochMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle) SendingTimeEpochNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle) OrdTypeId() uint16 {
	return 40
}

func (*NewOrderSingle) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrderSingle) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrdTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) TimeInForceId() uint16 {
	return 59
}

func (*NewOrderSingle) TimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TimeInForceSinceVersion()
}

func (*NewOrderSingle) TimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TimeInForceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*NewOrderSingle) ManualOrderIndicatorId() uint16 {
	return 1028
}

func (*NewOrderSingle) ManualOrderIndicatorSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) ManualOrderIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ManualOrderIndicatorSinceVersion()
}

func (*NewOrderSingle) ManualOrderIndicatorDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) ManualOrderIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}
