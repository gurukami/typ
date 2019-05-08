package typ

import (
	"database/sql/driver"
	"encoding/json"
)

// BoolCommon represents a bool with pointer and error.
type BoolCommon struct {
	P     *bool
	Error error
}

// Set saves value into current struct
func (n *BoolCommon) Set(value bool) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n BoolCommon) V() bool {
	if n.P == nil {
		return false
	}
	return *n.P
}

// Present determines whether a value has been set
func (n BoolCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n BoolCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n BoolCommon) Value() (driver.Value, error) {
	return n.V(), n.Err()
}

// Scan implements the sql Scanner interface.
func (n *BoolCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	v := Of(value).Bool().V()
	n.P = &v
	return nil
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *BoolCommon) UnmarshalJSON(b []byte) error {
	n.P, n.Error = nil, nil
	var uv interface{}
	if err := json.Unmarshal(b, &uv); err != nil {
		n.Error = err
		return n.Err()
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
func (n BoolCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n BoolCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n BoolCommon) Err() error {
	return n.Error
}

// BoolAccessor
type BoolAccessor interface {
	NullCommon
	V() bool
	Set(value bool)
	Clone() BoolAccessor
}

// NullBool represents a bool that may be null.
type NullBool struct {
	BoolCommon
}

// Value implements the sql driver Valuer interface.
func (n NullBool) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.BoolCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullBool) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.BoolCommon.MarshalJSON()
}

// Clone returns new instance of NullBool with preserved value & error
func (n NullBool) Clone() BoolAccessor {
	nv := &NullBool{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NBool returns NullBool under BoolAccessor from bool
func NBool(value bool) BoolAccessor {
	return &NullBool{BoolCommon{P: &value}}
}

// NotNullBool represents a bool with accessor.
type NotNullBool struct {
	BoolCommon
}

// Clone returns new instance of NullBool with preserved value & error
func (n NotNullBool) Clone() BoolAccessor {
	nv := &NotNullBool{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNBool returns NotNullBool under BoolAccessor from bool
func NNBool(value bool) BoolAccessor {
	return &NotNullBool{BoolCommon{P: &value}}
}

// BoolSlice returns slice of bool with filled values from slice of BoolAccessor
func BoolSlice(null []BoolAccessor, valid bool) []bool {
	slice := make([]bool, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
