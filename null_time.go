package typ

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// TimeCommon represents a time.Time that may be null.
type TimeCommon struct {
	P     *time.Time
	Error error
}

// Set saves value into current struct
func (n *TimeCommon) Set(value time.Time) {
	n.P = &value
}

// V returns value of underlying type if it was set, otherwise default value
func (n TimeCommon) V() time.Time {
	if n.P == nil {
		return time.Time{}
	}
	return *n.P
}

// Present determines whether a value has been set
func (n TimeCommon) Present() bool {
	return n.P != nil
}

// Valid determines whether a value has been valid
func (n TimeCommon) Valid() bool {
	return n.Err() == nil
}

// Value implements the sql driver Valuer interface.
func (n TimeCommon) Value() (driver.Value, error) {
	return n.V(), nil
}

// Scan implements the sql Scanner interface.
func (n *TimeCommon) Scan(value interface{}) error {
	n.Error = nil
	var tv time.Time
	tv, ok := value.(time.Time)
	if !ok {
		n.Error = ErrInvalidArgument
	}
	n.P = &tv
	return n.Err()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *TimeCommon) UnmarshalJSON(b []byte) error {
	v := n.V()
	n.Error = v.UnmarshalJSON(b)
	return n.Err()
}

// MarshalJSON implements the json Marshaler interface.
func (n TimeCommon) MarshalJSON() ([]byte, error) {
	return n.V().MarshalJSON()
}

// Typ returns new instance with himself value.
// If current value is invalid, nil *Type returned
func (n TimeCommon) Typ(options ...Option) *Type {
	if n.Err() != nil {
		return NewType(nil, n.Err())
	}
	return NewType(n.V(), n.Err(), options...)
}

// Err returns underlying error.
func (n TimeCommon) Err() error {
	return n.Error
}

// TimeAccessor accessor of time.Time type.
type TimeAccessor interface {
	Common
	V() time.Time
	Set(value time.Time)
	Clone() TimeAccessor
}

// NullTime represents a time.Time that may be null.
type NullTime struct {
	TimeCommon
}

// Value implements the sql driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if n.Err() != nil || !n.Present() {
		return nil, n.Err()
	}
	return n.TimeCommon.Value()
}

// MarshalJSON implements the json Marshaler interface.
func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.Err() != nil || !n.Present() {
		return json.Marshal(nil)
	}
	return n.TimeCommon.MarshalJSON()
}

// Clone returns new instance of NullTime with preserved value & error
func (n NullTime) Clone() TimeAccessor {
	nv := &NullTime{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NTime returns NullTime under TimeAccessor from time.Time
func NTime(value time.Time) TimeAccessor {
	return &NullTime{TimeCommon{P: &value}}
}

// NotNullTime represents a time.Time that may be null.
type NotNullTime struct {
	TimeCommon
}

// Clone returns new instance of NotNullTime with preserved value & error
func (n NotNullTime) Clone() TimeAccessor {
	nv := &NotNullTime{}
	if n.Present() {
		nv.Set(n.V())
	}
	nv.Error = n.Error
	return nv
}

// NNTime returns NullTime under TimeAccessor from time.Time
func NNTime(value time.Time) TimeAccessor {
	return &NotNullTime{TimeCommon{P: &value}}
}

// TimeSlice returns slice of time.Time with filled values from slice of TimeAccessor
func TimeSlice(null []TimeAccessor, valid bool) []time.Time {
	slice := make([]time.Time, 0, len(null))
	for _, v := range null {
		if valid && v.Err() != nil {
			continue
		}
		slice = append(slice, v.V())
	}
	return slice
}
