package typ

import (
	"database/sql/driver"
	"encoding/json"
)

// StringCommon represents a string that may be null.
type StringCommon struct {
	P     *string
	Error error
}

// Set saves value into current struct
func (n *StringCommon) Set(value string) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n StringCommon) V() string {
	if n.P == nil {
		return ""
	}
	return *n.P
}

// Present determines whether a value has been set
func (n StringCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n StringCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n StringCommon) Value() (driver.Value, error) {
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *StringCommon) Scan(value interface{}) error {
	n.P, n.Error = nil, nil
	if value == nil {
		return nil
	}
	if v, ok := value.([]byte); ok {
		value = string(v)
	}
	v := Of(value).String()
	n.Set(v.V())
	n.Error = v.Err()
	return v.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *StringCommon) UnmarshalJSON(b []byte) error {
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
	n.P = &v
	return nil
}

// MarshalJSON implements the json Marshaler interface.
func (n StringCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.V())
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n StringCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n StringCommon) Err() error {
	return n.Error
}

// StringAccessor
type StringAccessor interface {
	NullCommon
	V() string
	Set(value string)
	Clone() StringAccessor
}

// NullString represents a string that may be null.
type NullString struct {
	StringCommon
}

// Value implements the sql driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.StringCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullString) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.StringCommon.MarshalJSON()
}

// Clone returns new instance of NullString with preserved value & error
func (n NullString) Clone() StringAccessor {
	nv := &NullString{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NString returns NullString under StringAccessor from string
func NString(value string) StringAccessor {
	return &NullString{StringCommon{P: &value}}
}

// NotNullString represents a string that may be null.
type NotNullString struct {
	StringCommon
}

// Clone returns new instance of NotNullString with preserved value & error
func (n NotNullString) Clone() StringAccessor {
	nv := &NotNullString{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNString returns NullString under StringAccessor from string
func NNString(value string) StringAccessor {
	return &NotNullString{StringCommon{P: &value}}
}

// StringSlice returns slice of string with filled values from slice of StringAccessor
func StringSlice(null []StringAccessor, valid bool) []string {
	slice := make([]string, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
