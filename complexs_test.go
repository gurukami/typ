package typ

import (
	"database/sql/driver"
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"testing"
)

var complexReflectTypes = []reflect.Type{
	getDefaultType(reflect.Complex64),
	getDefaultType(reflect.Complex128),
}

func init() {
	// Test Data
	matrixSuite.Register(getDefaultType(reflect.Complex128), []dataItem{
		{reflect.ValueOf(complex128(0)), nil},
		{reflect.ValueOf(complex(float64(MinFloat64), float64(MinFloat64))), nil},
		{reflect.ValueOf(complex(float64(MaxFloat64), float64(MaxFloat64))), nil},
		{reflect.ValueOf(cmplx.Inf()), nil},
		{reflect.ValueOf(cmplx.NaN()), nil},
	})
	matrixSuite.Register(getDefaultType(reflect.Complex64), []dataItem{
		{reflect.ValueOf(complex64(0)), nil},
		{reflect.ValueOf(complex(float32(MinFloat32), float32(MinFloat32))), nil},
		{reflect.ValueOf(complex(float32(MaxFloat32), float32(MaxFloat32))), nil},
		{reflect.ValueOf(complex64(cmplx.Inf())), nil},
		{reflect.ValueOf(complex64(cmplx.NaN())), nil},
	})
	// Converters
	// - to bool
	complexBoolConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		c, b := rv.Complex(), false
		bp := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolPositive(true)))
		bh := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolHumanize(true)))
		if bp != nil || bh != nil {
			if real(c) > 0 {
				b = true
			}
		} else {
			if c != 0 {
				b = true
			}
		}
		return b, true
	}
	matrixSuite.SetConverters(complexReflectTypes, boolReflectTypes, complexBoolConverter)
	// - to float
	complexFloatConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		c, f, s := rv.Complex(), float64(0), false
		if s = isSafeComplexToFloat(c, bitSizeMap[to.Kind()]); s {
			f = float64(real(c))
		}
		switch to.Kind() {
		case reflect.Float32:
			return float32(f), s
		case reflect.Float64:
			return float64(f), s
		}
		return nil, false
	}
	matrixSuite.SetConverters(complexReflectTypes, floatReflectTypes, complexFloatConverter)
	// - to int
	complexIntConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		c, i, s := rv.Complex(), int64(0), false
		if s = isSafeComplexToInt(c, bitSizeMap[rv.Kind()], bitSizeMap[to.Kind()]); s {
			i = int64(real(c))
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
	matrixSuite.SetConverters(complexReflectTypes, intReflectTypes, complexIntConverter)
	// - to uint
	complexUintConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		c, i, s := rv.Complex(), uint64(0), false
		if s = isSafeComplexToUint(c, bitSizeMap[rv.Kind()], bitSizeMap[to.Kind()]); s {
			i = uint64(real(c))
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
	matrixSuite.SetConverters(complexReflectTypes, uintReflectTypes, complexUintConverter)
	// - to string
	complexStringConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return fmt.Sprintf("%v", from), true
	}
	matrixSuite.SetConverters(complexReflectTypes, stringReflectTypes, complexStringConverter)
	// - to &NullComplex*{}
	matrixSuite.SetConverters(complexReflectTypes, nullComplexReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Complex64 && to == reflect.TypeOf(&NullComplex64{}):
			v := complex64(rv.Complex())
			return &NullComplex64{P: &v}, true
		case rv.Kind() == reflect.Complex128 && to == reflect.TypeOf(&NullComplex{}):
			v := rv.Complex()
			return &NullComplex{P: &v}, true
		}
		return nil, false
	})
	// - to SqlValueType
	matrixSuite.SetConverters(complexReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		switch {
		case rv.Kind() == reflect.Complex64:
			nv := Complex64Float64(complex64(rv.Complex()))
			return SqlValueType{driver.Value(nv.V()), from}, nv.Valid()
		case rv.Kind() == reflect.Complex128:
			nv := ComplexFloat64(rv.Complex())
			return SqlValueType{driver.Value(nv.V()), from}, nv.Valid()
		}
		return nil, false
	})
	// For other types
	matrixSuite.SetConverters(interfaceReflectTypes, complexReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
	matrixSuite.SetComparators(func(a interface{}, b interface{}) bool {
		arv := reflect.ValueOf(a)
		arb := reflect.ValueOf(b)
		if arv.Kind() == arb.Kind() {
			cflags := func(c complex128) uint8 {
				var flags uint8
				if math.IsNaN(real(c)) {
					flags |= 1
				}
				if math.IsNaN(imag(c)) {
					flags |= 2
				}
				if math.IsInf(real(c), -1) {
					flags |= 4
				}
				if math.IsInf(imag(c), -1) {
					flags |= 8
				}
				if math.IsInf(real(c), 1) {
					flags |= 16
				}
				if math.IsInf(imag(c), 1) {
					flags |= 32
				}
				return flags
			}
			return cflags(arv.Complex()) == cflags(arb.Complex())
		}
		return false
	}, complexReflectTypes)
}

func TestComplex(t *testing.T) {
	testData := matrixSuite.Generate()
	dCmplx64 := complex(float32(magicNumber), float32(0))
	dCmplx128 := complex(float64(magicNumber), float64(0))
	for _, di := range testData {
		testOfDefault(t, di.value.Interface(), "Complex64", dCmplx64)
		testOfDefault(t, di.value.Interface(), "Complex", dCmplx128)
		switch di.value.Kind() {
		case reflect.Int:
			testNative(t, IntComplex64, []interface{}{di.value.Int(), dCmplx64})
			testNative(t, IntComplex, []interface{}{di.value.Int(), dCmplx128})
		case reflect.Uint:
			testNative(t, UintComplex64, []interface{}{di.value.Uint(), dCmplx64})
			testNative(t, UintComplex, []interface{}{di.value.Uint(), dCmplx128})
		case reflect.Float32:
			testNative(t, Float32Complex64, []interface{}{float32(di.value.Float()), dCmplx64})
		case reflect.Float64:
			testNative(t, FloatComplex64, []interface{}{di.value.Float(), dCmplx64})
		case reflect.String:
			testNative(t, StringComplex64, []interface{}{di.value.String(), dCmplx64})
			testNative(t, StringComplex, []interface{}{di.value.String(), dCmplx128})
		case reflect.Complex128:
			testNative(t, Complex64, []interface{}{di.value.Complex(), complex64(dCmplx64)})
		}
	}
	testOfDefaultNil(t, "Complex64")
	testOfDefaultNil(t, "Complex")
}

// BenchmarkOfComplex/StringComplex64-8         	 1000000	      1318 ns/op
// BenchmarkOfComplex/Complex64-8               	 5000000	       253 ns/op
// BenchmarkOfComplex/NativeFloat32Complex64-8  	20000000	       105 ns/op
// BenchmarkOfComplex/StringComplex-8           	 1000000	      1666 ns/op
// BenchmarkOfComplex/Complex-8                 	10000000	       158 ns/op
// BenchmarkOfComplex/NativeFloatComplex-8      	10000000	       130 ns/op
func BenchmarkOfComplex(b *testing.B) {
	c64 := "(3.4028235e+38+3.4028235e+38i)"
	c128 := "(1.7976931348623157e+308+1.7976931348623157e+308i)"
	b.Run("StringComplex64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(c64).Complex64()
		}
	})
	b.Run("Complex64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(MaxFloat32).Complex64()
		}
	})
	b.Run("NativeFloat32Complex64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Float32Complex64(MaxFloat32)
		}
	})
	b.Run("StringComplex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(c128).Complex()
		}
	})
	b.Run("Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Of(MaxFloat64).Complex()
		}
	})
	b.Run("NativeFloatComplex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FloatComplex64(MaxFloat64)
		}
	})
}
