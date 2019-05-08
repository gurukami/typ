package typ

import (
	"database/sql/driver"
	"encoding/json"
)

// Float32Common represents a float32 that may be null.
type Float32Common struct {
	P     *float32
	Error error
}

// Set saves value into current struct
func (n *Float32Common) Set(value float32) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n Float32Common) V() float32 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n Float32Common) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n Float32Common) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n Float32Common) Value() (driver.Value, error) {
	return float64(n.V()), nil
}

// Scan implements the sql Scanner interface.
func (n *Float32Common) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Float32()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *Float32Common) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return n.Err()
	}
	if uv == nil {
		return nil
	}
	vFloat, ok := uv.(float64)
	if !ok {
		n.Error = ErrConvert
		return n.Err()
	}
	v := Float32(vFloat)
	if v.Err() != nil {
		n.Error = ErrConvert
		return n.Err()
	}
	n.Set(v.V())
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n Float32Common) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n Float32Common) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n Float32Common) Err() error {
	return n.Error
}

// Float32Accessor accessor of float32 type.
type Float32Accessor interface {
	Common
	V() float32
	Set(value float32)
	Clone() Float32Accessor
}

// NullFloat32 represents a float32 that may be null.
type NullFloat32 struct {
	Float32Common
}

// Value implements the sql driver Valuer interface.
func (n NullFloat32) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.Float32Common.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullFloat32) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.Float32Common.MarshalJSON()
}

// Clone returns new instance of NullFloat32 with preserved value & error
func (n NullFloat32) Clone() Float32Accessor {
	nv := &NullFloat32{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NFloat32 returns NullFloat32 under from float32
func NFloat32(value float32) Float32Accessor {
	return &NullFloat32{Float32Common{P: &value}}
}

// NotNullFloat32 represents a float32 that may be null.
type NotNullFloat32 struct {
	Float32Common
}

// Clone returns new instance of NotNullFloat32 with preserved value & error
func (n NotNullFloat32) Clone() Float32Accessor {
	nv := &NotNullFloat32{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNFloat32 returns NullFloat32 under Float32Accessor from float32
func NNFloat32(value float32) Float32Accessor {
	return &NotNullFloat32{Float32Common{P: &value}}
}

// Float32Slice returns slice of float32 with filled values from slice of Float32Accessor
func Float32Slice(null []Float32Accessor, valid bool) []float32 {
	slice := make([]float32, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}

// FloatCommon represents a float64 that may be null.
type FloatCommon struct {
	P     *float64
	Error error
}

// Set saves value into current struct
func (n *FloatCommon) Set(value float64) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n FloatCommon) V() float64 {
	if n.P == nil {
		return 0
	}
	return *n.P
}

// Present determines whether a value has been set
func (n FloatCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n FloatCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n FloatCommon) Value() (driver.Value, error) {
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *FloatCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Float()
	if v.Err() != nil {
		n.Error = v.Err()
		return v.Err()
	}
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *FloatCommon) UnmarshalJSON(b []byte) error {
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
		return n.Err()
	}
	n.P = &v
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n FloatCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n FloatCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n FloatCommon) Err() error {
	return n.Error
}

// FloatAccessor accessor of float64 type.
type FloatAccessor interface {
	Common
	V() float64
	Set(value float64)
	Clone() FloatAccessor
}

// NullFloat represents a float64 that may be null.
type NullFloat struct {
	FloatCommon
}

// Value implements the sql driver Valuer interface.
func (n NullFloat) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.FloatCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullFloat) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.FloatCommon.MarshalJSON()
}

// Clone returns new instance of NullFloat with preserved value & error
func (n NullFloat) Clone() FloatAccessor {
	nv := &NullFloat{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NFloat returns NullFloat under FloatAccessor from float64
func NFloat(value float64) FloatAccessor {
	return &NullFloat{FloatCommon{P: &value}}
}

// NotNullFloat represents a float64 that may be null.
type NotNullFloat struct {
	FloatCommon
}

// Clone returns new instance of NotNullFloat with preserved value & error
func (n NotNullFloat) Clone() FloatAccessor {
	nv := &NotNullFloat{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNFloat returns NullFloat under FloatAccessor from float64
func NNFloat(value float64) FloatAccessor {
	return &NotNullFloat{FloatCommon{P: &value}}
}

// FloatSlice returns slice of float64 with filled values from slice of FloatAccessor
func FloatSlice(null []FloatAccessor, valid bool) []float64 {
	slice := make([]float64, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
