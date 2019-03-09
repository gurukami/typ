package typ

import (
	"reflect"
	"strings"
)

// Convert interface value to bool.
// Returns true for any non-zero values
func (t *Type) Bool() (nv NullBool) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	return t.toBool(false)
}

// Convert interface value to bool.
// Returns false for string 'false' in case-insensitive mode or string equals '0', for other types
// returns true only for positive values
func (t *Type) BoolHumanize() (nv NullBool) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	switch {
	case t.IsString(true):
		bf := strings.EqualFold("false", t.rv.String()) || t.rv.String() == "0"
		bt := strings.EqualFold("true", t.rv.String()) || t.rv.String() == "1"
		if !bf && !bt {
			nv.Error = ErrUnexpectedValue
		}
		nv.P = &bt
		return
	case t.IsBool(true):
		v := t.rv.Bool()
		nv.P = &v
		return
	default:
		return t.toBool(true)
	}
}

// Convert interface value to bool.
// Returns true only for positive values
func (t *Type) BoolPositive() (nv NullBool) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	return t.toBool(true)
}

// Convert interface value to bool.
// It check positive values if argument "positive" is true, otherwise always true for any non-zero values
func (t *Type) toBool(positive bool) (nv NullBool) {
	var v bool
	switch {
	case t.IsBool(true):
		v = t.rv.Bool()
		goto end
	case t.IsString(true):
		v = t.rv.Len() != 0
		goto end
	case t.IsInt(true):
		vInt := t.rv.Int()
		v = (vInt < 0 && !positive) || vInt > 0
		goto end
	case t.IsUint(true):
		vUint := t.rv.Uint()
		v = (vUint < 0 && !positive) || vUint > 0
		goto end
	case t.IsFloat(true):
		vFloat := t.rv.Float()
		v = (vFloat != 0 && !positive) || vFloat > 0
		goto end
	case t.IsComplex(true):
		value := t.rv.Complex()
		fr := real(value)
		v = (fr != 0 && !positive) || fr > 0
		goto end
	default:
		if positive {
			switch t.rv.Kind() {
			case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
				v = t.rv.Len() > 0
				goto end
			}
		}
		nv = t.Empty()
		*nv.P = !*nv.P
		return
	}
end:
	nv.P = &v
	return
}

// Convert value from string to bool.
// Returns false for string 'false' in case-insensitive mode or string equals '0'
func StringBoolHumanize(from string) (nv NullBool) {
	bf := strings.EqualFold("false", from) || from == "0"
	bt := strings.EqualFold("true", from) || from == "1"
	if !bf && !bt {
		nv.Error = ErrUnexpectedValue
		return
	}
	nv.P = &bt
	return
}
