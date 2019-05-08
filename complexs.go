package typ

import (
	"reflect"
	"regexp"
	"strconv"
)

const regexpComplexString = `(?i)\(?([+-]?\d+(?:\.\d+(?:e\+\d+)?)?)([+-]\d+(?:\.\d+(?:e\+\d+)?)?)i\)?`

var regexpComplex = regexp.MustCompile(regexpComplexString)

// Complex64 convert interface value to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Complex64(defaultValue ...complex64) Complex64Accessor {
	nv := &NullComplex64{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	vt := t.toComplex(reflect.Complex64)
	v := complex64(vt.V())
	nv.Error, nv.P = vt.Err(), &v
	defaultComplex64(nv, defaultValue...)
	return nv
}

// Complex convert interface value to complex128
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Complex(defaultValue ...complex128) ComplexAccessor {
	nv := &NullComplex{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	nvTo := t.toComplex(reflect.Complex128)
	nv.Error = nvTo.Err()
	nv.Set(nvTo.V())
	defaultComplex(nv, defaultValue...)
	return nv
}

// Convert interface value to any complex type.
// Returns error if type can't safely converted
func (t *Type) toComplex(typeTo reflect.Kind) ComplexAccessor {
	nv := &NullComplex{}
	if !t.rv.IsValid() || !isComplex(typeTo) {
		nv.Error = ErrConvert
		return nv
	}
	switch {
	case t.IsString(true):
		matches := regexpComplex.FindStringSubmatch(t.rv.String())
		if len(matches) == 3 {
			fr, re := strconv.ParseFloat(matches[1], bitSizeMap[typeTo])
			fi, ie := strconv.ParseFloat(matches[2], bitSizeMap[typeTo])
			v := complex(fr, fi)
			nv.P = &v
			if re != nil || ie != nil {
				nv.Error = ErrUnexpectedValue
				return nv
			}
			return nv
		}
		nv.Error = ErrConvert
		return nv
	case t.IsComplex(true):
		v := t.rv.Complex()
		nv.P = &v
		if !isSafeComplex(t.rv.Complex(), bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsBool(true):
		var v complex128
		if t.rv.Bool() {
			v = complex(1, 0)
			nv.P = &v
		} else {
			nv.P = &v
		}
		return nv
	}
	floatValue := t.toFloat(complexFloatMap[typeTo])
	v := complex(floatValue.V(), 0)
	nv.P, nv.Error = &v, floatValue.Err()
	return nv
}

// IntComplex64 convert value from int64 to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntComplex64(from int64, defaultValue ...complex64) Complex64Accessor {
	nv := &NullComplex64{}
	if safe := isSafeIntToFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultComplex64(nv, defaultValue...) {
			return nv
		}
	}
	v := complex(float32(from), 0)
	nv.P = &v
	return nv
}

// IntComplex convert value from int64 to complex128.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntComplex(from int64, defaultValue ...complex128) ComplexAccessor {
	nv := &NullComplex{}
	if safe := isSafeIntToFloat(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultComplex(nv, defaultValue...) {
			return nv
		}
	}
	v := complex(float64(from), 0)
	nv.P = &v
	return nv
}

// UintComplex64 convert value from uint64 to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintComplex64(from uint64, defaultValue ...complex64) Complex64Accessor {
	nv := &NullComplex64{}
	if safe := isSafeUintToFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultComplex64(nv, defaultValue...) {
			return nv
		}
	}
	v := complex(float32(from), 0)
	nv.P = &v
	return nv
}

// UintComplex convert value from uint64 to complex128.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintComplex(from uint64, defaultValue ...complex128) ComplexAccessor {
	nv := &NullComplex{}
	if safe := isSafeUintToFloat(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultComplex(nv, defaultValue...) {
			return nv
		}
	}
	v := complex(float64(from), 0)
	nv.P = &v
	return nv
}

// Float32Complex64 convert value from float32 to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Complex64(from float32, defaultValue ...complex64) Complex64Accessor {
	return FloatComplex64(float64(from), defaultValue...)
}

// FloatComplex64 convert value from float64 to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatComplex64(from float64, defaultValue ...complex64) Complex64Accessor {
	nv := &NullComplex64{}
	if safe := isSafeFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultComplex64(nv, defaultValue...) {
			return nv
		}
	}
	v := complex(float32(from), 0)
	nv.P = &v
	return nv
}

// Complex64 convert value from complex128 to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64(from complex128, defaultValue ...complex64) Complex64Accessor {
	nv := &NullComplex64{}
	if safe := isSafeComplex(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultComplex64(nv, defaultValue...) {
			return nv
		}
	}
	v := complex(float32(real(from)), float32(imag(from)))
	nv.P = &v
	return nv
}

// StringComplex64 convert value from string to complex64.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringComplex64(from string, defaultValue ...complex64) Complex64Accessor {
	nv := &NullComplex64{}
	matches := regexpComplex.FindStringSubmatch(from)
	if len(matches) < 3 {
		nv.Error = ErrConvert
		defaultComplex64(nv, defaultValue...)
		return nv
	}
	fr, re := strconv.ParseFloat(matches[1], 32)
	fi, ie := strconv.ParseFloat(matches[2], 32)
	if re != nil || ie != nil {
		nv.Error = ErrUnexpectedValue
	}
	if defaultComplex64(nv, defaultValue...) {
		return nv
	}
	v := complex(float32(fr), float32(fi))
	nv.P = &v
	return nv
}

// StringComplex convert value from string to complex128.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringComplex(from string, defaultValue ...complex128) ComplexAccessor {
	nv := &NullComplex{}
	matches := regexpComplex.FindStringSubmatch(from)
	if len(matches) < 3 {
		nv.Error = ErrConvert
		defaultComplex(nv, defaultValue...)
		return nv
	}
	fr, re := strconv.ParseFloat(matches[1], 64)
	fi, ie := strconv.ParseFloat(matches[2], 64)
	if re != nil || ie != nil {
		nv.Error = ErrUnexpectedValue
	}
	if defaultComplex(nv, defaultValue...) {
		return nv
	}
	v := complex(fr, fi)
	nv.P = &v
	return nv
}

func defaultComplex64(nv *NullComplex64, defaultValue ...complex64) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		nv.P = &defaultValue[0]
		return true
	}
	return false
}

func defaultComplex(nv *NullComplex, defaultValue ...complex128) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		nv.P = &defaultValue[0]
		return true
	}
	return false
}
