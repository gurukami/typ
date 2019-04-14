package typ

import (
	"database/sql/driver"
	"math"
	"reflect"
	"strconv"
	"testing"
)

var floatReflectTypes = []reflect.Type{
	getDefaultType(reflect.Float32),
	getDefaultType(reflect.Float64),
}

func init() {
	// Test Data
	matrixSuite.Register(getDefaultType(reflect.Float64), []dataItem{
		{reflect.ValueOf(float64(0)), nil},
		{reflect.ValueOf(MinFloat64), nil},
		{reflect.ValueOf(MaxFloat64), nil},
		{reflect.ValueOf(MinSafeIntFloat64), nil},
		{reflect.ValueOf(MaxSafeIntFloat64), nil},
		{reflect.ValueOf(math.SmallestNonzeroFloat64), nil},
		{reflect.ValueOf(math.Inf(0)), nil},
		{reflect.ValueOf(math.NaN()), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Float32), []dataItem{
		{reflect.ValueOf(float32(0)), nil},
		{reflect.ValueOf(MinFloat32), nil},
		{reflect.ValueOf(MaxFloat32), nil},
		{reflect.ValueOf(MinSafeIntFloat32), nil},
		{reflect.ValueOf(MaxSafeIntFloat32), nil},
		{reflect.ValueOf(float32(math.SmallestNonzeroFloat32)), nil},
		{reflect.ValueOf(float32(math.Inf(0))), nil},
		{reflect.ValueOf(float32(math.NaN())), nil},
	})
	// Converters
	// - to bool
	floatBoolConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		f, b := rv.Float(), false
		bp := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolPositive{}))
		bh := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolHumanize{}))
		if bp != nil || bh != nil {
			if f > 0 {
				b = true
			}
		} else {
			if f != 0 {
				b = true
			}
		}
		return b, true
	}
	matrixSuite.SetConverters(floatReflectTypes, boolReflectTypes, floatBoolConverter)
	// - to int
	floatIntConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		f, i := rv.Float(), int64(0)
		var s bool
		if s = isSafeFloatToInt(f, bitSizeMap[rv.Kind()], bitSizeMap[to.Kind()]); s {
			i = int64(f)
		}
		switch to.Kind() {
		case reflect.Int:
			return int(i), s
		case reflect.Int64:
			return int64(i), s
		case reflect.Int32:
			return int32(i), s
		case reflect.Int16:
			return int16(i), s
		case reflect.Int8:
			return int8(i), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(floatReflectTypes, intReflectTypes, floatIntConverter)
	// - to uint
	floatUintConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		f, i := rv.Float(), uint64(0)
		var s bool
		if s = isSafeFloatToUint(f, bitSizeMap[rv.Kind()], bitSizeMap[to.Kind()]); s {
			i = uint64(f)
		}
		switch to.Kind() {
		case reflect.Uint:
			return uint(i), s
		case reflect.Uint64:
			return uint64(i), s
		case reflect.Uint32:
			return uint32(i), s
		case reflect.Uint16:
			return uint16(i), s
		case reflect.Uint8:
			return uint8(i), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(floatReflectTypes, uintReflectTypes, floatUintConverter)
	// - to complex
	floatComplexConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		f, c := rv.Float(), complex128(0)
		var s bool
		if s = isSafeFloat(f, bitSizeMap[to.Kind()]); s {
			c = complex(f, 0)
		}
		switch to.Kind() {
		case reflect.Complex64:
			return complex64(c), s
		case reflect.Complex128:
			return complex128(c), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(floatReflectTypes, complexReflectTypes, floatComplexConverter)
	// - to string
	floatStringConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		return strconv.FormatFloat(rv.Float(), 'g', -1, bitSizeMap[rv.Kind()]), true
	}
	matrixSuite.SetConverters(floatReflectTypes, stringReflectTypes, floatStringConverter)
	// - to &NullFloat*{}
	matrixSuite.SetConverters(floatReflectTypes, nullFloatReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Float32 && to == reflect.TypeOf(&NullFloat32{}):
			var err error
			if valid := isSafeFloat(rv.Float(), 32); !valid {
				err = ErrConvert
			}
			v := float32(rv.Float())
			return &NullFloat32{P: &v, Error: err}, true
		case rv.Kind() == reflect.Float64 && to == reflect.TypeOf(&NullFloat{}):
			v := rv.Float()
			return &NullFloat{P: &v}, true
		}
		return nil, false
	})
	// - to SQLValueType
	matrixSuite.SetConverters(floatReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Float32:
			return SQLValueType{driver.Value(rv.Float()), from}, true
		case rv.Kind() == reflect.Float64:
			return SQLValueType{driver.Value(rv.Float()), from}, true
		}
		return nil, false
	})
	// For other types
	matrixSuite.SetConverters(interfaceReflectTypes, floatReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
	matrixSuite.SetComparators(func(a interface{}, b interface{}) bool {
		arv := reflect.ValueOf(a)
		arb := reflect.ValueOf(b)
		if arv.Kind() == arb.Kind() {
			cflags := func(f float64) uint8 {
				var flags uint8
				if math.IsNaN(f) {
					flags |= 1
				}
				if math.IsInf(f, -1) {
					flags |= 2
				}
				if math.IsInf(f, 1) {
					flags |= 4
				}
				return flags
			}
			return cflags(arv.Float()) == cflags(arb.Float())
		}
		return false
	}, floatReflectTypes)
}

func TestFloat(t *testing.T) {
	testData := matrixSuite.Generate()
	dFloat32 := float32(magicNumber)
	dFloat64 := float64(magicNumber)
	for _, di := range testData {
		testOfDefault(t, di.value.Interface(), "Float32", dFloat32)
		testOfDefault(t, di.value.Interface(), "Float", dFloat64)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Float32", errPassed)
		testOfPassedErr(t, NewType(di.value.Interface(), errPassed), "Float", errPassed)
		testOfDefaultErr(t, di.value.Interface(), "Float32", dFloat32, ErrDefaultValue)
		testOfDefaultErr(t, di.value.Interface(), "Float", dFloat64, ErrDefaultValue)
		switch di.value.Kind() {
		case reflect.Int:
			testNative(t, IntFloat32, []interface{}{di.value.Int(), dFloat32})
			testNative(t, IntFloat, []interface{}{di.value.Int(), dFloat64})
		case reflect.Uint:
			testNative(t, UintFloat32, []interface{}{di.value.Uint(), dFloat32})
			testNative(t, UintFloat, []interface{}{di.value.Uint(), dFloat64})
		case reflect.Float64:
			testNative(t, Float32, []interface{}{di.value.Float(), dFloat32})
		case reflect.Complex64:
			testNative(t, Complex64Float32, []interface{}{complex64(di.value.Complex()), dFloat32})
			testNative(t, Complex64Float64, []interface{}{complex64(di.value.Complex()), dFloat64})
		case reflect.Complex128:
			testNative(t, ComplexFloat32, []interface{}{di.value.Complex(), dFloat32})
			testNative(t, ComplexFloat64, []interface{}{di.value.Complex(), dFloat64})
		case reflect.String:
			testNative(t, StringFloat32, []interface{}{di.value.String(), dFloat32})
			testNative(t, StringFloat, []interface{}{di.value.String(), dFloat64})
		}
	}
	testOfDefaultNil(t, "Float32")
	testOfDefaultNil(t, "Float")
}

// BenchmarkOfFloat/StringFloat32-8         	 5000000	       262 ns/op
// BenchmarkOfFloat/Float32-8               	10000000	       145 ns/op
// BenchmarkOfFloat/NativeStringFloat32-8   	20000000	       107 ns/op
// BenchmarkOfFloat/NativeIntFloat32-8      	50000000	        27.1 ns/op
// BenchmarkOfFloat/StringFloat-8           	 5000000	       277 ns/op
// BenchmarkOfFloat/Float-8                 	10000000	       139 ns/op
// BenchmarkOfFloat/NativeStringFloat-8     	10000000	       126 ns/op
// BenchmarkOfFloat/NativeIntFloat-8        	50000000	        31.9 ns/op
func BenchmarkOfFloat(b *testing.B) {
	c32 := "3.4028235e+38"
	c64 := "1.7976931348623157e+308"
	i16 := int64(MaxInt16)
	i32 := int64(MaxInt32)
	b.Run("StringFloat32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(c32).Float32()
		}
	})
	b.Run("Float32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(MaxInt16).Float32()
		}
	})
	b.Run("NativeStringFloat32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StringFloat32(c32)
		}
	})
	b.Run("NativeIntFloat32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IntFloat32(i16)
		}
	})
	b.Run("StringFloat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(c64).Float()
		}
	})
	b.Run("Float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(MaxInt32).Float()
		}
	})
	b.Run("NativeStringFloat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StringFloat(c64)
		}
	})
	b.Run("NativeIntFloat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IntFloat(i32)
		}
	})
}
