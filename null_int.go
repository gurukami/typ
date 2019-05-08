package typ

import (
	"database/sql/driver"
	"encoding/json"
)

// IntCommon represents an int with pointer and error.
type IntCommon struct {
	P     *int
	Error error
}

// Set saves value into current struct
func (n *IntCommon) Set(value int) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n IntCommon) V() int {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n IntCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n IntCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n IntCommon) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *IntCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *IntCommon) UnmarshalJSON(b []byte) error {
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
	v := FloatInt(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n IntCommon) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	if v.Err() != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n IntCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n IntCommon) Err() error {
	return n.Error
}

// IntAccessor accessor of int type.
type IntAccessor interface {
	Common
	V() int
	Set(value int)
	Clone() IntAccessor
}

// NullInt represents an int that may be null.
type NullInt struct {
	IntCommon
}

// Value implements the sql driver Valuer interface.
func (n NullInt) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.IntCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.IntCommon.MarshalJSON()
}

// Clone returns new instance of NullInt with preserved value & error
func (n NullInt) Clone() IntAccessor {
	nv := &NullInt{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NInt returns NullInt under IntAccessor from int
func NInt(value int) IntAccessor {
	return &NullInt{IntCommon{P: &value}}
}

// NotNullInt represents an int with accessor.
type NotNullInt struct {
	IntCommon
}

// Clone returns new instance of NotNullInt with preserved value & error
func (n NotNullInt) Clone() IntAccessor {
	nv := &NotNullInt{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNInt returns NotNullInt under IntAccessor from int
func NNInt(value int) IntAccessor {
	return &NotNullInt{IntCommon{P: &value}}
}

// IntSlice returns slice of int with filled values from slice of IntAccessor
func IntSlice(null []IntAccessor, valid bool) []int {
	slice := make([]int, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Int8Common represents an int8 with pointer and error.
type Int8Common struct {
	P     *int8
	Error error
}

// Set saves value into current struct
func (n *Int8Common) Set(value int8) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Int8Common) V() int8 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Int8Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Int8Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Int8Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Int8Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int8()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Int8Common) UnmarshalJSON(b []byte) error {
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
	v := FloatInt8(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Int8Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Int8Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Int8Common) Err() error {
	return n.Error
}

// Int8Accessor accessor of int8 type.
type Int8Accessor interface {
	Common
	V() int8
	Set(value int8)
	Clone() Int8Accessor
}

// NullInt8 represents an int8 that may be null.
type NullInt8 struct {
	Int8Common
}

// Value implements the sql driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Int8Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt8) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Int8Common.MarshalJSON()
}

// Clone returns new instance of NullInt8 with preserved value & error
func (n NullInt8) Clone() Int8Accessor {
	nv := &NullInt8{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NInt8 returns NullInt8 under Int8Accessor from int8
func NInt8(value int8) Int8Accessor {
	return &NullInt8{Int8Common{P: &value}}
}

// NotNullInt8 represents an int8 with accessor.
type NotNullInt8 struct {
	Int8Common
}

// Clone returns new instance of NotNullInt8 with preserved value & error
func (n NotNullInt8) Clone() Int8Accessor {
	nv := &NotNullInt8{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNInt8 returns NotNullInt8 under Int8Accessor from int8
func NNInt8(value int8) Int8Accessor {
	return &NotNullInt8{Int8Common{P: &value}}
}

// Int8Slice returns slice of int8 with filled values from slice of Int8Accessor
func Int8Slice(null []Int8Accessor, valid bool) []int8 {
	slice := make([]int8, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Int16Common represents an int16 with pointer and error.
type Int16Common struct {
	P     *int16
	Error error
}

// Set saves value into current struct
func (n *Int16Common) Set(value int16) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Int16Common) V() int16 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Int16Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Int16Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Int16Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Int16Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int16()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Int16Common) UnmarshalJSON(b []byte) error {
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
	v := FloatInt16(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Int16Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Int16Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Int16Common) Err() error {
	return n.Error
}

// Int16Accessor accessor of int16 type.
type Int16Accessor interface {
	Common
	V() int16
	Set(value int16)
	Clone() Int16Accessor
}

// NullInt16 represents an int16 that may be null.
type NullInt16 struct {
	Int16Common
}

// Value implements the sql driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Int16Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt16) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Int16Common.MarshalJSON()
}

// Clone returns new instance of NullInt16 with preserved value & error
func (n NullInt16) Clone() Int16Accessor {
	nv := &NullInt16{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NInt16 returns NullInt16 under Int16Accessor from int16
func NInt16(value int16) Int16Accessor {
	return &NullInt16{Int16Common{P: &value}}
}

// NotNullInt16 represents an int16 with accessor.
type NotNullInt16 struct {
	Int16Common
}

// Clone returns new instance of NotNullInt16 with preserved value & error
func (n NotNullInt16) Clone() Int16Accessor {
	nv := &NotNullInt16{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNInt16 returns NotNullInt16 under Int16Accessor from int16
func NNInt16(value int16) Int16Accessor {
	return &NotNullInt16{Int16Common{P: &value}}
}

// Int16Slice returns slice of int16 with filled values from slice of Int16Accessor
func Int16Slice(null []Int16Accessor, valid bool) []int16 {
	slice := make([]int16, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Int32Common represents an int32 with pointer and error.
type Int32Common struct {
	P     *int32
	Error error
}

// Set saves value into current struct
func (n *Int32Common) Set(value int32) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Int32Common) V() int32 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Int32Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Int32Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Int32Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Int32Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int32()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Int32Common) UnmarshalJSON(b []byte) error {
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
	v := FloatInt32(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Int32Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Int32Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Int32Common) Err() error {
	return n.Error
}

// Int32Accessor accessor of int32 type.
type Int32Accessor interface {
	Common
	V() int32
	Set(value int32)
	Clone() Int32Accessor
}

// NullInt32 represents an int32 that may be null.
type NullInt32 struct {
	Int32Common
}

// Value implements the sql driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Int32Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Int32Common.MarshalJSON()
}

// Clone returns new instance of NullInt32 with preserved value & error
func (n NullInt32) Clone() Int32Accessor {
	nv := &NullInt32{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NInt32 returns NullInt32 under Int32Accessor from int32
func NInt32(value int32) Int32Accessor {
	return &NullInt32{Int32Common{P: &value}}
}

// NotNullInt32 represents an int32 with accessor.
type NotNullInt32 struct {
	Int32Common
}

// Clone returns new instance of NotNullInt32 with preserved value & error
func (n NotNullInt32) Clone() Int32Accessor {
	nv := &NotNullInt32{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNInt32 returns NotNullInt32 under Int32Accessor from int32
func NNInt32(value int32) Int32Accessor {
	return &NotNullInt32{Int32Common{P: &value}}
}

// Int32Slice returns slice of int32 with filled values from slice of Int32Accessor
func Int32Slice(null []Int32Accessor, valid bool) []int32 {
	slice := make([]int32, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Int64Common represents an int64 with pointer and error.
type Int64Common struct {
	P     *int64
	Error error
}

// Set saves value into current struct
func (n *Int64Common) Set(value int64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Int64Common) V() int64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Int64Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Int64Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Int64Common) Value() (driver.Value, error) {
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Int64Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int64()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Int64Common) UnmarshalJSON(b []byte) error {
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
	v := FloatInt64(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Int64Common) MarshalJSON() ([]byte, error) {
	v := Of(n.V()).Float()
	if v.Err() != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Int64Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Int64Common) Err() error {
	return n.Error
}

// Int64Accessor accessor of int64 type.
type Int64Accessor interface {
	Common
	V() int64
	Set(value int64)
	Clone() Int64Accessor
}

// NullInt64 represents an int64 that may be null.
type NullInt64 struct {
	Int64Common
}

// Value implements the sql driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Int64Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Int64Common.MarshalJSON()
}

// Clone returns new instance of NullInt64 with preserved value & error
func (n NullInt64) Clone() Int64Accessor {
	nv := &NullInt64{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NInt64 returns NullInt64 under Int64Accessor from int64
func NInt64(value int64) Int64Accessor {
	return &NullInt64{Int64Common{P: &value}}
}

// NotNullInt64 represents an int64 with accessor.
type NotNullInt64 struct {
	Int64Common
}

// Clone returns new instance of NotNullInt64 with preserved value & error
func (n NotNullInt64) Clone() Int64Accessor {
	nv := &NotNullInt64{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNInt64 returns NotNullInt64 under from int64
func NNInt64(value int64) Int64Accessor {
	return &NotNullInt64{Int64Common{P: &value}}
}

// Int64Slice returns slice of int64 with filled values from slice of Int64Accessor
func Int64Slice(null []Int64Accessor, valid bool) []int64 {
	slice := make([]int64, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
