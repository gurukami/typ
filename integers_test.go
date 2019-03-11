package typ

import (
	"database/sql/driver"
	"reflect"
	"strconv"
	"testing"
)

var intReflectTypes = []reflect.Type{
	getDefaultType(reflect.Int),
	getDefaultType(reflect.Int64),
	getDefaultType(reflect.Int32),
	getDefaultType(reflect.Int16),
	getDefaultType(reflect.Int8),
}

var uintReflectTypes = []reflect.Type{
	getDefaultType(reflect.Uint),
	getDefaultType(reflect.Uint64),
	getDefaultType(reflect.Uint32),
	getDefaultType(reflect.Uint16),
	getDefaultType(reflect.Uint8),
}

func init() {
	// Test Data
	matrixSuite.Register(getDefaultType(reflect.Int), []dataItem{
		{reflect.ValueOf(MaxInt), nil},
		{reflect.ValueOf(MinInt), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Int64), []dataItem{
		{reflect.ValueOf(MaxInt64), nil},
		{reflect.ValueOf(MinInt64), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Int32), []dataItem{
		{reflect.ValueOf(MaxInt32), nil},
		{reflect.ValueOf(MinInt32), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Int16), []dataItem{
		{reflect.ValueOf(MaxInt16), nil},
		{reflect.ValueOf(MinInt16), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Int8), []dataItem{
		{reflect.ValueOf(MaxInt8), nil},
		{reflect.ValueOf(MinInt8), nil},
	})
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
	intUintBoolConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, b := reflect.ValueOf(from), false
		switch {
		case isUint(rv.Kind()):
			if rv.Uint() != 0 {
				b = true
			}
		case isInt(rv.Kind()):
			bp := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolPositive(true)))
			bh := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolHumanize(true)))
			if bp != nil || bh != nil {
				if rv.Int() > 0 {
					b = true
				}
			} else {
				if rv.Int() != 0 {
					b = true
				}
			}
		}
		return b, true
	}
	matrixSuite.SetConverters(uintReflectTypes, boolReflectTypes, intUintBoolConverter)
	matrixSuite.SetConverters(intReflectTypes, boolReflectTypes, intUintBoolConverter)
	// - to complex
	intUintComplexConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, c, s := reflect.ValueOf(from), complex128(0), false
		switch {
		case isInt(rv.Kind()):
			i := rv.Int()
			if s = isSafeIntToFloat(i, bitSizeMap[to.Kind()]); s {
				c = complex(float64(i), 0)
			}
		case isUint(rv.Kind()):
			i := rv.Uint()
			if s = isSafeUintToFloat(i, bitSizeMap[to.Kind()]); s {
				c = complex(float64(i), 0)
			}
		}
		switch to.Kind() {
		case reflect.Complex64:
			return complex64(c), s
		case reflect.Complex128:
			return complex128(c), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(uintReflectTypes, complexReflectTypes, intUintComplexConverter)
	matrixSuite.SetConverters(intReflectTypes, complexReflectTypes, intUintComplexConverter)
	// - to float
	intUintFloatConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, f, s := reflect.ValueOf(from), float64(0), false
		switch {
		case isInt(rv.Kind()):
			i := rv.Int()
			if s = isSafeIntToFloat(i, bitSizeMap[to.Kind()]); s {
				f = float64(i)
			}
		case isUint(rv.Kind()):
			i := rv.Uint()
			if s = isSafeUintToFloat(i, bitSizeMap[to.Kind()]); s {
				f = float64(i)
			}
		}
		switch to.Kind() {
		case reflect.Float32:
			return float32(f), s
		case reflect.Float64:
			return float64(f), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(uintReflectTypes, floatReflectTypes, intUintFloatConverter)
	matrixSuite.SetConverters(intReflectTypes, floatReflectTypes, intUintFloatConverter)
	// - to string
	intUintStringConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, str := reflect.ValueOf(from), ""
		switch {
		case isInt(rv.Kind()):
			str = strconv.FormatInt(rv.Int(), 10)
		case isUint(rv.Kind()):
			str = strconv.FormatUint(rv.Uint(), 10)
		}
		return str, true
	}
	matrixSuite.SetConverters(uintReflectTypes, stringReflectTypes, intUintStringConverter)
	matrixSuite.SetConverters(intReflectTypes, stringReflectTypes, intUintStringConverter)
	// - to &NullInt*{}
	matrixSuite.SetConverters(intReflectTypes, nullIntReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Int && to == reflect.TypeOf(&NullInt{}):
			v := int(rv.Int())
			return &NullInt{P: &v}, true
		case rv.Kind() == reflect.Int64 && to == reflect.TypeOf(&NullInt64{}):
			v := int64(rv.Int())
			return &NullInt64{P: &v}, true
		case rv.Kind() == reflect.Int32 && to == reflect.TypeOf(&NullInt32{}):
			v := int32(rv.Int())
			return &NullInt32{P: &v}, true
		case rv.Kind() == reflect.Int16 && to == reflect.TypeOf(&NullInt16{}):
			v := int16(rv.Int())
			return &NullInt16{P: &v}, true
		case rv.Kind() == reflect.Int8 && to == reflect.TypeOf(&NullInt8{}):
			v := int8(rv.Int())
			return &NullInt8{P: &v}, true
		}
		return nil, false
	})
	// - to &NullUint*{}
	matrixSuite.SetConverters(uintReflectTypes, nullUintReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Uint && to == reflect.TypeOf(&NullUint{}):
			v := uint(rv.Uint())
			return &NullUint{P: &v}, true
		case rv.Kind() == reflect.Uint64 && to == reflect.TypeOf(&NullUint64{}):
			v := uint64(rv.Uint())
			return &NullUint64{P: &v}, true
		case rv.Kind() == reflect.Uint32 && to == reflect.TypeOf(&NullUint32{}):
			v := uint32(rv.Uint())
			return &NullUint32{P: &v}, true
		case rv.Kind() == reflect.Uint16 && to == reflect.TypeOf(&NullUint16{}):
			v := uint16(rv.Uint())
			return &NullUint16{P: &v}, true
		case rv.Kind() == reflect.Uint8 && to == reflect.TypeOf(&NullUint8{}):
			v := uint8(rv.Uint())
			return &NullUint8{P: &v}, true
		}
		return nil, false
	})
	// - to SQLValueType
	intUintSqValueConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case isInt(rv.Kind()):
			return SQLValueType{driver.Value(int64(rv.Int())), from}, true
		case isUint(rv.Kind()):
			return SQLValueType{driver.Value(int64(rv.Uint())), from}, isSafeUintToInt(rv.Uint(), 64)
		}
		return nil, false
	}
	matrixSuite.SetConverters(intReflectTypes, sqlValueReflectTypes, intUintSqValueConverter)
	matrixSuite.SetConverters(uintReflectTypes, sqlValueReflectTypes, intUintSqValueConverter)
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
	matrixSuite.SetConverters(interfaceReflectTypes, intReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		if isInt(rv.Kind()) {
			s := isSafeInt(rv.Int(), bitSizeMap[to.Kind()])
			switch to.Kind() {
			case reflect.Int:
				return int(rv.Int()), s
			case reflect.Int64:
				return int64(rv.Int()), s
			case reflect.Int32:
				return int32(rv.Int()), s
			case reflect.Int16:
				return int16(rv.Int()), s
			case reflect.Int8:
				return int8(rv.Int()), s
			}
		}
		return nil, false
	})
}

func TestInt(t *testing.T) {
	testData := matrixSuite.Generate()
	dInt := int(magicNumber)
	dInt8 := int8(magicNumber)
	dInt16 := int16(magicNumber)
	dInt32 := int32(magicNumber)
	dInt64 := int64(magicNumber)
	for _, di := range testData {
		if di.value.Kind() == reflect.Complex128 {
			testNative(t, ComplexInt32, []interface{}{di.value.Complex(), dInt32})
		}
		testOfDefault(t, di.value.Interface(), "Int", dInt)
		testOfDefault(t, di.value.Interface(), "Int8", dInt8)
		testOfDefault(t, di.value.Interface(), "Int16", dInt16)
		testOfDefault(t, di.value.Interface(), "Int32", dInt32)
		testOfDefault(t, di.value.Interface(), "Int64", dInt64)
		switch di.value.Kind() {
		case reflect.Int64:
			testNative(t, Int8, []interface{}{di.value.Int(), dInt8})
			testNative(t, Int16, []interface{}{di.value.Int(), dInt16})
			testNative(t, Int32, []interface{}{di.value.Int(), dInt32})
		case reflect.Uint64:
			testNative(t, UintInt8, []interface{}{di.value.Uint(), dInt8})
			testNative(t, UintInt16, []interface{}{di.value.Uint(), dInt16})
			testNative(t, UintInt32, []interface{}{di.value.Uint(), dInt32})
			testNative(t, UintInt64, []interface{}{di.value.Uint(), dInt64})
		case reflect.Float32:
			testNative(t, Float32Int, []interface{}{float32(di.value.Float()), dInt})
			testNative(t, Float32Int8, []interface{}{float32(di.value.Float()), dInt8})
			testNative(t, Float32Int16, []interface{}{float32(di.value.Float()), dInt16})
			testNative(t, Float32Int32, []interface{}{float32(di.value.Float()), dInt32})
			testNative(t, Float32Int64, []interface{}{float32(di.value.Float()), dInt64})
		case reflect.Float64:
			testNative(t, FloatInt, []interface{}{di.value.Float(), dInt})
			testNative(t, FloatInt8, []interface{}{di.value.Float(), dInt8})
			testNative(t, FloatInt16, []interface{}{di.value.Float(), dInt16})
			testNative(t, FloatInt32, []interface{}{di.value.Float(), dInt32})
			testNative(t, FloatInt64, []interface{}{di.value.Float(), dInt64})
		case reflect.Complex64:
			testNative(t, Complex64Int, []interface{}{complex64(di.value.Complex()), dInt})
			testNative(t, Complex64Int8, []interface{}{complex64(di.value.Complex()), dInt8})
			testNative(t, Complex64Int16, []interface{}{complex64(di.value.Complex()), dInt16})
			testNative(t, Complex64Int32, []interface{}{complex64(di.value.Complex()), dInt32})
			testNative(t, Complex64Int64, []interface{}{complex64(di.value.Complex()), dInt64})
		case reflect.Complex128:
			testNative(t, ComplexInt, []interface{}{di.value.Complex(), dInt})
			testNative(t, ComplexInt8, []interface{}{di.value.Complex(), dInt8})
			testNative(t, ComplexInt16, []interface{}{di.value.Complex(), dInt16})
			testNative(t, ComplexInt32, []interface{}{di.value.Complex(), dInt32})
			testNative(t, ComplexInt64, []interface{}{di.value.Complex(), dInt64})
		case reflect.String:
			testNative(t, StringInt, []interface{}{di.value.String(), dInt})
			testNative(t, StringInt8, []interface{}{di.value.String(), dInt8})
			testNative(t, StringInt16, []interface{}{di.value.String(), dInt16})
			testNative(t, StringInt32, []interface{}{di.value.String(), dInt32})
			testNative(t, StringInt64, []interface{}{di.value.String(), dInt64})
		}
	}
	testOfDefaultNil(t, "Int")
	testOfDefaultNil(t, "Int64")
	testOfDefaultNil(t, "Int32")
	testOfDefaultNil(t, "Int16")
	testOfDefaultNil(t, "Int8")
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
