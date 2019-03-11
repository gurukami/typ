package typ

import (
	"database/sql/driver"
	"reflect"
	"strings"
	"testing"
)

type BoolPositive bool
type BoolHumanize bool

var boolReflectTypes = []reflect.Type{
	getDefaultType(reflect.Bool),
}

func init() {
	// Test Data
	matrixSuite.Register(getDefaultType(reflect.Bool), []dataItem{
		{reflect.ValueOf(true), nil},
		{reflect.ValueOf(false), nil},
	})
	// Converters
	// - to int, uint
	boolIntUintConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		b, i := rv.Bool(), 0
		if b {
			i = 1
		}
		switch to.Kind() {
		case reflect.Int:
			return int(i), true
		case reflect.Int64:
			return int64(i), true
		case reflect.Int32:
			return int32(i), true
		case reflect.Int16:
			return int16(i), true
		case reflect.Int8:
			return int8(i), true
		case reflect.Uint:
			return uint(i), true
		case reflect.Uint64:
			return uint64(i), true
		case reflect.Uint32:
			return uint32(i), true
		case reflect.Uint16:
			return uint16(i), true
		case reflect.Uint8:
			return uint8(i), true
		}
		return nil, false
	}
	matrixSuite.SetConverters(boolReflectTypes, intReflectTypes, boolIntUintConverter)
	matrixSuite.SetConverters(boolReflectTypes, uintReflectTypes, boolIntUintConverter)
	// - to complex
	boolComplexConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		b, c := rv.Bool(), complex128(0)
		if b {
			c = complex128(1)
		}
		switch to.Kind() {
		case reflect.Complex128:
			return complex128(c), true
		case reflect.Complex64:
			return complex64(c), true
		}
		return nil, false
	}
	matrixSuite.SetConverters(boolReflectTypes, complexReflectTypes, boolComplexConverter)
	// - to float
	boolFloatConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		b, f := rv.Bool(), float64(0)
		if b {
			f = float64(1)
		}
		switch to.Kind() {
		case reflect.Float64:
			return float64(f), true
		case reflect.Float32:
			return float32(f), true
		}
		return nil, false
	}
	matrixSuite.SetConverters(boolReflectTypes, floatReflectTypes, boolFloatConverter)
	// - to string
	matrixSuite.SetConverter(getDefaultType(reflect.Bool), getDefaultType(reflect.String), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		b, s := rv.Bool(), "0"
		if b {
			s = "1"
		}
		return s, true
	})
	// - to &NullBool{}
	matrixSuite.SetConverter(getDefaultType(reflect.Bool), reflect.TypeOf(&NullBool{}), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		v := from.(bool)
		nb := &NullBool{P: &v}
		return nb, true
	})
	// - to SQLValueType
	matrixSuite.SetConverter(getDefaultType(reflect.Bool), reflect.TypeOf(SQLValueType{}), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {

		return SQLValueType{driver.Value(from), from}, true
	})
	// For other types
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Bool), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestBool(t *testing.T) {
	testData := matrixSuite.Generate()
	for _, v := range testData {
		expectedValue, expectedValid, _ := matrixSuite.Test(v.value.Interface(), getDefaultType(reflect.Bool))
		if expectedValue == nil {
			continue
		}
		nv := Of(v.value.Interface()).Bool()
		if !matrixSuite.Compare(nv.V(), expectedValue) || nv.Valid() != expectedValid {
			t.Errorf("Of(%T(%+[1]v)).Bool(), %s", v.value.Interface(), errNull{
				expectedValue, expectedValid, nil,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestBoolPositive(t *testing.T) {
	testData := matrixSuite.Generate()
	for _, v := range testData {
		expectedValue, expectedValid, _ := matrixSuite.Test(v.value.Interface(), getDefaultType(reflect.Bool), BoolPositive(true))
		if expectedValue == nil {
			continue
		}
		nv := Of(v.value.Interface()).BoolPositive()
		if !matrixSuite.Compare(nv.V(), expectedValue) || nv.Valid() != expectedValid {
			t.Errorf("Of(%T(%+[1]v)).BoolPositive(), %s", v.value.Interface(), errNull{
				expectedValue, expectedValid, nil,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestBoolHumanize(t *testing.T) {
	testData := matrixSuite.Generate()
	for _, di := range testData {
		eValue, expectedValid, _ := matrixSuite.Test(di.value.Interface(), getDefaultType(reflect.Bool), BoolHumanize(true))
		if eValue == nil {
			continue
		}
		nv := Of(di.value.Interface()).BoolHumanize()
		if !matrixSuite.Compare(nv.V(), eValue) || nv.Valid() != expectedValid {
			t.Errorf("Of(%T(%+[1]v)).BoolHumanize(), %s", di.value.Interface(), errNull{
				eValue, expectedValid, nil,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
		if stringNv, ok := di.value.Interface().(string); ok {
			nv := StringBoolHumanize(stringNv)
			if !matrixSuite.Compare(nv.V(), eValue) || nv.Valid() != expectedValid {
				t.Errorf("StringBoolHumanize(%T(%+[1]v)), %s", stringNv, errNull{
					eValue, expectedValid, nil,
					nv.V(), nv.Valid(), nv.Error,
				})
			}
		}
	}
}

// BenchmarkOfBool-8   	 10000000	       124 ns/op
func BenchmarkOfBool(b *testing.B) {
	v := map[interface{}]interface{}{}
	for i := 0; i < b.N; i++ {
		Of(v).Bool()
	}
}

// BenchmarkOfBoolPositive-8   	 20000000	       98.4 ns/op
func BenchmarkOfBoolPositive(b *testing.B) {
	v := map[interface{}]interface{}{}
	for i := 0; i < b.N; i++ {
		Of(v).BoolPositive()
	}
}

// BenchmarkOfBoolHumanize-8   	 10000000	       160 ns/op
func BenchmarkOfBoolHumanize(b *testing.B) {
	v := strings.Repeat("a", 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Of(v).BoolHumanize()
	}
}

// BenchmarkNativeStringBoolHumanize-8   	 30000000	       38.3 ns/op
func BenchmarkNativeStringBoolHumanize(b *testing.B) {
	v := strings.Repeat("a", 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringBoolHumanize(v)
	}
}
