package typ

import (
	"reflect"
	"strconv"
)

// Convert interface value to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Float32(defaultValue ...float32) (nv NullFloat32) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toFloat(reflect.Float32)
	nv = NullFloat32{Error: valueTo.Error}
	if defaultFloat32(&nv, defaultValue...) {
		return
	}
	v := float32(valueTo.V())
	nv.P = &v
	return
}

// Convert interface value to float64.
// Returns value if type can safely converted, otherwise error & default value in result values
func (t *Type) Float(defaultValue ...float64) (nv NullFloat) {
	if nv.Error = t.err; t.err != nil {
		return
	}
	valueTo := t.toFloat(reflect.Float64)
	nv = NullFloat{Error: valueTo.Error}
	if defaultFloat(&nv, defaultValue...) {
		return
	}
	v := valueTo.V()
	nv.P = &v
	return
}

// Convert interface value to any float type.
// Returns error if type can't safely converted
func (t *Type) toFloat(typeTo reflect.Kind) (nv NullFloat) {
	if !t.rv.IsValid() || !isFloat(typeTo) {
		nv.Error = ErrConvert
		return
	}
	switch {
	case t.IsString(true):
		value, err := strconv.ParseFloat(t.rv.String(), bitSizeMap[typeTo])
		nv.P = &value
		nv.Error = err
		return
	case t.IsFloat(true):
		floatValue := t.rv.Float()
		nv.P = &floatValue
		if !isSafeFloat(floatValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsInt(true):
		intValue := t.rv.Int()
		v := float64(intValue)
		nv.P = &v
		if !isSafeIntToFloat(intValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsUint(true):
		uintValue := t.rv.Uint()
		v := float64(uintValue)
		nv.P = &v
		if !isSafeUintToFloat(uintValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsComplex(true):
		complexValue := t.rv.Complex()
		v := float64(real(complexValue))
		nv.P = &v
		if !isSafeComplexToFloat(complexValue, bitSizeMap[typeTo]) {
			nv.Error = ErrConvert
		}
		return
	case t.IsBool(true):
		var v float64
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

// Convert value from int to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntFloat32(from int64, defaultValue ...float32) (nv NullFloat32) {
	if safe := isSafeIntToFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultFloat32(&nv, defaultValue...) {
			return
		}
	}
	v := float32(from)
	nv.P = &v
	return
}

// Convert value from int to float64.
// Returns value if type can safely converted, otherwise error & default value in result values
func IntFloat(from int64, defaultValue ...float64) (nv NullFloat) {
	if safe := isSafeIntToFloat(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultFloat(&nv, defaultValue...) {
			return
		}
	}
	v := float64(from)
	nv.P = &v
	return
}

// Convert value from uint to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintFloat32(from uint64, defaultValue ...float32) (nv NullFloat32) {
	if safe := isSafeUintToFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultFloat32(&nv, defaultValue...) {
			return
		}
	}
	v := float32(from)
	nv.P = &v
	return
}

// Convert value from uint to float64.
// Returns value if type can safely converted, otherwise error & default value in result values
func UintFloat(from uint64, defaultValue ...float64) (nv NullFloat) {
	if safe := isSafeUintToFloat(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultFloat(&nv, defaultValue...) {
			return
		}
	}
	v := float64(from)
	nv.P = &v
	return
}

// Convert value from float64 to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Float32(from float64, defaultValue ...float32) (nv NullFloat32) {
	if safe := isSafeFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultFloat32(&nv, defaultValue...) {
			return
		}
	}
	v := float32(from)
	nv.P = &v
	return
}

// Convert value from complex64 to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Float32(from complex64, defaultValue ...float32) (nv NullFloat32) {
	if safe := isSafeComplexToFloat(complex128(from), 32); !safe {
		nv.Error = ErrConvert
		if defaultFloat32(&nv, defaultValue...) {
			return
		}
	}
	v := float32(real(from))
	nv.P = &v
	return
}

// Convert value from complex64 to float64.
// Returns value if type can safely converted, otherwise error & default value in result values
func Complex64Float64(from complex64, defaultValue ...float64) (nv NullFloat) {
	if safe := isSafeComplexToFloat(complex128(from), 64); !safe {
		nv.Error = ErrConvert
		if defaultFloat(&nv, defaultValue...) {
			return
		}
	}
	v := float64(real(from))
	nv.P = &v
	return
}

// Convert value from complex128 to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexFloat32(from complex128, defaultValue ...float32) (nv NullFloat32) {
	if safe := isSafeComplexToFloat(from, 32); !safe {
		nv.Error = ErrConvert
		if defaultFloat32(&nv, defaultValue...) {
			return
		}
	}
	v := float32(real(from))
	nv.P = &v
	return
}

// Convert value from complex128 to float64.
// Returns value if type can safely converted, otherwise error & default value in result values
func ComplexFloat64(from complex128, defaultValue ...float64) (nv NullFloat) {
	if safe := isSafeComplexToFloat(from, 64); !safe {
		nv.Error = ErrConvert
		if defaultFloat(&nv, defaultValue...) {
			return
		}
	}
	v := real(from)
	nv.P = &v
	return
}

// Convert value from string to float32.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringFloat32(from string, defaultValue ...float32) (nv NullFloat32) {
	pv, err := strconv.ParseFloat(from, 32)
	nv.Error = err
	if defaultFloat32(&nv, defaultValue...) {
		return
	}
	v := float32(pv)
	nv.P = &v
	return
}

// Convert value from string to float64.
// Returns value if type can safely converted, otherwise error & default value in result values
func StringFloat(from string, defaultValue ...float64) (nv NullFloat) {
	pv, err := strconv.ParseFloat(from, 64)
	nv.Error = err
	if defaultFloat(&nv, defaultValue...) {
		return
	}
	v := pv
	nv.P = &v
	return
}

func defaultFloat32(nv *NullFloat32, defaultValue ...float32) bool {
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

func defaultFloat(nv *NullFloat, defaultValue ...float64) bool {
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
