package typ

import (
	"fmt"
	"strconv"
)

// Convert interface value to string
func (t *Type) String() (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	nv = t.toString()
	transition(t, &nv)
	return
}

// Convert interface value to string and replace zero values to empty string if needed
func (t *Type) StringEmpty() (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	r, empty := t.toString(), t.Empty()
	if empty.V() {
		nv = NullString{Error: empty.Error}
		transition(t, &nv)
		return
	}
	transition(t, &r)
	return r
}

// Convert interface value to string with default value if it wasn't converted or it has zero value
func (t *Type) StringDefault(defaultValue string) (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	r, empty := t.toString(), t.Empty()
	if r.V() != "" && !empty.V() {
		transition(t, &r)
		return r
	}
	nv = NullString{P: &defaultValue, Error: r.Error}
	transition(t, &nv)
	return
}

// Convert interface value to string
func (t *Type) toString() (nv NullString) {
	if !t.rv.IsValid() {
		nv.Error = ErrUnexpectedValue
		return
	}
	if t.IsString(true) {
		v := t.rv.String()
		nv.P = &v
		return
	}
	if t.IsNumeric(true) {
		v := NumericToString(t.rv.Interface(), *t.opts.base, *t.opts.fmtByte, *t.opts.precision)
		nv.P = &v
		return
	}
	if t.IsBool(true) {
		if t.rv.Bool() {
			v := "true"
			nv.P = &v
			return
		} else {
			v := "false"
			nv.P = &v
			return
		}
	}
	v := fmt.Sprintf("%+v", t.rv.Interface())
	nv.P = &v
	return
}

// Convert interface value to string represents as int
func (t *Type) StringInt() (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	v := t.Int()
	return NewType(v.V(), v.Error).String()
}

// Convert interface value to string represent as bool
func (t *Type) StringBool() (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	v := t.Bool()
	return NewType(v.V(), v.Error).String()
}

// Convert interface value to string represent as float
func (t *Type) StringFloat(defaultValue ...float32) (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	v := t.Float()
	return NewType(v.V(), v.Error).String()
}

// Convert interface value to string represent as complex
func (t *Type) StringComplex(defaultValue ...complex64) (nv NullString) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	v := t.Complex()
	return NewType(v.V(), v.Error).String()
}

// Concatenate interface values to string
func Concat(values []interface{}, options ...Option) NullString {
	var opts opts
	if len(options) > 0 {
		for _, v := range options {
			if optErr := v(&opts); optErr != nil {
				return NullString{Error: optErr}
			}
		}
	}
	out := ""
	for i, v := range values {
		nv := Of(v, options...).String()
		if nv.Error != nil {
			return NullString{Error: nv.Error}
		}
		out += nv.V()
		if opts.delimiter != nil && len(values) > i+1 {
			out += *opts.delimiter
		}
	}
	return NullString{P: &out, Error: nil}
}

// Convert value from bool to string
func BoolString(from bool) (nv NullString) {
	if from {
		v := "true"
		nv.P = &v
		return
	}
	v := "false"
	nv.P = &v
	return
}

// Convert value from int64 to string
func IntString(from int64, options ...IntStringOption) (nv NullString) {
	opts := intStrOpts{base: 10}
	if len(options) > 0 {
		for _, v := range options {
			if optErr := v(&opts); optErr != nil {
				nv.Error = optErr
				break
			}
		}
	}
	if opts.defaultValue != nil {
		v := *opts.defaultValue
		nv.P = &v
		return
	}
	v := strconv.FormatInt(from, opts.base)
	nv.P = &v
	return
}

// Convert value from uint64 to string
func UintString(from uint64, options ...UintStringOption) (nv NullString) {
	opts := uintStrOpts{base: 10}
	if len(options) > 0 {
		for _, v := range options {
			if optErr := v(&opts); optErr != nil {
				nv.Error = optErr
				break
			}
		}
	}
	if opts.defaultValue != nil {
		v := *opts.defaultValue
		nv.P = &v
		return
	}
	v := strconv.FormatUint(from, opts.base)
	nv.P = &v
	return
}

// Convert value from float to string
func FloatString(from float64, options ...FloatStringOption) (nv NullString) {
	opts := floatStrOpts{fmtByte: byte('e'), precision: -1, bitSize: 64}
	if len(options) > 0 {
		for _, v := range options {
			if optErr := v(&opts); optErr != nil {
				nv.Error = optErr
				break
			}
		}
	}
	if opts.defaultValue != nil {
		v := *opts.defaultValue
		nv.P = &v
		return
	}
	v := strconv.FormatFloat(from, opts.fmtByte, opts.precision, opts.bitSize)
	nv.P = &v
	return
}

// Convert value from complex128 to string
func ComplexString(from complex128, options ...ComplexStringOption) (nv NullString) {
	opts := complexStrOpts{}
	if len(options) > 0 {
		for _, v := range options {
			if optErr := v(&opts); optErr != nil {
				nv.Error = optErr
				break
			}
		}
	}
	if opts.defaultValue != nil {
		v := *opts.defaultValue
		nv.P = &v
		return
	}
	v := fmt.Sprintf("%v", from)
	nv.P = &v
	return
}

func transition(t *Type, nv *NullString) {
	if t.opts.prefix != nil {
		v := *t.opts.prefix + nv.V()
		nv.P = &v
	}
	if t.opts.suffix != nil {
		v := nv.V() + *t.opts.suffix
		nv.P = &v
	}
}
