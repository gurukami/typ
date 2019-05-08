package typ

import (
	"fmt"
	"strconv"
)

// String convert interface value to string
func (t *Type) String() StringAccessor {
	var nv StringAccessor
	nv = &NullString{StringCommon{Error: t.err}}
	if t.err != nil {
		return nv
	}
	nv = t.toString()
	transition(t, nv)
	return nv
}

// StringEmpty convert interface value to string and replace zero values to empty string if needed
func (t *Type) StringEmpty() StringAccessor {
	nv := &NullString{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	r, empty := t.toString(), t.Empty()
	if empty.V() {
		nv = &NullString{StringCommon{Error: empty.Err()}}
		transition(t, nv)
		return nv
	}
	transition(t, r)
	return r
}

// StringDefault convert interface value to string with default value if it wasn't converted or it has zero value
func (t *Type) StringDefault(defaultValue string) StringAccessor {
	var nv StringAccessor
	nv = &NullString{StringCommon{Error: t.err}}
	if t.err != nil {
		return nv
	}
	r, empty := t.toString(), t.Empty()
	if r.V() != "" && !empty.V() {
		transition(t, r)
		return r
	}
	nv = &NullString{StringCommon{P: &defaultValue, Error: r.Err()}}
	transition(t, nv)
	return nv
}

// toString convert interface value to string
func (t *Type) toString() StringAccessor {
	nv := &NullString{}
	if !t.rv.IsValid() {
		nv.Error = ErrUnexpectedValue
		return nv
	}
	if t.IsString(true) {
		v := t.rv.String()
		nv.P = &v
		return nv
	}
	if t.IsNumeric(true) {
		v := NumericToString(t.rv.Interface(), *t.opts.base, *t.opts.fmtByte, *t.opts.precision)
		nv.P = &v
		return nv
	}
	if t.IsBool(true) {
		if t.rv.Bool() {
			v := "true"
			nv.P = &v
			return nv
		}
		v := "false"
		nv.P = &v
		return nv
	}
	v := fmt.Sprintf("%+v", t.rv.Interface())
	nv.P = &v
	return nv
}

// StringInt convert interface value to string represents as int
func (t *Type) StringInt() StringAccessor {
	var nv StringAccessor
	nv = &NullString{StringCommon{Error: t.err}}
	if t.err != nil {
		return nv
	}
	v := t.Int()
	return NewType(v.V(), v.Err()).String()
}

// StringBool convert interface value to string represent as bool
func (t *Type) StringBool() StringAccessor {
	var nv StringAccessor
	nv = &NullString{StringCommon{Error: t.err}}
	if t.err != nil {
		return nv
	}
	v := t.Bool()
	return NewType(v.V(), v.Err()).String()
}

// StringFloat convert interface value to string represent as float
func (t *Type) StringFloat(defaultValue ...float32) StringAccessor {
	var nv StringAccessor
	nv = &NullString{StringCommon{Error: t.err}}
	if t.err != nil {
		return nv
	}
	v := t.Float()
	return NewType(v.V(), v.Err()).String()
}

// StringComplex convert interface value to string represent as complex
func (t *Type) StringComplex(defaultValue ...complex64) StringAccessor {
	var nv StringAccessor
	nv = &NullString{StringCommon{Error: t.err}}
	if t.err != nil {
		return nv
	}
	v := t.Complex()
	return NewType(v.V(), v.Err()).String()
}

// Concat returns concatenated interface values to string
func Concat(values []interface{}, options ...Option) StringAccessor {
	var opts opts
	if len(options) > 0 {
		for _, v := range options {
			if optErr := v(&opts); optErr != nil {
				return &NullString{StringCommon{Error: optErr}}
			}
		}
	}
	out := ""
	for i, v := range values {
		nv := Of(v, options...).String()
		if nv.Err() != nil {
			return &NullString{StringCommon{Error: nv.Err()}}
		}
		out += nv.V()
		if opts.delimiter != nil && len(values) > i+1 {
			out += *opts.delimiter
		}
	}
	return &NullString{StringCommon{P: &out}}
}

// BoolString convert value from bool to string
func BoolString(from bool) StringAccessor {
	nv := &NullString{}
	if from {
		v := "true"
		nv.P = &v
		return nv
	}
	v := "false"
	nv.P = &v
	return nv
}

// IntString convert value from int64 to string
func IntString(from int64, options ...IntStringOption) StringAccessor {
	nv := &NullString{}
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
		return nv
	}
	v := strconv.FormatInt(from, opts.base)
	nv.P = &v
	return nv
}

// UintString convert value from uint64 to string
func UintString(from uint64, options ...UintStringOption) StringAccessor {
	nv := &NullString{}
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
		return nv
	}
	v := strconv.FormatUint(from, opts.base)
	nv.P = &v
	return nv
}

// FloatString convert value from float to string
func FloatString(from float64, options ...FloatStringOption) StringAccessor {
	nv := &NullString{}
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
		return nv
	}
	v := strconv.FormatFloat(from, opts.fmtByte, opts.precision, opts.bitSize)
	nv.P = &v
	return nv
}

// ComplexString convert value from complex128 to string
func ComplexString(from complex128, options ...ComplexStringOption) StringAccessor {
	nv := &NullString{}
	opts := complexStrOpts{}
	if len(options) > 0 {
		for _, v := range options {
			_ = v(&opts)
		}
	}
	if opts.defaultValue != nil {
		v := *opts.defaultValue
		nv.P = &v
		return nv
	}
	v := fmt.Sprintf("%v", from)
	nv.P = &v
	return nv
}

func transition(t *Type, nv StringAccessor) {
	if t.opts.prefix != nil {
		v := *t.opts.prefix + nv.V()
		nv.Set(v)
	}
	if t.opts.suffix != nil {
		v := nv.V() + *t.opts.suffix
		nv.Set(v)
	}
}
