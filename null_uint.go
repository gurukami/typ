package typ

import (
	"database/sql/driver"
	"encoding/json"
)

// UintCommon represents an uint with pointer and error.
type UintCommon struct {
	P     *uint
	Error error
}

// Set saves value into current struct
func (n *UintCommon) Set(value uint) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n UintCommon) V() uint {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n UintCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n UintCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n UintCommon) Value() (driver.Value, error) {
	nv := UintInt64(uint64(n.V()))
	return nv.V(), nv.Err()
}

// Scan implements the sql Scanner interface.
func (n *UintCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *UintCommon) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Err()
	}
	v := FloatUint(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n UintCommon) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	if v.Err() != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n UintCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n UintCommon) Err() error {
	return n.Error
}

// UintAccessor
type UintAccessor interface {
	NullCommon
	V() uint
	Set(value uint)
	Clone() UintAccessor
}

// NullUint represents an uint that may be null.
type NullUint struct {
	UintCommon
}

// Value implements the sql driver Valuer interface.
func (n NullUint) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.UintCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.UintCommon.MarshalJSON()
}

// Clone returns new instance of NullUint with preserved value & error
func (n NullUint) Clone() UintAccessor {
	nv := &NullUint{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NUint returns NullUint under UintAccessor from uint
func NUint(value uint) UintAccessor {
	return &NullUint{UintCommon{P: &value}}
}

// NotNullUint represents an uint with accessor.
type NotNullUint struct {
	UintCommon
}

// Clone returns new instance of NotNullUint with preserved value & error
func (n NotNullUint) Clone() UintAccessor {
	nv := &NotNullUint{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNUint returns NotNullUint under UintAccessor from uint
func NNUint(value uint) UintAccessor {
	return &NotNullUint{UintCommon{P: &value}}
}

// UintSlice returns slice of uint with filled values from slice of UintAccessor
func UintSlice(null []UintAccessor, valid bool) []uint {
	slice := make([]uint, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Uint8Common represents an uint8 with pointer and error.
type Uint8Common struct {
	P     *uint8
	Error error
}

// Set saves value into current struct
func (n *Uint8Common) Set(value uint8) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Uint8Common) V() uint8 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Uint8Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Uint8Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Uint8Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Uint8Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint8()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Uint8Common) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Err()
	}
	v := FloatUint8(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Uint8Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Uint8Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Uint8Common) Err() error {
	return n.Error
}

// Uint8Accessor
type Uint8Accessor interface {
	NullCommon
	V() uint8
	Set(value uint8)
	Clone() Uint8Accessor
}

// NullUint8 represents an uint8 that may be null.
type NullUint8 struct {
	Uint8Common
}

// Value implements the sql driver Valuer interface.
func (n NullUint8) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Uint8Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint8) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Uint8Common.MarshalJSON()
}

// Clone returns new instance of NullUint8 with preserved value & error
func (n NullUint8) Clone() Uint8Accessor {
	nv := &NullUint8{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NUint8 returns NullUint8 under Uint8Accessor from uint8
func NUint8(value uint8) Uint8Accessor {
	return &NullUint8{Uint8Common{P: &value}}
}

// NotNullUint represents an uint8 with accessor.
type NotNullUint8 struct {
	Uint8Common
}

// Clone returns new instance of NotNullUint8 with preserved value & error
func (n NotNullUint8) Clone() Uint8Accessor {
	nv := &NotNullUint8{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNUint8 returns NotNullUint8 under Uint8Accessor from uint8
func NNUint8(value uint8) Uint8Accessor {
	return &NotNullUint8{Uint8Common{P: &value}}
}

// Uint8Slice returns slice of uint8 with filled values from slice of Uint8Accessor
func Uint8Slice(null []Uint8Accessor, valid bool) []uint8 {
	slice := make([]uint8, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Uint16Common represents an uint16 with pointer and error.
type Uint16Common struct {
	P     *uint16
	Error error
}

// Set saves value into current struct
func (n *Uint16Common) Set(value uint16) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Uint16Common) V() uint16 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Uint16Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Uint16Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Uint16Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Uint16Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint16()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Uint16Common) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Err()
	}
	v := FloatUint16(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Uint16Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Uint16Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Uint16Common) Err() error {
	return n.Error
}

// Uint16Accessor
type Uint16Accessor interface {
	NullCommon
	V() uint16
	Set(value uint16)
	Clone() Uint16Accessor
}

// NullUint16 represents an uint16 that may be null.
type NullUint16 struct {
	Uint16Common
}

// Value implements the sql driver Valuer interface.
func (n NullUint16) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Uint16Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint16) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Uint16Common.MarshalJSON()
}

// Clone returns new instance of NullUint16 with preserved value & error
func (n NullUint16) Clone() Uint16Accessor {
	nv := &NullUint16{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NUint16 returns NullUint16 under Uint16Accessor from uint16
func NUint16(value uint16) Uint16Accessor {
	return &NullUint16{Uint16Common{P: &value}}
}

// NotNullUint16 represents an uint16 with accessor.
type NotNullUint16 struct {
	Uint16Common
}

// Clone returns new instance of NotNullUint16 with preserved value & error
func (n NotNullUint16) Clone() Uint16Accessor {
	nv := &NotNullUint16{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNUint16 returns NotNullUint16 under Uint16Accessor from uint16
func NNUint16(value uint16) Uint16Accessor {
	return &NotNullUint16{Uint16Common{P: &value}}
}

// Uint16Slice returns slice of uint16 with filled values from slice of Uint16Accessor
func Uint16Slice(null []Uint16Accessor, valid bool) []uint16 {
	slice := make([]uint16, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Uint32Common represents an uint32 with pointer and error.
type Uint32Common struct {
	P     *uint32
	Error error
}

// Set saves value into current struct
func (n *Uint32Common) Set(value uint32) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Uint32Common) V() uint32 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Uint32Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Uint32Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Uint32Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Uint32Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint32()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Uint32Common) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Err()
	}
	v := FloatUint32(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Uint32Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Uint32Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Uint32Common) Err() error {
	return n.Error
}

// Uint32Accessor
type Uint32Accessor interface {
	NullCommon
	V() uint32
	Set(value uint32)
	Clone() Uint32Accessor
}

// NullUint32 represents an uint32 that may be null.
type NullUint32 struct {
	Uint32Common
}

// Value implements the sql driver Valuer interface.
func (n NullUint32) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Uint32Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint32) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Uint32Common.MarshalJSON()
}

// Clone returns new instance of NullUint32 with preserved value & error
func (n NullUint32) Clone() Uint32Accessor {
	nv := &NullUint32{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NUint32 returns NullUint32 under Uint32Accessor from uint32
func NUint32(value uint32) Uint32Accessor {
	return &NullUint32{Uint32Common{P: &value}}
}

// NotNullUint32 represents an uint32 with accessor.
type NotNullUint32 struct {
	Uint32Common
}

// Clone returns new instance of NotNullUint32 with preserved value & error
func (n NotNullUint32) Clone() Uint32Accessor {
	nv := &NotNullUint32{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNUint32 returns NotNullUint32 under Uint32Accessor from uint32
func NNUint32(value uint32) Uint32Accessor {
	return &NotNullUint32{Uint32Common{P: &value}}
}

// Uint32Slice returns slice of uint32 with filled values from slice of Uint32Accessor
func Uint32Slice(null []Uint32Accessor, valid bool) []uint32 {
	slice := make([]uint32, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Uint64Common represents an uint64 with pointer and error.
type Uint64Common struct {
	P     *uint64
	Error error
}

// Set saves value into current struct
func (n *Uint64Common) Set(value uint64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Uint64Common) V() uint64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Uint64Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Uint64Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Uint64Common) Value() (driver.Value, error) {
	nv := UintInt64(n.V())
	return nv.V(), nv.Err()
}

// Scan implements the sql Scanner interface.
func (n *Uint64Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint64()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Uint64Common) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Err()
	}
	v := FloatUint64(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Uint64Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	if v.Err() != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Uint64Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Uint64Common) Err() error {
	return n.Error
}

// Uint64Accessor
type Uint64Accessor interface {
	NullCommon
	V() uint64
	Set(value uint64)
	Clone() Uint64Accessor
}

// NullUint64 represents an uint64 that may be null.
type NullUint64 struct {
	Uint64Common
}

// Value implements the sql driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Uint64Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint64) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Uint64Common.MarshalJSON()
}

// Clone returns new instance of NullUint64 with preserved value & error
func (n NullUint64) Clone() Uint64Accessor {
	nv := &NullUint64{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NUint64 returns NullUint64 under Uint64Accessor from uint64
func NUint64(value uint64) Uint64Accessor {
	return &NullUint64{Uint64Common{P: &value}}
}

// NotNullUint64 represents an uint64 with accessor.
type NotNullUint64 struct {
	Uint64Common
}

// Clone returns new instance of NotNullUint64 with preserved value & error
func (n NotNullUint64) Clone() Uint64Accessor {
	nv := &NotNullUint64{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNUint64 returns NotNullUint64 under Uint64Accessor from uint64
func NNUint64(value uint64) Uint64Accessor {
	return &NotNullUint64{Uint64Common{P: &value}}
}

// Uint64Slice returns slice of uint64 with filled values from slice of Uint64Accessor
func Uint64Slice(null []Uint64Accessor, valid bool) []uint64 {
	slice := make([]uint64, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
