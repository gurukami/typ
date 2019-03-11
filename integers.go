package typ

import (
	"reflect"
	"strconv"
)

// Int convert interface value to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int(defaultValue ...int) (nv NullInt) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toInt(reflect.Int)
	nv = NullInt{Error: valueTo.Error}
	if defaultInt(&nv, defaultValue...) {
		return
	}
	v := int(valueTo.V())
	nv.P = &v
	return
}

// Int8 convert interface value to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int8(defaultValue ...int8) (nv NullInt8) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toInt(reflect.Int8)
	nv = NullInt8{Error: valueTo.Error}
	if defaultInt8(&nv, defaultValue...) {
		return
	}
	v := int8(valueTo.V())
	nv.P = &v
	return
}

// Int16 convert interface value to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int16(defaultValue ...int16) (nv NullInt16) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toInt(reflect.Int16)
	nv = NullInt16{Error: valueTo.Error}
	if defaultInt16(&nv, defaultValue...) {
		return
	}
	v := int16(valueTo.V())
	nv.P = &v
	return
}

// Int32 convert interface value to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int32(defaultValue ...int32) (nv NullInt32) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toInt(reflect.Int32)
	nv = NullInt32{Error: valueTo.Error}
	if defaultInt32(&nv, defaultValue...) {
		return
	}
	v := int32(valueTo.V())
	nv.P = &v
	return
}

// Int64 convert interface value to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Int64(defaultValue ...int64) (nv NullInt64) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toInt(reflect.Int64)
	nv = NullInt64{Error: valueTo.Error}
	if defaultInt64(&nv, defaultValue...) {
		return
	}
	v := int64(valueTo.V())
	nv.P = &v
	return
}

// Uint convert interface value to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint(defaultValue ...uint) (nv NullUint) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toUint(reflect.Uint)
	nv = NullUint{Error: valueTo.Error}
	if defaultUint(&nv, defaultValue...) {
		return
	}
	v := uint(valueTo.V())
	nv.P = &v
	return
}

// Uint8 convert interface value to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint8(defaultValue ...uint8) (nv NullUint8) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toUint(reflect.Uint8)
	nv = NullUint8{Error: valueTo.Error}
	if defaultUint8(&nv, defaultValue...) {
		return
	}
	v := uint8(valueTo.V())
	nv.P = &v
	return
}

// Uint16 convert interface value to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint16(defaultValue ...uint16) (nv NullUint16) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toUint(reflect.Uint16)
	nv = NullUint16{Error: valueTo.Error}
	if defaultUint16(&nv, defaultValue...) {
		return
	}
	v := uint16(valueTo.V())
	nv.P = &v
	return
}

// Uint32 convert interface value to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint32(defaultValue ...uint32) (nv NullUint32) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toUint(reflect.Uint32)
	nv = NullUint32{Error: valueTo.Error}
	if defaultUint32(&nv, defaultValue...) {
		return
	}
	v := uint32(valueTo.V())
	nv.P = &v
	return
}

// Uint64 convert interface value to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Uint64(defaultValue ...uint64) (nv NullUint64) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toUint(reflect.Uint64)
	nv = NullUint64{Error: valueTo.Error}
	if defaultUint64(&nv, defaultValue...) {
		return
	}
	v := uint64(valueTo.V())
	nv.P = &v
	return
}

// Convert interface value to any signed integer type.
// Returns error if type can't safely converted
func (t *Type) toInt(typeTo reflect.Kind) (nv NullInt64) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	if !t.rv.IsValid() || !isInt(typeTo) {
		nv.Error = ErrConvert
		return
	}
	switch {
	case t.IsString(true):
		value, err := strconv.ParseInt(t.rv.String(), 0, bitSizeMap[typeTo])
		nv.Error = err
		v := value
		nv.P = &v
		return
	case t.IsInt(true):
		intValue := t.rv.Int()
		v := intValue
		nv.P = &v
		if !isSafeInt(intValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsUint(true):
		uintValue := t.rv.Uint()
		v := int64(uintValue)
		nv.P = &v
		if !isSafeUintToInt(uintValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsFloat(true):
		floatValue := t.rv.Float()
		v := int64(floatValue)
		nv.P = &v
		if !isSafeFloatToInt(floatValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsComplex(true):
		complexValue := t.rv.Complex()
		v := int64(real(complexValue))
		nv.P = &v
		if !isSafeComplexToInt(complexValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsBool(true):
		var v int64
		if t.rv.Bool() {
			v = 1
			nv.P = &v
		} else {
			nv.P = &v
		}
		return
	}
	nv.Error = ErrConvert
	return
}

// Convert interface value to any unsigned integer type.
// Returns error if type can't safely converted
func (t *Type) toUint(typeTo reflect.Kind) (nv NullUint64) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	if !t.rv.IsValid() || !isUint(typeTo) {
		nv.Error = ErrConvert
		return
	}
	switch {
	case t.IsString(true):
		value, err := strconv.ParseUint(t.rv.String(), 0, bitSizeMap[typeTo])
		nv.Error = err
		v := value
		nv.P = &v
		return
	case t.IsUint(true):
		uintValue := t.rv.Uint()
		v := uintValue
		nv.P = &v
		if !isSafeUint(uintValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsInt(true):
		intValue := t.rv.Int()
		v := uint64(intValue)
		nv.P = &v
		if !isSafeIntToUint(intValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsFloat(true):
		floatValue := t.rv.Float()
		v := uint64(floatValue)
		nv.P = &v
		if !isSafeFloatToUint(floatValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsComplex(true):
		complexValue := t.rv.Complex()
		v := uint64(real(complexValue))
		nv.P = &v
		if !isSafeComplexToUint(complexValue, bitSizeMap[t.Kind()], bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsBool(true):
		var v uint64
		if t.rv.Bool() {
			v = 1
			nv.P = &v
		} else {
			nv.P = &v
		}
		return
	}
	nv.Error = ErrConvert
	return
}

// Int32 convert value from int64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Int32(from int64, defaultValue ...int32) (nv NullInt32) {
	if safe := isSafeInt(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(&nv, defaultValue...) {
			return
		}
	}
	v := int32(from)
	nv.P = &v
	return
}

// Int16 convert value from int64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Int16(from int64, defaultValue ...int16) (nv NullInt16) {
	if safe := isSafeInt(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(&nv, defaultValue...) {
			return
		}
	}
	v := int16(from)
	nv.P = &v
	return
}

// Int8 convert value from int64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Int8(from int64, defaultValue ...int8) (nv NullInt8) {
	if safe := isSafeInt(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(&nv, defaultValue...) {
			return
		}
	}
	v := int8(from)
	nv.P = &v
	return
}

// IntUint64 convert value from int64 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint64(from int64, defaultValue ...uint64) (nv NullUint64) {
	if safe := isSafeIntToUint(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint64(&nv, defaultValue...) {
			return
		}
	}
	v := uint64(from)
	nv.P = &v
	return
}

// IntUint32 convert value from int64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint32(from int64, defaultValue ...uint32) (nv NullUint32) {
	if safe := isSafeIntToUint(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(&nv, defaultValue...) {
			return
		}
	}
	v := uint32(from)
	nv.P = &v
	return
}

// IntUint16 convert value from int64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint16(from int64, defaultValue ...uint16) (nv NullUint16) {
	if safe := isSafeIntToUint(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(&nv, defaultValue...) {
			return
		}
	}
	v := uint16(from)
	nv.P = &v
	return
}

// IntUint8 convert value from int64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntUint8(from int64, defaultValue ...uint8) (nv NullUint8) {
	if safe := isSafeIntToUint(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(&nv, defaultValue...) {
			return
		}
	}
	v := uint8(from)
	nv.P = &v
	return
}

// UintInt64 convert value from uint64 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt64(from uint64, defaultValue ...int64) (nv NullInt64) {
	if safe := isSafeUintToInt(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt64(&nv, defaultValue...) {
			return
		}
	}
	v := int64(from)
	nv.P = &v
	return
}

// UintInt32 convert value from uint64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt32(from uint64, defaultValue ...int32) (nv NullInt32) {
	if safe := isSafeUintToInt(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(&nv, defaultValue...) {
			return
		}
	}
	v := int32(from)
	nv.P = &v
	return
}

// UintInt16 convert value from uint64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt16(from uint64, defaultValue ...int16) (nv NullInt16) {
	if safe := isSafeUintToInt(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(&nv, defaultValue...) {
			return
		}
	}
	v := int16(from)
	nv.P = &v
	return
}

// UintInt8 convert value from uint64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintInt8(from uint64, defaultValue ...int8) (nv NullInt8) {
	if safe := isSafeUintToInt(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(&nv, defaultValue...) {
			return
		}
	}
	v := int8(from)
	nv.P = &v
	return
}

// Uint32 convert value from uint64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Uint32(from uint64, defaultValue ...uint32) (nv NullUint32) {
	if safe := isSafeUint(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(&nv, defaultValue...) {
			return
		}
	}
	v := uint32(from)
	nv.P = &v
	return
}

// Uint16 convert value from uint64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Uint16(from uint64, defaultValue ...uint16) (nv NullUint16) {
	if safe := isSafeUint(from, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(&nv, defaultValue...) {
			return
		}
	}
	v := uint16(from)
	nv.P = &v
	return
}

// Uint8 convert value from uint64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Uint8(from uint64, defaultValue ...uint8) (nv NullUint8) {
	if safe := isSafeUint(from, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(&nv, defaultValue...) {
			return
		}
	}
	v := uint8(from)
	nv.P = &v
	return
}

// Float32Int convert value from float32 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int(from float32, defaultValue ...int) (nv NullInt) {
	if safe := isSafeFloatToInt(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt(&nv, defaultValue...) {
			return
		}
	}
	v := int(from)
	nv.P = &v
	return
}

// Float32Int64 convert value from float32 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int64(from float32, defaultValue ...int64) (nv NullInt64) {
	if safe := isSafeFloatToInt(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt64(&nv, defaultValue...) {
			return
		}
	}
	v := int64(from)
	nv.P = &v
	return
}

// Float32Int32 convert value from float32 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int32(from float32, defaultValue ...int32) (nv NullInt32) {
	if safe := isSafeFloatToInt(float64(from), 32, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(&nv, defaultValue...) {
			return
		}
	}
	v := int32(from)
	nv.P = &v
	return
}

// Float32Int16 convert value from float32 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int16(from float32, defaultValue ...int16) (nv NullInt16) {
	if safe := isSafeFloatToInt(float64(from), 32, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(&nv, defaultValue...) {
			return
		}
	}
	v := int16(from)
	nv.P = &v
	return
}

// Float32Int8 convert value from float32 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Int8(from float32, defaultValue ...int8) (nv NullInt8) {
	if safe := isSafeFloatToInt(float64(from), 32, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(&nv, defaultValue...) {
			return
		}
	}
	v := int8(from)
	nv.P = &v
	return
}

// Float32Uint convert value from float32 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint(from float32, defaultValue ...uint) (nv NullUint) {
	if safe := isSafeFloatToUint(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint(&nv, defaultValue...) {
			return
		}
	}
	v := uint(from)
	nv.P = &v
	return
}

// Float32Uint64 convert value from float32 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint64(from float32, defaultValue ...uint64) (nv NullUint64) {
	if safe := isSafeFloatToUint(float64(from), 32, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint64(&nv, defaultValue...) {
			return
		}
	}
	v := uint64(from)
	nv.P = &v
	return
}

// Float32Uint32 convert value from float32 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint32(from float32, defaultValue ...uint32) (nv NullUint32) {
	if safe := isSafeFloatToUint(float64(from), 32, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(&nv, defaultValue...) {
			return
		}
	}
	v := uint32(from)
	nv.P = &v
	return
}

// Float32Uint16 convert value from float32 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint16(from float32, defaultValue ...uint16) (nv NullUint16) {
	if safe := isSafeFloatToUint(float64(from), 32, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(&nv, defaultValue...) {
			return
		}
	}
	v := uint16(from)
	nv.P = &v
	return
}

// Float32Uint8 convert value from float32 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32Uint8(from float32, defaultValue ...uint8) (nv NullUint8) {
	if safe := isSafeFloatToUint(float64(from), 32, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(&nv, defaultValue...) {
			return
		}
	}
	v := uint8(from)
	nv.P = &v
	return
}

// FloatInt convert value from float64 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt(from float64, defaultValue ...int) (nv NullInt) {
	if safe := isSafeFloatToInt(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt(&nv, defaultValue...) {
			return
		}
	}
	v := int(from)
	nv.P = &v
	return
}

// FloatInt64 convert value from float64 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt64(from float64, defaultValue ...int64) (nv NullInt64) {
	if safe := isSafeFloatToInt(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultInt64(&nv, defaultValue...) {
			return
		}
	}
	v := int64(from)
	nv.P = &v
	return
}

// FloatInt32 convert value from float64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt32(from float64, defaultValue ...int32) (nv NullInt32) {
	if safe := isSafeFloatToInt(float64(from), 64, 32); !safe {
		nv.Error = ErrConvert
		if defaultInt32(&nv, defaultValue...) {
			return
		}
	}
	v := int32(from)
	nv.P = &v
	return
}

// FloatInt16 convert value from float64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt16(from float64, defaultValue ...int16) (nv NullInt16) {
	if safe := isSafeFloatToInt(float64(from), 32, 16); !safe {
		nv.Error = ErrConvert
		if defaultInt16(&nv, defaultValue...) {
			return
		}
	}
	v := int16(from)
	nv.P = &v
	return
}

// FloatInt8 convert value from float64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatInt8(from float64, defaultValue ...int8) (nv NullInt8) {
	if safe := isSafeFloatToInt(float64(from), 32, 8); !safe {
		nv.Error = ErrConvert
		if defaultInt8(&nv, defaultValue...) {
			return
		}
	}
	v := int8(from)
	nv.P = &v
	return
}

// FloatUint convert value from float64 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint(from float64, defaultValue ...uint) (nv NullUint) {
	if safe := isSafeFloatToUint(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint(&nv, defaultValue...) {
			return
		}
	}
	v := uint(from)
	nv.P = &v
	return
}

// FloatUint64 convert value from float64 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint64(from float64, defaultValue ...uint64) (nv NullUint64) {
	if safe := isSafeFloatToUint(float64(from), 64, 64); !safe {
		nv.Error = ErrConvert
		if defaultUint64(&nv, defaultValue...) {
			return
		}
	}
	v := uint64(from)
	nv.P = &v
	return
}

// FloatUint32 convert value from float64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint32(from float64, defaultValue ...uint32) (nv NullUint32) {
	if safe := isSafeFloatToUint(float64(from), 64, 32); !safe {
		nv.Error = ErrConvert
		if defaultUint32(&nv, defaultValue...) {
			return
		}
	}
	v := uint32(from)
	nv.P = &v
	return
}

// FloatUint16 convert value from float64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint16(from float64, defaultValue ...uint16) (nv NullUint16) {
	if safe := isSafeFloatToUint(float64(from), 64, 16); !safe {
		nv.Error = ErrConvert
		if defaultUint16(&nv, defaultValue...) {
			return
		}
	}
	v := uint16(from)
	nv.P = &v
	return
}

// FloatUint8 convert value from float64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func FloatUint8(from float64, defaultValue ...uint8) (nv NullUint8) {
	if safe := isSafeFloatToUint(float64(from), 64, 8); !safe {
		nv.Error = ErrConvert
		if defaultUint8(&nv, defaultValue...) {
			return
		}
	}
	v := uint8(from)
	nv.P = &v
	return
}

// Complex64Int convert value from complex64 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int(from complex64, defaultValue ...int) (nv NullInt) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Int64 convert value from complex64 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int64(from complex64, defaultValue ...int64) (nv NullInt64) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int64(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt64(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Int32 convert value from complex64 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int32(from complex64, defaultValue ...int32) (nv NullInt32) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int32(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt32(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Int16 convert value from complex64 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int16(from complex64, defaultValue ...int16) (nv NullInt16) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int16(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt16(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Int8 convert value from complex64 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Int8(from complex64, defaultValue ...int8) (nv NullInt8) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Int8(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt8(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Uint convert value from complex64 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint(from complex64, defaultValue ...uint) (nv NullUint) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Uint64 convert value from complex64 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint64(from complex64, defaultValue ...uint64) (nv NullUint64) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint64(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint64(&nv, defaultValue...) {
		v := defaultValue[0]
		nv.P = &v
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Uint32 convert value from complex64 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint32(from complex64, defaultValue ...uint32) (nv NullUint32) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint32(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint32(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Uint16 convert value from complex64 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint16(from complex64, defaultValue ...uint16) (nv NullUint16) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint16(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint16(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// Complex64Uint8 convert value from complex64 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Uint8(from complex64, defaultValue ...uint8) (nv NullUint8) {
	fr := float32(real(from))
	vi := float32(imag(from))
	cv := Float32Uint8(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint8(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexInt convert value from complex128 to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt(from complex128, defaultValue ...int) (nv NullInt) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexInt64 convert value from complex128 to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt64(from complex128, defaultValue ...int64) (nv NullInt64) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt64(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt64(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexInt32 convert value from complex128 to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt32(from complex128, defaultValue ...int32) (nv NullInt32) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt32(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt32(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexInt16 convert value from complex128 to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt16(from complex128, defaultValue ...int16) (nv NullInt16) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt16(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt16(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexInt8 convert value from complex128 to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexInt8(from complex128, defaultValue ...int8) (nv NullInt8) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatInt8(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultInt8(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexUint convert value from complex128 to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint(from complex128, defaultValue ...uint) (nv NullUint) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexUint64 convert value from complex128 to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint64(from complex128, defaultValue ...uint64) (nv NullUint64) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint64(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint64(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexUint32 convert value from complex128 to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint32(from complex128, defaultValue ...uint32) (nv NullUint32) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint32(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint32(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexUint16 convert value from complex128 to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint16(from complex128, defaultValue ...uint16) (nv NullUint16) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint16(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint16(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// ComplexUint8 convert value from complex128 to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexUint8(from complex128, defaultValue ...uint8) (nv NullUint8) {
	fr, vi := float64(real(from)), float64(imag(from))
	cv := FloatUint8(fr, defaultValue...)
	nv.Error = cv.Error
	if vi != 0 {
	}
	if defaultUint8(&nv, defaultValue...) {
		return
	}
	v := cv.V()
	nv.P = &v
	return
}

// StringInt convert value from string to int.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt(from string, defaultValue ...int) (nv NullInt) {
	pv, err := strconv.ParseInt(from, 0, 64)
	nv.Error = err
	if defaultInt(&nv, defaultValue...) {
		return
	}
	v := int(pv)
	nv.P = &v
	return
}

// StringInt64 convert value from string to int64.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt64(from string, defaultValue ...int64) (nv NullInt64) {
	pv, err := strconv.ParseInt(from, 0, 64)
	nv.Error = err
	if defaultInt64(&nv, defaultValue...) {
		return
	}
	v := int64(pv)
	nv.P = &v
	return
}

// StringInt32 convert value from string to int32.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt32(from string, defaultValue ...int32) (nv NullInt32) {
	pv, err := strconv.ParseInt(from, 0, 32)
	nv.Error = err
	if defaultInt32(&nv, defaultValue...) {
		return
	}
	v := int32(pv)
	nv.P = &v
	return
}

// StringInt16 convert value from string to int16.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt16(from string, defaultValue ...int16) (nv NullInt16) {
	pv, err := strconv.ParseInt(from, 0, 16)
	nv.Error = err
	if defaultInt16(&nv, defaultValue...) {
		return
	}
	v := int16(pv)
	nv.P = &v
	return
}

// StringInt8 convert value from string to int8.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringInt8(from string, defaultValue ...int8) (nv NullInt8) {
	pv, err := strconv.ParseInt(from, 0, 8)
	nv.Error = err
	if defaultInt8(&nv, defaultValue...) {
		v := defaultValue[0]
		nv.P = &v
		return
	}
	v := int8(pv)
	nv.P = &v
	return
}

// StringUint convert value from string to uint.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint(from string, defaultValue ...uint) (nv NullUint) {
	pv, err := strconv.ParseUint(from, 0, 64)
	nv.Error = err
	if defaultUint(&nv, defaultValue...) {
		return
	}
	v := uint(pv)
	nv.P = &v
	return
}

// StringUint64 convert value from string to uint64.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint64(from string, defaultValue ...uint64) (nv NullUint64) {
	pv, err := strconv.ParseUint(from, 0, 64)
	nv.Error = err
	if defaultUint64(&nv, defaultValue...) {
		v := defaultValue[0]
		nv.P = &v
		return
	}
	v := uint64(pv)
	nv.P = &v
	return
}

// StringUint32 convert value from string to uint32.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint32(from string, defaultValue ...uint32) (nv NullUint32) {
	pv, err := strconv.ParseUint(from, 0, 32)
	nv.Error = err
	if defaultUint32(&nv, defaultValue...) {
		return
	}
	v := uint32(pv)
	nv.P = &v
	return
}

// StringUint16 convert value from string to uint16.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint16(from string, defaultValue ...uint16) (nv NullUint16) {
	pv, err := strconv.ParseUint(from, 0, 16)
	nv.Error = err
	if defaultUint16(&nv, defaultValue...) {
		return
	}
	v := uint16(pv)
	nv.P = &v
	return
}

// StringUint8 convert value from string to uint8.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringUint8(from string, defaultValue ...uint8) (nv NullUint8) {
	pv, err := strconv.ParseUint(from, 0, 8)
	nv.Error = err
	if defaultUint8(&nv, defaultValue...) {
		return
	}
	v := uint8(pv)
	nv.P = &v
	return
}

func defaultInt8(nv *NullInt8, defaultValue ...int8) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultInt16(nv *NullInt16, defaultValue ...int16) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultInt32(nv *NullInt32, defaultValue ...int32) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultInt64(nv *NullInt64, defaultValue ...int64) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultInt(nv *NullInt, defaultValue ...int) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint8(nv *NullUint8, defaultValue ...uint8) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint16(nv *NullUint16, defaultValue ...uint16) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint32(nv *NullUint32, defaultValue ...uint32) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint64(nv *NullUint64, defaultValue ...uint64) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}

func defaultUint(nv *NullUint, defaultValue ...uint) bool {
	if !nv.Valid() && len(defaultValue) > 0 {
		if len(defaultValue) > 1 {
			nv.Error = ErrDefaultValue
			return true
		}
		v := defaultValue[0]
		nv.P = &v
		return true
	}
	return false
}
