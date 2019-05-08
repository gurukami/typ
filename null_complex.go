package typ

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// ComplexCommon represents a complex128 with pointer and error.
type ComplexCommon struct {
	P     *complex128
	Error error
}

// Set saves value into current struct
func (n *ComplexCommon) Set(value complex128) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n ComplexCommon) V() complex128 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n ComplexCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n ComplexCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n ComplexCommon) Value() (driver.Value, error) {
	nv := ComplexFloat64(n.V())
	return nv.V(), nv.Err()
}

// Scan implements the sql Scanner interface.
func (n *ComplexCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Complex()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *ComplexCommon) UnmarshalJSON(b []byte) error {
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
		return n.Err()
	}
	vCmplx := StringComplex(v)
	if !vCmplx.Valid() {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(vCmplx.V())
	return n.Err()
}

// MarshalJSON implements the json Marshaler interface.
func (n ComplexCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%v", n.V()))
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n ComplexCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n ComplexCommon) Err() error {
	return n.Error
}

// ComplexAccessor accessor of complex128 type.
type ComplexAccessor interface {
	Common
	V() complex128
	Set(value complex128)
	Clone() ComplexAccessor
}

// NullComplex represents a complex128 that may be null.
type NullComplex struct {
	ComplexCommon
}

// Value implements the sql driver Valuer interface.
func (n NullComplex) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.ComplexCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullComplex) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.ComplexCommon.MarshalJSON()
}

// Clone returns new instance of NullComplex with preserved value & error
func (n NullComplex) Clone() ComplexAccessor {
	nv := &NullComplex{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NComplex returns NullComplex under ComplexAccessor from complex128
func NComplex(value complex128) ComplexAccessor {
	return &NullComplex{ComplexCommon{P: &value}}
}

// NotNullComplex represents a complex128 that may be null.
type NotNullComplex struct {
	ComplexCommon
}

// Clone returns new instance of NotNullComplex with preserved value & error
func (n NotNullComplex) Clone() ComplexAccessor {
	nv := &NotNullComplex{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNComplex returns NotNullComplex under ComplexAccessor from complex128
func NNComplex(value complex128) ComplexAccessor {
	return &NotNullComplex{ComplexCommon{P: &value}}
}

// ComplexSlice returns slice of complex128 with filled values from slice of ComplexAccessor
func ComplexSlice(null []ComplexAccessor, valid bool) []complex128 {
	slice := make([]complex128, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// Complex64Common represents a complex64 with pointer and error.
type Complex64Common struct {
	P     *complex64
	Error error
}

// Set saves value into current struct
func (n *Complex64Common) Set(value complex64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Complex64Common) V() complex64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Complex64Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Complex64Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Complex64Common) Value() (driver.Value, error) {
	nv := Complex64Float64(n.V())
	return nv.V(), nv.Err()
}

// Scan implements the sql Scanner interface.
func (n *Complex64Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Complex64()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return nil
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Complex64Common) UnmarshalJSON(b []byte) error {
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
		return n.Err()
	}
	vCmplx := StringComplex64(v)
	if !vCmplx.Valid() {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(vCmplx.V())
	return n.Err()
}

// MarshalJSON implements the json Marshaler interface.
func (n Complex64Common) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%v", n.V()))
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Complex64Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Complex64Common) Err() error {
	return n.Error
}

// Complex64Accessor accessor of complex64 type.
type Complex64Accessor interface {
	Common
	V() complex64
	Set(value complex64)
	Clone() Complex64Accessor
}

// NullComplex64 represents a complex64 that may be null.
type NullComplex64 struct {
	Complex64Common
}

// Value implements the sql driver Valuer interface.
func (n NullComplex64) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Complex64Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullComplex64) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Complex64Common.MarshalJSON()
}

// Clone returns new instance of NullComplex64 with preserved value & error
func (n NullComplex64) Clone() Complex64Accessor {
	nv := &NullComplex64{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NComplex64 returns NullComplex64 under Complex64Accessor from complex64
func NComplex64(value complex64) Complex64Accessor {
	return &NullComplex64{Complex64Common{P: &value}}
}

// NotNullComplex64 represents a complex64 that may be null.
type NotNullComplex64 struct {
	Complex64Common
}

// Clone returns new instance of NotNullComplex64 with preserved value & error
func (n NotNullComplex64) Clone() Complex64Accessor {
	nv := &NotNullComplex64{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNComplex64 returns NotNullComplex64 under Complex64Accessor from complex64
func NNComplex64(value complex64) Complex64Accessor {
	return &NotNullComplex64{Complex64Common{P: &value}}
}

// Complex64Slice returns slice of complex64 with filled values from slice of Complex64Accessor
func Complex64Slice(null []Complex64Accessor, valid bool) []complex64 {
	slice := make([]complex64, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
