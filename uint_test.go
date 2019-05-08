package typ

import (
	"database/sql/driver"
	"reflect"
	"strconv"
	"testing"
)

var uintReflectTypes = []reflect.Type{
	getDefaultType(reflect.Uint),
	getDefaultType(reflect.Uint64),
	getDefaultType(reflect.Uint32),
	getDefaultType(reflect.Uint16),
	getDefaultType(reflect.Uint8),
}

func init() {
	// Test Data
	matrixSuite.Register(getDefaultType(reflect.Uint), []dataItem{
		{reflect.ValueOf(uint(MaxUint)), nil},
		{reflect.ValueOf(uint(0)), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Uint64), []dataItem{
		{reflect.ValueOf(uint64(MaxUint)), nil},
		{reflect.ValueOf(uint64(0)), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Uint32), []dataItem{
		{reflect.ValueOf(uint32(MaxUint32)), nil},
		{reflect.ValueOf(uint32(0)), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Uint16), []dataItem{
		{reflect.ValueOf(uint16(MaxUint16)), nil},
		{reflect.ValueOf(uint16(0)), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Uint8), []dataItem{
		{reflect.ValueOf(uint8(MaxUint8)), nil},
		{reflect.ValueOf(uint8(0)), nil},
	})
	// Converters
	// - to bool
	uintBoolConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, b := reflect.ValueOf(from), false
		if rv.Uint() != 0 {
			b = true
		}
		return b, true
	}
	matrixSuite.SetConverters(uintReflectTypes, boolReflectTypes, uintBoolConverter)
	// - to complex
	uintComplexConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var s bool
		rv, c := reflect.ValueOf(from), complex128(0)
		i := rv.Uint()
		if s = isSafeUintToFloat(i, bitSizeMap[to.Kind()]); s {
			c = complex(float64(i), 0)
		}
		switch to.Kind() {
		case reflect.Complex64:
			return complex64(c), s
		case reflect.Complex128:
			return complex128(c), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(uintReflectTypes, complexReflectTypes, uintComplexConverter)
	// - to float
	uintFloatConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var s bool
		rv, f := reflect.ValueOf(from), float64(0)
		i := rv.Uint()
		if s = isSafeUintToFloat(i, bitSizeMap[to.Kind()]); s {
			f = float64(i)
		}
		switch to.Kind() {
		case reflect.Float32:
			return float32(f), s
		case reflect.Float64:
			return float64(f), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(uintReflectTypes, floatReflectTypes, uintFloatConverter)
	// - to string
	uintStringConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, str := reflect.ValueOf(from), ""
		str = strconv.FormatUint(rv.Uint(), 10)
		return str, true
	}
	matrixSuite.SetConverters(uintReflectTypes, stringReflectTypes, uintStringConverter)
	// - to &NullUint*{}, &NotNullUint*{}
	matrixSuite.SetConverters(uintReflectTypes, nullUintReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Uint && to == reflect.TypeOf(&NullUint{}):
			v := uint(rv.Uint())
			return &NullUint{UintCommon{P: &v}}, true
		case rv.Kind() == reflect.Uint64 && to == reflect.TypeOf(&NullUint64{}):
			v := uint64(rv.Uint())
			return &NullUint64{Uint64Common{P: &v}}, true
		case rv.Kind() == reflect.Uint32 && to == reflect.TypeOf(&NullUint32{}):
			v := uint32(rv.Uint())
			return &NullUint32{Uint32Common{P: &v}}, true
		case rv.Kind() == reflect.Uint16 && to == reflect.TypeOf(&NullUint16{}):
			v := uint16(rv.Uint())
			return &NullUint16{Uint16Common{P: &v}}, true
		case rv.Kind() == reflect.Uint8 && to == reflect.TypeOf(&NullUint8{}):
			v := uint8(rv.Uint())
			return &NullUint8{Uint8Common{P: &v}}, true
		case rv.Kind() == reflect.Uint && to == reflect.TypeOf(&NotNullUint{}):
			v := uint(rv.Uint())
			return &NotNullUint{UintCommon{P: &v}}, true
		case rv.Kind() == reflect.Uint64 && to == reflect.TypeOf(&NotNullUint64{}):
			v := uint64(rv.Uint())
			return &NotNullUint64{Uint64Common{P: &v}}, true
		case rv.Kind() == reflect.Uint32 && to == reflect.TypeOf(&NotNullUint32{}):
			v := uint32(rv.Uint())
			return &NotNullUint32{Uint32Common{P: &v}}, true
		case rv.Kind() == reflect.Uint16 && to == reflect.TypeOf(&NotNullUint16{}):
			v := uint16(rv.Uint())
			return &NotNullUint16{Uint16Common{P: &v}}, true
		case rv.Kind() == reflect.Uint8 && to == reflect.TypeOf(&NotNullUint8{}):
			v := uint8(rv.Uint())
			return &NotNullUint8{Uint8Common{P: &v}}, true
		}
		return nil, false
	})
	// - to SQLValueType
	uintSqValueConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		return SQLValueType{driver.Value(int64(rv.Uint())), from}, isSafeUintToInt(rv.Uint(), 64)
	}
	matrixSuite.SetConverters(uintReflectTypes, sqlValueReflectTypes, uintSqValueConverter)
	// For other types
	matrixSuite.SetConverters(interfaceReflectTypes, uintReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		if isUint(rv.Kind()) {
			s := isSafeUint(rv.Uint(), bitSizeMap[to.Kind()])
			switch to.Kind() {
			case reflect.Uint:
				return uint(rv.Uint()), s
			case reflect.Uint64:
				return uint64(rv.Uint()), s
			case reflect.Uint32:
				return uint32(rv.Uint()), s
			case reflect.Uint16:
				return uint16(rv.Uint()), s
			case reflect.Uint8:
				return uint8(rv.Uint()), s
			}
		}
		return nil, false
	})
}

func TestUint(t *testing.T) {
	testData := matrixSuite.Generate()
	dUint := uint(magicNumber)
	dUint8 := uint8(magicNumber)
	dUint16 := uint16(magicNumber)
	dUint32 := uint32(magicNumber)
	dUint64 := uint64(magicNumber)
	for _, di := range testData {
		testOfDefault(t, di.value.Interface(), "Uint", dUint)
		testOfDefault(t, di.value.Interface(), "Uint8", dUint8)
		testOfDefault(t, di.value.Interface(), "Uint16", dUint16)
		testOfDefault(t, di.value.Interface(), "Uint32", dUint32)
		testOfDefault(t, di.value.Interface(), "Uint64", dUint64)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Uint", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Uint8", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Uint16", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Uint32", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Uint64", errPassed)
		testOfDefaultErr(t, di.value.Interface(), "Uint", dUint, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Uint8", dUint8, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Uint16", dUint16, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Uint32", dUint32, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Uint64", dUint64, ErrDefaultValue)
		switch di.value.Kind() {
		case reflect.Int64:
			testNative(t, IntUint8, []interface{}{di.value.Int(), dUint8})
			testNative(t, IntUint16, []interface{}{di.value.Int(), dUint16})
			testNative(t, IntUint32, []interface{}{di.value.Int(), dUint32})
			testNative(t, IntUint64, []interface{}{di.value.Int(), dUint64})
		case reflect.Uint64:
			testNative(t, Uint8, []interface{}{di.value.Uint(), dUint8})
			testNative(t, Uint16, []interface{}{di.value.Uint(), dUint16})
			testNative(t, Uint32, []interface{}{di.value.Uint(), dUint32})
		case reflect.Float32:
			testNative(t, Float32Uint, []interface{}{float32(di.value.Float()), dUint})
			testNative(t, Float32Uint8, []interface{}{float32(di.value.Float()), dUint8})
			testNative(t, Float32Uint16, []interface{}{float32(di.value.Float()), dUint16})
			testNative(t, Float32Uint32, []interface{}{float32(di.value.Float()), dUint32})
			testNative(t, Float32Uint64, []interface{}{float32(di.value.Float()), dUint64})
		case reflect.Float64:
			testNative(t, FloatUint, []interface{}{di.value.Float(), dUint})
			testNative(t, FloatUint8, []interface{}{di.value.Float(), dUint8})
			testNative(t, FloatUint16, []interface{}{di.value.Float(), dUint16})
			testNative(t, FloatUint32, []interface{}{di.value.Float(), dUint32})
			testNative(t, FloatUint64, []interface{}{di.value.Float(), dUint64})
		case reflect.Complex64:
			testNative(t, Complex64Uint, []interface{}{complex64(di.value.Complex()), dUint})
			testNative(t, Complex64Uint8, []interface{}{complex64(di.value.Complex()), dUint8})
			testNative(t, Complex64Uint16, []interface{}{complex64(di.value.Complex()), dUint16})
			testNative(t, Complex64Uint32, []interface{}{complex64(di.value.Complex()), dUint32})
			testNative(t, Complex64Uint64, []interface{}{complex64(di.value.Complex()), dUint64})
		case reflect.Complex128:
			testNative(t, ComplexUint, []interface{}{di.value.Complex(), dUint})
			testNative(t, ComplexUint8, []interface{}{di.value.Complex(), dUint8})
			testNative(t, ComplexUint16, []interface{}{di.value.Complex(), dUint16})
			testNative(t, ComplexUint32, []interface{}{di.value.Complex(), dUint32})
			testNative(t, ComplexUint64, []interface{}{di.value.Complex(), dUint64})
		case reflect.String:
			testNative(t, StringUint, []interface{}{di.value.String(), dUint})
			testNative(t, StringUint8, []interface{}{di.value.String(), dUint8})
			testNative(t, StringUint16, []interface{}{di.value.String(), dUint16})
			testNative(t, StringUint32, []interface{}{di.value.String(), dUint32})
			testNative(t, StringUint64, []interface{}{di.value.String(), dUint64})
		}
	}
	testOfDefaultNil(t, "Uint")
	testOfDefaultNil(t, "Uint64")
	testOfDefaultNil(t, "Uint32")
	testOfDefaultNil(t, "Uint16")
	testOfDefaultNil(t, "Uint8")
}

// TODO: Benchmark
