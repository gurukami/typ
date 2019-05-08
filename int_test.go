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
	// Converters
	// - to bool
	intBoolConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, b := reflect.ValueOf(from), false
		bp := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolPositive{}))
		bh := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolHumanize{}))
		if bp != nil || bh != nil {
			if rv.Int() > 0 {
				b = true
			}
		} else {
			if rv.Int() != 0 {
				b = true
			}
		}
		return b, true
	}
	matrixSuite.SetConverters(intReflectTypes, boolReflectTypes, intBoolConverter)
	// - to complex
	intComplexConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, c, s := reflect.ValueOf(from), complex128(0), false
		i := rv.Int()
		if s = isSafeIntToFloat(i, bitSizeMap[to.Kind()]); s {
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
	matrixSuite.SetConverters(intReflectTypes, complexReflectTypes, intComplexConverter)
	// - to float
	intFloatConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, f, s := reflect.ValueOf(from), float64(0), false
		i := rv.Int()
		if s = isSafeIntToFloat(i, bitSizeMap[to.Kind()]); s {
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
	matrixSuite.SetConverters(intReflectTypes, floatReflectTypes, intFloatConverter)
	// - to string
	intStringConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv, str := reflect.ValueOf(from), ""
		str = strconv.FormatInt(rv.Int(), 10)
		return str, true
	}
	matrixSuite.SetConverters(intReflectTypes, stringReflectTypes, intStringConverter)
	// - to &NullInt*{}, &NotNullInt*{}
	matrixSuite.SetConverters(intReflectTypes, nullIntReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Int && to == reflect.TypeOf(&NullInt{}):
			v := int(rv.Int())
			return &NullInt{IntCommon{P: &v}}, true
		case rv.Kind() == reflect.Int64 && to == reflect.TypeOf(&NullInt64{}):
			v := int64(rv.Int())
			return &NullInt64{Int64Common{P: &v}}, true
		case rv.Kind() == reflect.Int32 && to == reflect.TypeOf(&NullInt32{}):
			v := int32(rv.Int())
			return &NullInt32{Int32Common{P: &v}}, true
		case rv.Kind() == reflect.Int16 && to == reflect.TypeOf(&NullInt16{}):
			v := int16(rv.Int())
			return &NullInt16{Int16Common{P: &v}}, true
		case rv.Kind() == reflect.Int8 && to == reflect.TypeOf(&NullInt8{}):
			v := int8(rv.Int())
			return &NullInt8{Int8Common{P: &v}}, true
		case rv.Kind() == reflect.Int && to == reflect.TypeOf(&NotNullInt{}):
			v := int(rv.Int())
			return &NotNullInt{IntCommon{P: &v}}, true
		case rv.Kind() == reflect.Int64 && to == reflect.TypeOf(&NotNullInt64{}):
			v := int64(rv.Int())
			return &NotNullInt64{Int64Common{P: &v}}, true
		case rv.Kind() == reflect.Int32 && to == reflect.TypeOf(&NotNullInt32{}):
			v := int32(rv.Int())
			return &NotNullInt32{Int32Common{P: &v}}, true
		case rv.Kind() == reflect.Int16 && to == reflect.TypeOf(&NotNullInt16{}):
			v := int16(rv.Int())
			return &NotNullInt16{Int16Common{P: &v}}, true
		case rv.Kind() == reflect.Int8 && to == reflect.TypeOf(&NotNullInt8{}):
			v := int8(rv.Int())
			return &NotNullInt8{Int8Common{P: &v}}, true
		}
		return nil, false
	})
	// - to SQLValueType
	intSqValueConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		return SQLValueType{driver.Value(int64(rv.Int())), from}, true
	}
	matrixSuite.SetConverters(intReflectTypes, sqlValueReflectTypes, intSqValueConverter)
	// For other types
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
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Int", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Int8", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Int16", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Int32", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Int64", errPassed)
		testOfDefaultErr(t, di.value.Interface(), "Int", dInt, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Int8", dInt8, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Int16", dInt16, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Int32", dInt32, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Int64", dInt64, ErrDefaultValue)
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

// TODO: Benchmark
