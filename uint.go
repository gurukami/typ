package typ

import (
	"reflect"
	"strconv"
)

// Convert interface value to any unsigned integer type.
// Returns error if type can't safely converted
func (t *Type) toUint(typeTo reflect.Kind) Uint64Accessor {
	nv := &NullUint64{}
	if !t.rv.IsValid() || !isUint(typeTo) {
		nv.Error = ErrConvert
		return nv
	}
	switch {
	case t.IsString(true):
		value, err := strconv.ParseUint(t.rv.String(), 0, bitSizeMap[typeTo])
		nv.Error = err
		v := value
		nv.P = &v
		return nv
	case t.IsUint(true):
		uintValue := t.rv.Uint()
		v := uintValue
		nv.P = &v
		if !isSafeUint(uintValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsInt(true):
		intValue := t.rv.Int()
		v := uint64(intValue)
		nv.P = &v
		if !isSafeIntToUint(intValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsFloat(true):
		floatValue := t.rv.Float()
		v := uint64(floatValue)
		nv.P = &v
		if !isSafeFloatToUint(floatValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsComplex(true):
		complexValue := t.rv.Complex()
		v := uint64(real(complexValue))
		nv.P = &v
		if !isSafeComplexToUint(complexValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsBool(true):
		var v uint64
		if t.rv.Bool() {
			v = 1
			nv.P = &v
		} else {
			nv.P = &v
		}
		return nv
	}
	nv.Error = ErrConvert
	return nv
}

// Uint convert interface value to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint(defaultValue ...uint) UintAccessor {
	nv := &NullUint{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toUint(reflect.Uint)
	nv = &NullUint{UintCommon{Error: valueTo.Err()}}
	if defaultUint(nv, defaultValue...) {
		return nv
	}
	v := uint(valueTo.V())
	nv.P = &v
	return nv
}

// Uint8 convert interface value to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint8(defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toUint(reflect.Uint8)
	nv = &NullUint8{Uint8Common{Error: valueTo.Err()}}
	if defaultUint8(nv, defaultValue...) {
		return nv
	}
	v := uint8(valueTo.V())
	nv.P = &v
	return nv
}

// Uint16 convert interface value to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint16(defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toUint(reflect.Uint16)
	nv = &NullUint16{Uint16Common{Error: valueTo.Err()}}
	if defaultUint16(nv, defaultValue...) {
		return nv
	}
	v := uint16(valueTo.V())
	nv.P = &v
	return nv
}

// Uint32 convert interface value to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint32(defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toUint(reflect.Uint32)
	nv = &NullUint32{Uint32Common{Error: valueTo.Err()}}
	if defaultUint32(nv, defaultValue...) {
		return nv
	}
	v := uint32(valueTo.V())
	nv.P = &v
	return nv
}

// Uint64 convert interface value to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint64(defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toUint(reflect.Uint64)
	nv = &NullUint64{Uint64Common{Error: valueTo.Err()}}
	if defaultUint64(nv, defaultValue...) {
		return nv
	}
	v := uint64(valueTo.V())
	nv.P = &v
	return nv
}

// UintInt64 convert value from uint64 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt64(from uint64, defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	if safe := isSafeUintToInt(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt64(nv, defaultValue...) {
			return nv
		}
	}
	v := int64(from)
	nv.P = &v
	return nv
}

// UintInt32 convert value from uint64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt32(from uint64, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	if safe := isSafeUintToInt(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(nv, defaultValue...) {
			return nv
		}
	}
	v := int32(from)
	nv.P = &v
	return nv
}

// UintInt16 convert value from uint64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt16(from uint64, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	if safe := isSafeUintToInt(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(nv, defaultValue...) {
			return nv
		}
	}
	v := int16(from)
	nv.P = &v
	return nv
}

// UintInt8 convert value from uint64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt8(from uint64, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	if safe := isSafeUintToInt(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(nv, defaultValue...) {
			return nv
		}
	}
	v := int8(from)
	nv.P = &v
	return nv
}

// Uint32 convert value from uint64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Uint32(from uint64, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	if safe := isSafeUint(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(nv, defaultValue...) {
			return nv
		}
	}
	v := uint32(from)
	nv.P = &v
	return nv
}

// Uint16 convert value from uint64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Uint16(from uint64, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	if safe := isSafeUint(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(nv, defaultValue...) {
			return nv
		}
	}
	v := uint16(from)
	nv.P = &v
	return nv
}

// Uint8 convert value from uint64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Uint8(from uint64, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	if safe := isSafeUint(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(nv, defaultValue...) {
			return nv
		}
	}
	v := uint8(from)
	nv.P = &v
	return nv
}

// Float32Uint convert value from float32 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint(from float32, defaultValue ...uint) UintAccessor {
	nv := &NullUint{}
	if safe := isSafeFloatToUint(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint(nv, defaultValue...) {
			return nv
		}
	}
	v := uint(from)
	nv.P = &v
	return nv
}

// Float32Uint64 convert value from float32 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint64(from float32, defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	if safe := isSafeFloatToUint(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint64(nv, defaultValue...) {
			return nv
		}
	}
	v := uint64(from)
	nv.P = &v
	return nv
}

// Float32Uint32 convert value from float32 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint32(from float32, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	if safe := isSafeFloatToUint(float64(from), 32, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(nv, defaultValue...) {
			return nv
		}
	}
	v := uint32(from)
	nv.P = &v
	return nv
}

// Float32Uint16 convert value from float32 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint16(from float32, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	if safe := isSafeFloatToUint(float64(from), 32, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(nv, defaultValue...) {
			return nv
		}
	}
	v := uint16(from)
	nv.P = &v
	return nv
}

// Float32Uint8 convert value from float32 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint8(from float32, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	if safe := isSafeFloatToUint(float64(from), 32, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(nv, defaultValue...) {
			return nv
		}
	}
	v := uint8(from)
	nv.P = &v
	return nv
}

// FloatUint convert value from float64 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint(from float64, defaultValue ...uint) UintAccessor {
	nv := &NullUint{}
	if safe := isSafeFloatToUint(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint(nv, defaultValue...) {
			return nv
		}
	}
	v := uint(from)
	nv.P = &v
	return nv
}

// FloatUint64 convert value from float64 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint64(from float64, defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	if safe := isSafeFloatToUint(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint64(nv, defaultValue...) {
			return nv
		}
	}
	v := uint64(from)
	nv.P = &v
	return nv
}

// FloatUint32 convert value from float64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint32(from float64, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	if safe := isSafeFloatToUint(float64(from), 64, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(nv, defaultValue...) {
			return nv
		}
	}
	v := uint32(from)
	nv.P = &v
	return nv
}

// FloatUint16 convert value from float64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint16(from float64, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	if safe := isSafeFloatToUint(float64(from), 64, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(nv, defaultValue...) {
			return nv
		}
	}
	v := uint16(from)
	nv.P = &v
	return nv
}

// FloatUint8 convert value from float64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint8(from float64, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	if safe := isSafeFloatToUint(float64(from), 64, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(nv, defaultValue...) {
			return nv
		}
	}
	v := uint8(from)
	nv.P = &v
	return nv
}

// Complex64Uint convert value from complex64 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint(from complex64, defaultValue ...uint) UintAccessor {
	nv := &NullUint{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Uint64 convert value from complex64 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint64(from complex64, defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint64(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint64(nv, defaultValue...) {
		v := defaultValue[0]
		nv.P = &v
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Uint32 convert value from complex64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint32(from complex64, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint32(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint32(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Uint16 convert value from complex64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint16(from complex64, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint16(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint16(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Uint8 convert value from complex64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint8(from complex64, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint8(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint8(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexUint convert value from complex128 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint(from complex128, defaultValue ...uint) UintAccessor {
	nv := &NullUint{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexUint64 convert value from complex128 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint64(from complex128, defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint64(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint64(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexUint32 convert value from complex128 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint32(from complex128, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint32(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint32(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexUint16 convert value from complex128 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint16(from complex128, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint16(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint16(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexUint8 convert value from complex128 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint8(from complex128, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint8(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultUint8(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// StringUint convert value from string to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint(from string, defaultValue ...uint) UintAccessor {
	nv := &NullUint{}
	pv, err := strconv.ParseUint(from, 0, 64)
	nv.Error = err
	if defaultUint(nv, defaultValue...) {
		return nv
	}
	v := uint(pv)
	nv.P = &v
	return nv
}

// StringUint64 convert value from string to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint64(from string, defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	pv, err := strconv.ParseUint(from, 0, 64)
	nv.Error = err
	if defaultUint64(nv, defaultValue...) {
		v := defaultValue[0]
		nv.P = &v
		return nv
	}
	v := uint64(pv)
	nv.P = &v
	return nv
}

// StringUint32 convert value from string to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint32(from string, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	pv, err := strconv.ParseUint(from, 0, 32)
	nv.Error = err
	if defaultUint32(nv, defaultValue...) {
		return nv
	}
	v := uint32(pv)
	nv.P = &v
	return nv
}

// StringUint16 convert value from string to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint16(from string, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	pv, err := strconv.ParseUint(from, 0, 16)
	nv.Error = err
	if defaultUint16(nv, defaultValue...) {
		return nv
	}
	v := uint16(pv)
	nv.P = &v
	return nv
}

// StringUint8 convert value from string to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint8(from string, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	pv, err := strconv.ParseUint(from, 0, 8)
	nv.Error = err
	if defaultUint8(nv, defaultValue...) {
		return nv
	}
	v := uint8(pv)
	nv.P = &v
	return nv
}

func defaultUint8(nv *NullUint8, defaultValue ...uint8) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint16(nv *NullUint16, defaultValue ...uint16) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint32(nv *NullUint32, defaultValue ...uint32) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint64(nv *NullUint64, defaultValue ...uint64) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint(nv *NullUint, defaultValue ...uint) bool {
	if len(defaultValue) > 1 {
		nv.Error = ErrDefaultValue
		return true
	}
	if !nv.Valid() && len(defaultValue) > 0 {
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}
