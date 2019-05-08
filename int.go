package typ

import (
	"reflect"
	"strconv"
)

// Convert interface value to any signed integer type.
// Returns error if type can't safely converted
func (t *Type) toInt(typeTo reflect.Kind) Int64Accessor {
	nv := &NullInt64{}
	if !t.rv.IsValid() || !isInt(typeTo) {
		nv.Error = ErrConvert
		return nv
	}
	switch {
	case t.IsString(true):
		value, err := strconv.ParseInt(t.rv.String(), 0, bitSizeMap[typeTo])
		nv.Error = err
		v := value
		nv.P = &v
		return nv
	case t.IsInt(true):
		intValue := t.rv.Int()
		v := intValue
		nv.P = &v
		if !isSafeInt(intValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsUint(true):
		uintValue := t.rv.Uint()
		v := int64(uintValue)
		nv.P = &v
		if !isSafeUintToInt(uintValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsFloat(true):
		floatValue := t.rv.Float()
		v := int64(floatValue)
		nv.P = &v
		if !isSafeFloatToInt(floatValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsComplex(true):
		complexValue := t.rv.Complex()
		v := int64(real(complexValue))
		nv.P = &v
		if !isSafeComplexToInt(complexValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return nv
	case t.IsBool(true):
		var v int64
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

// Int convert interface value to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int(defaultValue ...int) IntAccessor {
	nv := &NullInt{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toInt(reflect.Int)
	nv = &NullInt{IntCommon{Error: valueTo.Err()}}
	if defaultInt(nv, defaultValue...) {
		return nv
	}
	v := int(valueTo.V())
	nv.P = &v
	return nv
}

// Int8 convert interface value to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int8(defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toInt(reflect.Int8)
	nv = &NullInt8{Int8Common{Error: valueTo.Err()}}
	if defaultInt8(nv, defaultValue...) {
		return nv
	}
	v := int8(valueTo.V())
	nv.P = &v
	return nv
}

// Int16 convert interface value to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int16(defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toInt(reflect.Int16)
	nv = &NullInt16{Int16Common{Error: valueTo.Err()}}
	if defaultInt16(nv, defaultValue...) {
		return nv
	}
	v := int16(valueTo.V())
	nv.P = &v
	return nv
}

// Int32 convert interface value to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int32(defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toInt(reflect.Int32)
	nv = &NullInt32{Int32Common{Error: valueTo.Err()}}
	if defaultInt32(nv, defaultValue...) {
		return nv
	}
	v := int32(valueTo.V())
	nv.P = &v
	return nv
}

// Int64 convert interface value to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int64(defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	if nv.Error = t.err; t.err != nil {
		return nv
	}
	valueTo := t.toInt(reflect.Int64)
	nv = &NullInt64{Int64Common{Error: valueTo.Err()}}
	if defaultInt64(nv, defaultValue...) {
		return nv
	}
	v := int64(valueTo.V())
	nv.P = &v
	return nv
}

// Int32 convert value from int64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Int32(from int64, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	if safe := isSafeInt(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(nv, defaultValue...) {
			return nv
		}
	}
	v := int32(from)
	nv.P = &v
	return nv
}

// Int16 convert value from int64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Int16(from int64, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	if safe := isSafeInt(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(nv, defaultValue...) {
			return nv
		}
	}
	v := int16(from)
	nv.P = &v
	return nv
}

// Int8 convert value from int64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Int8(from int64, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	if safe := isSafeInt(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(nv, defaultValue...) {
			return nv
		}
	}
	v := int8(from)
	nv.P = &v
	return nv
}

// IntUint64 convert value from int64 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint64(from int64, defaultValue ...uint64) Uint64Accessor {
	nv := &NullUint64{}
	if safe := isSafeIntToUint(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint64(nv, defaultValue...) {
			return nv
		}
	}
	v := uint64(from)
	nv.P = &v
	return nv
}

// IntUint32 convert value from int64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint32(from int64, defaultValue ...uint32) Uint32Accessor {
	nv := &NullUint32{}
	if safe := isSafeIntToUint(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(nv, defaultValue...) {
			return nv
		}
	}
	v := uint32(from)
	nv.P = &v
	return nv
}

// IntUint16 convert value from int64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint16(from int64, defaultValue ...uint16) Uint16Accessor {
	nv := &NullUint16{}
	if safe := isSafeIntToUint(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(nv, defaultValue...) {
			return nv
		}
	}
	v := uint16(from)
	nv.P = &v
	return nv
}

// IntUint8 convert value from int64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint8(from int64, defaultValue ...uint8) Uint8Accessor {
	nv := &NullUint8{}
	if safe := isSafeIntToUint(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(nv, defaultValue...) {
			return nv
		}
	}
	v := uint8(from)
	nv.P = &v
	return nv
}

// Float32Int convert value from float32 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int(from float32, defaultValue ...int) IntAccessor {
	nv := &NullInt{}
	if safe := isSafeFloatToInt(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt(nv, defaultValue...) {
			return nv
		}
	}
	v := int(from)
	nv.P = &v
	return nv
}

// Float32Int64 convert value from float32 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int64(from float32, defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	if safe := isSafeFloatToInt(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt64(nv, defaultValue...) {
			return nv
		}
	}
	v := int64(from)
	nv.P = &v
	return nv
}

// Float32Int32 convert value from float32 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int32(from float32, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	if safe := isSafeFloatToInt(float64(from), 32, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(nv, defaultValue...) {
			return nv
		}
	}
	v := int32(from)
	nv.P = &v
	return nv
}

// Float32Int16 convert value from float32 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int16(from float32, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	if safe := isSafeFloatToInt(float64(from), 32, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(nv, defaultValue...) {
			return nv
		}
	}
	v := int16(from)
	nv.P = &v
	return nv
}

// Float32Int8 convert value from float32 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int8(from float32, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	if safe := isSafeFloatToInt(float64(from), 32, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(nv, defaultValue...) {
			return nv
		}
	}
	v := int8(from)
	nv.P = &v
	return nv
}

// FloatInt convert value from float64 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt(from float64, defaultValue ...int) IntAccessor {
	nv := &NullInt{}
	if safe := isSafeFloatToInt(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt(nv, defaultValue...) {
			return nv
		}
	}
	v := int(from)
	nv.P = &v
	return nv
}

// FloatInt64 convert value from float64 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt64(from float64, defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	if safe := isSafeFloatToInt(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt64(nv, defaultValue...) {
			return nv
		}
	}
	v := int64(from)
	nv.P = &v
	return nv
}

// FloatInt32 convert value from float64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt32(from float64, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	if safe := isSafeFloatToInt(float64(from), 64, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(nv, defaultValue...) {
			return nv
		}
	}
	v := int32(from)
	nv.P = &v
	return nv
}

// FloatInt16 convert value from float64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt16(from float64, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	if safe := isSafeFloatToInt(float64(from), 32, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(nv, defaultValue...) {
			return nv
		}
	}
	v := int16(from)
	nv.P = &v
	return nv
}

// FloatInt8 convert value from float64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt8(from float64, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	if safe := isSafeFloatToInt(float64(from), 32, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(nv, defaultValue...) {
			return nv
		}
	}
	v := int8(from)
	nv.P = &v
	return nv
}

// Complex64Int convert value from complex64 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int(from complex64, defaultValue ...int) IntAccessor {
	nv := &NullInt{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Int64 convert value from complex64 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int64(from complex64, defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int64(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt64(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Int32 convert value from complex64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int32(from complex64, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int32(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt32(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Int16 convert value from complex64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int16(from complex64, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int16(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt16(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// Complex64Int8 convert value from complex64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int8(from complex64, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int8(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt8(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexInt convert value from complex128 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt(from complex128, defaultValue ...int) IntAccessor {
	nv := &NullInt{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexInt64 convert value from complex128 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt64(from complex128, defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt64(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt64(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexInt32 convert value from complex128 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt32(from complex128, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt32(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt32(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexInt16 convert value from complex128 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt16(from complex128, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt16(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt16(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// ComplexInt8 convert value from complex128 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt8(from complex128, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt8(fr, defaultValue...)
	nv.Error = cv.Err()
	if vi != 0 {
	}
	if defaultInt8(nv, defaultValue...) {
		return nv
	}
	v := cv.V()
	nv.P = &v
	return nv
}

// StringInt convert value from string to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt(from string, defaultValue ...int) IntAccessor {
	nv := &NullInt{}
	pv, err := strconv.ParseInt(from, 0, 64)
	nv.Error = err
	if defaultInt(nv, defaultValue...) {
		return nv
	}
	v := int(pv)
	nv.P = &v
	return nv
}

// StringInt64 convert value from string to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt64(from string, defaultValue ...int64) Int64Accessor {
	nv := &NullInt64{}
	pv, err := strconv.ParseInt(from, 0, 64)
	nv.Error = err
	if defaultInt64(nv, defaultValue...) {
		return nv
	}
	v := int64(pv)
	nv.P = &v
	return nv
}

// StringInt32 convert value from string to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt32(from string, defaultValue ...int32) Int32Accessor {
	nv := &NullInt32{}
	pv, err := strconv.ParseInt(from, 0, 32)
	nv.Error = err
	if defaultInt32(nv, defaultValue...) {
		return nv
	}
	v := int32(pv)
	nv.P = &v
	return nv
}

// StringInt16 convert value from string to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt16(from string, defaultValue ...int16) Int16Accessor {
	nv := &NullInt16{}
	pv, err := strconv.ParseInt(from, 0, 16)
	nv.Error = err
	if defaultInt16(nv, defaultValue...) {
		return nv
	}
	v := int16(pv)
	nv.P = &v
	return nv
}

// StringInt8 convert value from string to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt8(from string, defaultValue ...int8) Int8Accessor {
	nv := &NullInt8{}
	pv, err := strconv.ParseInt(from, 0, 8)
	nv.Error = err
	if defaultInt8(nv, defaultValue...) {
		v := defaultValue[0]
		nv.P = &v
		return nv
	}
	v := int8(pv)
	nv.P = &v
	return nv
}

func defaultInt8(nv *NullInt8, defaultValue ...int8) bool {
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

func defaultInt16(nv *NullInt16, defaultValue ...int16) bool {
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

func defaultInt32(nv *NullInt32, defaultValue ...int32) bool {
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

func defaultInt64(nv *NullInt64, defaultValue ...int64) bool {
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

func defaultInt(nv *NullInt, defaultValue ...int) bool {
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
