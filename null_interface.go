package typ

import (
	"database/sql/driver"
	"encoding/json"
)

// InterfaceCommon represents an interface{} that may be null.
type InterfaceCommon struct {
	P     interface{}
	Error error
}

// Set saves value into current struct
func (n *InterfaceCommon) Set(value interface{}) {
	n.P = value
}

// V returns value of underlying type if it was set, otherwise default value
func (n InterfaceCommon) V() interface{} {
	return n.P
}

// Present determines whether a value has been set
func (n InterfaceCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n InterfaceCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n InterfaceCommon) Value() (driver.Value, error) {
	if n.Err() == nil && !driver.IsValue(n.V()) {
		return nil, ErrConvert
	}
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *InterfaceCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if !driver.IsValue(value) {
		n.Error = ErrInvalidArgument
		return n.Err()
	}
	n.P = value
	return nil
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *InterfaceCommon) UnmarshalJSON(b []byte) error {
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
func (n InterfaceCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n InterfaceCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n InterfaceCommon) Err() error {
	return n.Error
}

// InterfaceAccessor
type InterfaceAccessor interface {
	NullCommon
	V() interface{}
	Set(value interface{})
	Clone() InterfaceAccessor
}

// NullInterface represents an interface{} that may be null.
type NullInterface struct {
	InterfaceCommon
}

// Value implements the sql driver Valuer interface.
func (n NullInterface) Value() (driver.Value, error) {
	if n.Err() == nil && !driver.IsValue(n.V()) {
		return nil, ErrConvert
	}
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.InterfaceCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullInterface) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.InterfaceCommon.MarshalJSON()
}

// Clone returns new instance of NullInterface with preserved value & error
func (n NullInterface) Clone() InterfaceAccessor {
	nv := &NullInterface{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NInterface returns NullInterface under InterfaceAccessor from interface{}
func NInterface(value interface{}) InterfaceAccessor {
	return &NullInterface{InterfaceCommon{P: value}}
}

// NotNullInterface represents an interface{} that may be null.
type NotNullInterface struct {
	InterfaceCommon
}

// Clone returns new instance of NotNullInterface with preserved value & error
func (n NotNullInterface) Clone() InterfaceAccessor {
	nv := &NotNullInterface{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNInterface returns NotNullInterface under InterfaceAccessor from interface{}
func NNInterface(value interface{}) InterfaceAccessor {
	return &NotNullInterface{InterfaceCommon{P: value}}
}

// InterfaceSlice returns slice of interface{} with filled values from slice of InterfaceAccessor
func InterfaceSlice(null []InterfaceAccessor, valid bool) []interface{} {
	slice := make([]interface{}, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
