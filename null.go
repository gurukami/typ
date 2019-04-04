package typ

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// NullBool represents an bool that may be null.
type NullBool struct {
	P     *bool
	Error error
}

// Set saves value into current struct
func (n *NullBool) Set(value bool) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullBool) V() bool {
	if n.P == nil {
		return false
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullBool) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullBool) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullBool) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *NullBool) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Bool().V()
	n.P = &v
	return nil
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullBool) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return n.Error
	}
	if uv == nil {
		return nil
	}
	v, ok := uv.(bool)
	if !ok {
		n.Error = ErrConvert
		return ErrConvert
	}
	n.P = &v
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullBool) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullBool) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullBoolSlice returns slice of bool with filled values from slice of NullBool
func NullBoolSlice(null []NullBool, valid bool) []bool {
	slice := make([]bool, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NBool returns NullBool from bool
func NBool(value bool) NullBool {
	return NullBool{P: &value}
}

// NullComplex64 represents an complex64 that may be null.
type NullComplex64 struct {
	P     *complex64
	Error error
}

// Set saves value into current struct
func (n *NullComplex64) Set(value complex64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullComplex64) V() complex64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullComplex64) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullComplex64) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullComplex64) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	nv := Complex64Float64(n.V())
	return nv.V(), nv.Error
}

// Scan implements the sql Scanner interface.
func (n *NullComplex64) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Complex64()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return nil
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullComplex64) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	v, ok := uv.(string)
	if !ok {
		n.Error = ErrConvert
		return n.Error
	}
	vCmplx := StringComplex64(v)
	if !vCmplx.Valid() {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = vCmplx.P
	return n.Error
}

// MarshalJSON implements the json Marshaler interface.
func (n NullComplex64) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(fmt.Sprintf("%v", n.V()))
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullComplex64) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullComplex64Slice returns slice of complex64 with filled values from slice of NullComplex64
func NullComplex64Slice(null []NullComplex64, valid bool) []complex64 {
	slice := make([]complex64, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NComplex64 returns NullComplex64 from complex64
func NComplex64(value complex64) NullComplex64 {
	return NullComplex64{P: &value}
}

// NullComplex represents a complex128 that may be null.
type NullComplex struct {
	P     *complex128
	Error error
}

// Set saves value into current struct
func (n *NullComplex) Set(value complex128) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullComplex) V() complex128 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullComplex) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullComplex) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullComplex) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	nv := ComplexFloat64(n.V())
	return nv.V(), nv.Error
}

// Scan implements the sql Scanner interface.
func (n *NullComplex) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Complex()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullComplex) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	v, ok := uv.(string)
	if !ok {
		n.Error = ErrConvert
		return n.Error
	}
	vCmplx := StringComplex(v)
	if !vCmplx.Valid() {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = vCmplx.P
	return n.Error
}

// MarshalJSON implements the json Marshaler interface.
func (n NullComplex) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(fmt.Sprintf("%v", n.V()))
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullComplex) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullComplexSlice returns slice of complex128 with filled values from slice of NullComplex
func NullComplexSlice(null []NullComplex, valid bool) []complex128 {
	slice := make([]complex128, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NComplex returns NullComplex from complex128
func NComplex(value complex128) NullComplex {
	return NullComplex{P: &value}
}

// NullInt represents an int that may be null.
type NullInt struct {
	P     *int
	Error error
}

// Set saves value into current struct
func (n *NullInt) Set(value int) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullInt) V() int {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullInt) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullInt) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullInt) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullInt) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullInt) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatInt(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	if v.Error != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullInt) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullIntSlice returns slice of int with filled values from slice of NullInt
func NullIntSlice(null []NullInt, valid bool) []int {
	slice := make([]int, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NInt returns NullInt from int
func NInt(value int) NullInt {
	return NullInt{P: &value}
}

// NullInt8 represents an int8 that may be null.
type NullInt8 struct {
	P     *int8
	Error error
}

// Set saves value into current struct
func (n *NullInt8) Set(value int8) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullInt8) V() int8 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullInt8) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullInt8) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullInt8) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int8()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullInt8) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatInt8(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt8) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullInt8) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullInt8Slice returns slice of int8 with filled values from slice of NullInt8
func NullInt8Slice(null []NullInt8, valid bool) []int8 {
	slice := make([]int8, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NInt8 returns NullInt8 from int8
func NInt8(value int8) NullInt8 {
	return NullInt8{P: &value}
}

// NullInt16 represents an int16 that may be null.
type NullInt16 struct {
	P     *int16
	Error error
}

// Set saves value into current struct
func (n *NullInt16) Set(value int16) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullInt16) V() int16 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullInt16) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullInt16) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullInt16) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int16()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullInt16) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatInt16(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt16) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullInt16) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullInt16Slice returns slice of int16 with filled values from slice of NullInt16
func NullInt16Slice(null []NullInt16, valid bool) []int16 {
	slice := make([]int16, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NInt16 returns NullInt16 from int16
func NInt16(value int16) NullInt16 {
	return NullInt16{P: &value}
}

// NullInt32 represents an int32 that may be null.
type NullInt32 struct {
	P     *int32
	Error error
}

// Set saves value into current struct
func (n *NullInt32) Set(value int32) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullInt32) V() int32 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullInt32) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullInt32) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullInt32) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int32()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullInt32) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatInt32(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullInt32) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullInt32Slice returns slice of int32 with filled values from slice of NullInt32
func NullInt32Slice(null []NullInt32, valid bool) []int32 {
	slice := make([]int32, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NInt32 returns NullInt32 from int32
func NInt32(value int32) NullInt32 {
	return NullInt32{P: &value}
}

// NullInt64 represents an int64 that may be null.
type NullInt64 struct {
	P     *int64
	Error error
}

// Set saves value into current struct
func (n *NullInt64) Set(value int64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullInt64) V() int64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullInt64) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullInt64) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullInt64) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Int64()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullInt64) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatInt64(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	if v.Error != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullInt64) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullInt64Slice returns slice of int64 with filled values from slice of NullInt64
func NullInt64Slice(null []NullInt64, valid bool) []int64 {
	slice := make([]int64, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NInt64 returns NullInt64 from int64
func NInt64(value int64) NullInt64 {
	return NullInt64{P: &value}
}

// NullUint represents a uint that may be null.
type NullUint struct {
	P     *uint
	Error error
}

// Set saves value into current struct
func (n *NullUint) Set(value uint) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullUint) V() uint {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullUint) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullUint) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullUint) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	nv := UintInt64(uint64(n.V()))
	return nv.V(), nv.Error
}

// Scan implements the sql Scanner interface.
func (n *NullUint) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullUint) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatUint(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	if v.Error != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullUint) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullUintSlice returns slice of uint with filled values from slice of NullUint
func NullUintSlice(null []NullUint, valid bool) []uint {
	slice := make([]uint, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NUint returns NullUint from uint
func NUint(value uint) NullUint {
	return NullUint{P: &value}
}

// NullUint8 represents a uint8 that may be null.
type NullUint8 struct {
	P     *uint8
	Error error
}

// Set saves value into current struct
func (n *NullUint8) Set(value uint8) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullUint8) V() uint8 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullUint8) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullUint8) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullUint8) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullUint8) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint8()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullUint8) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatUint8(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint8) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullUint8) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullUint8Slice returns slice of uint8 with filled values from slice of NullUint8
func NullUint8Slice(null []NullUint8, valid bool) []uint8 {
	slice := make([]uint8, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NUint8 returns NullUint8 from uint8
func NUint8(value uint8) NullUint8 {
	return NullUint8{P: &value}
}

// NullUint16 represents a uint16 that may be null.
type NullUint16 struct {
	P     *uint16
	Error error
}

// Set saves value into current struct
func (n *NullUint16) Set(value uint16) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullUint16) V() uint16 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullUint16) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullUint16) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullUint16) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullUint16) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint16()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullUint16) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatUint16(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint16) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullUint16) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullUint16Slice returns slice of uint16 with filled values from slice of NullUint16
func NullUint16Slice(null []NullUint16, valid bool) []uint16 {
	slice := make([]uint16, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NUint16 returns NullUint16 from uint16
func NUint16(value uint16) NullUint16 {
	return NullUint16{P: &value}
}

// NullUint32 represents a uint32 that may be null.
type NullUint32 struct {
	P     *uint32
	Error error
}

// Set saves value into current struct
func (n *NullUint32) Set(value uint32) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullUint32) V() uint32 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullUint32) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullUint32) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullUint32) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return int64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullUint32) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint32()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullUint32) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatUint32(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint32) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullUint32) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullUint32Slice returns slice of uint32 with filled values from slice of NullUint32
func NullUint32Slice(null []NullUint32, valid bool) []uint32 {
	slice := make([]uint32, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NUint32 returns NullUint32 from uint32
func NUint32(value uint32) NullUint32 {
	return NullUint32{P: &value}
}

// NullUint64 represents a uint64 that may be null.
type NullUint64 struct {
	P     *uint64
	Error error
}

// Set saves value into current struct
func (n *NullUint64) Set(value uint64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullUint64) V() uint64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullUint64) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullUint64) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	nv := UintInt64(n.V())
	return nv.V(), nv.Error
}

// Scan implements the sql Scanner interface.
func (n *NullUint64) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Uint64()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullUint64) UnmarshalJSON(b []byte) error {
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
		return n.Error
	}
	v := FloatUint64(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullUint64) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	v := Of(n.V()).Float()
	if v.Error != nil {
		return nil, ErrConvert
	}
	return json.Marshal(v.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullUint64) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullUint64Slice returns slice of uint64 with filled values from slice of NullUint64
func NullUint64Slice(null []NullUint64, valid bool) []uint64 {
	slice := make([]uint64, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NUint64 returns NullUint64 from uint64
func NUint64(value uint64) NullUint64 {
	return NullUint64{P: &value}
}

// NullFloat32 represents a float32 that may be null.
type NullFloat32 struct {
	P     *float32
	Error error
}

// Set saves value into current struct
func (n *NullFloat32) Set(value float32) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullFloat32) V() float32 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullFloat32) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullFloat32) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullFloat32) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return float64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *NullFloat32) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Float32()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullFloat32) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return n.Error
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Error
	}
	v := Float32(vFloat)
	if v.Error != nil {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = v.P
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullFloat32) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullFloat32) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullFloat32Slice returns slice of float32 with filled values from slice of NullFloat32
func NullFloat32Slice(null []NullFloat32, valid bool) []float32 {
	slice := make([]float32, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NFloat32 returns NullFloat32 from float32
func NFloat32(value float32) NullFloat32 {
	return NullFloat32{P: &value}
}

// NullFloat represents a float64 that may be null.
type NullFloat struct {
	P     *float64
	Error error
}

// Set saves value into current struct
func (n *NullFloat) Set(value float64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullFloat) V() float64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullFloat) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullFloat) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullFloat) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *NullFloat) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Float()
	if v.Error != nil {
		n.Error = v.Error
		return v.Error
	}
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullFloat) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	v, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = &v
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullFloat) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullFloat) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullFloatSlice returns slice of float64 with filled values from slice of NullFloat
func NullFloatSlice(null []NullFloat, valid bool) []float64 {
	slice := make([]float64, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NFloat returns NullFloat from float64
func NFloat(value float64) NullFloat {
	return NullFloat{P: &value}
}

// NullString represents a string that may be null.
type NullString struct {
	P     *string
	Error error
}

// Set saves value into current struct
func (n *NullString) Set(value string) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullString) V() string {
	if n.P == nil {
		return ""
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullString) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullString) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *NullString) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	if v, ok := value.([]byte); ok {
		value = string(v)
	}
	v := Of(value).String()
	n.P, n.Error = v.P, v.Error
	return v.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullString) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return err
	}
	if uv == nil {
		return nil
	}
	v, ok := uv.(string)
	if !ok {
		n.Error = ErrConvert
		return n.Error
	}
	n.P = &v
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullString) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullString) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullStringSlice returns slice of string with filled values from slice of NullString
func NullStringSlice(null []NullString, valid bool) []string {
	slice := make([]string, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NString returns NullString from string
func NString(value string) NullString {
	return NullString{P: &value}
}

// NullInterface represents an interface{} that may be null.
type NullInterface struct {
	P     interface{}
	Error error
}

// Set saves value into current struct
func (n *NullInterface) Set(value interface{}) {
	n.P = value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullInterface) V() interface{} {
	return n.P
}

// Present determines whether a value has been set
func (n NullInterface) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullInterface) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullInterface) Value() (driver.Value, error) {
	if n.Error == nil && !driver.IsValue(n.V()) {
		return nil, ErrConvert
	}
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *NullInterface) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if !driver.IsValue(value) {
		n.Error = ErrInvalidArgument
		return n.Error
	}
	n.P = value
	return nil
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullInterface) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var value interface{}
	if err := json.Unmarshal(b, &value); err != nil {
		n.Error = err
		return err
	}
	if value == nil {
		return nil
	}
	n.P = value
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInterface) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n NullInterface) Typ(options ...Option) *Type {
	if n.Error != nil {
		return NewType(nil, n.Error)
	}
	return NewType(n.V(), n.Error, options...)
}

// NullInterfaceSlice returns slice of interface{} with filled values from slice of NullInterface
func NullInterfaceSlice(null []NullInterface, valid bool) []interface{} {
	slice := make([]interface{}, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NInterface returns NullInterface from interface{}
func NInterface(value interface{}) NullInterface {
	return NullInterface{P: value}
}

// NullTime represents a time.Time that may be null.
type NullTime struct {
	P     *time.Time
	Error error
}

// Set saves value into current struct
func (n *NullTime) Set(value time.Time) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n NullTime) V() time.Time {
	if n.P == nil {
		return time.Time{}
	}
	return *n.P
}

// Present determines whether a value has been set
func (n NullTime) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n NullTime) Valid() bool {
	return n.Error == nil
}

// Value implements the sql driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if n.Error != nil || !n.Present() {
		return nil, nil
	}
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *NullTime) Scan(value interface{}) error {
	n.Error = nil
	var tv time.Time
	tv, ok := value.(time.Time)
	if !ok {
		n.Error = ErrInvalidArgument
	}
	n.P = &tv
	return n.Error
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullTime) UnmarshalJSON(b []byte) error {
	v := n.V()
	n.Error = v.UnmarshalJSON(b)
	return n.Error
}

// MarshalJSON implements the json Marshaler interface.
func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.Error != nil {
		return []byte("null"), nil
	}
	return n.V().MarshalJSON()
}

// NullTimeSlice returns slice of time.Time with filled values from slice of NullTime
func NullTimeSlice(null []NullTime, valid bool) []time.Time {
	slice := make([]time.Time, 0, len(null))
	for _, v := range null {
		if valid && v.Error != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// NTime returns NullTime from time.Time
func NTime(value time.Time) NullTime {
	return NullTime{P: &value}
}
