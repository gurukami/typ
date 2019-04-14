package typ

import (
	"database/sql/driver"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var stringReflectTypes = []reflect.Type{
	getDefaultType(reflect.String),
}

func init() {

	// Test Data

	matrixSuite.Register(getDefaultType(reflect.String), []dataItem{
		{reflect.ValueOf(""), nil},
		{reflect.ValueOf("string"), nil},
		{reflect.ValueOf("string_bytes"), []interface{}{NullStringBytes{}}},
		{reflect.ValueOf("true"), nil},
		{reflect.ValueOf("True"), nil},
		{reflect.ValueOf("false"), nil},
		{reflect.ValueOf("False"), nil},
		{reflect.ValueOf("(" + overFlowValue + "+" + overFlowValue + "i)"), nil},
		{reflect.ValueOf("(0+" + overFlowValue + "i)"), nil},
	})

	// Converters

	// - to bool

	matrixSuite.SetConverter(getDefaultType(reflect.String), getDefaultType(reflect.Bool), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		s, b := rv.String(), false
		if opt := matrixSuite.GetOptByType(opts, reflect.TypeOf(BoolHumanize{})); opt != nil {
			// false for string 'false' in case-insensitive mode or string equals '0'
			bf := strings.EqualFold("false", s) || s == "0"
			bt := strings.EqualFold("true", s) || s == "1"
			if bt {
				b = true
			}
			return b, bt || bf
		}
		if len(s) > 0 {
			b = true
		}
		return b, true
	})

	// - to int

	stringIntConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		str, i, s := rv.String(), int64(0), false
		if pi, err := strconv.ParseInt(str, 0, bitSizeMap[to.Kind()]); err == nil {
			i, s = pi, true
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

	matrixSuite.SetConverters(stringReflectTypes, intReflectTypes, stringIntConverter)

	// - to uint

	stringUintConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		str, i, s := rv.String(), uint64(0), false
		if pi, err := strconv.ParseUint(str, 0, bitSizeMap[to.Kind()]); err == nil {
			i, s = pi, true
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

	matrixSuite.SetConverters(stringReflectTypes, uintReflectTypes, stringUintConverter)

	// - to float

	stringFloatConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		str, f, s := rv.String(), float64(0), false
		if pf, err := strconv.ParseFloat(str, bitSizeMap[to.Kind()]); err == nil {
			f, s = pf, true
		}
		switch to.Kind() {
		case reflect.Float32:
			return float32(f), s
		case reflect.Float64:
			return float64(f), s
		}
		return nil, false
	}

	matrixSuite.SetConverters(stringReflectTypes, floatReflectTypes, stringFloatConverter)

	// - to complex

	stringComplexConverter := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		str, c, s := rv.String(), complex128(0), false
		matches := regexpComplex.FindStringSubmatch(str)
		if len(matches) != 0 {
			fr, re := strconv.ParseFloat(matches[1], bitSizeMap[to.Kind()])
			fi, ie := strconv.ParseFloat(matches[2], bitSizeMap[to.Kind()])
			s = re == nil && ie == nil
			if s {
				c = complex(fr, fi)
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

	matrixSuite.SetConverters(stringReflectTypes, complexReflectTypes, stringComplexConverter)

	// - to &NullString{}

	matrixSuite.SetConverters(stringReflectTypes, nullStringReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		v := from.(string)
		nb := &NullString{P: &v}
		return nb, true
	})

	// - to SQLValueType

	matrixSuite.SetConverter(getDefaultType(reflect.String), reflect.TypeOf(SQLValueType{}), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		nsb := matrixSuite.GetOptByType(opts, reflect.TypeOf(NullStringBytes{}))
		if nsb != nil {
			return SQLValueType{driver.Value([]byte(from.(string))), from}, true
		}
		return SQLValueType{driver.Value(from), from}, true
	})
}

func TestTransition(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{true, false},
			[]interface{}{"Prefix:true", "Prefix:false"},
			[]Option{Prefix("Prefix:")},
		},
		{
			[]interface{}{true, false},
			[]interface{}{"true:Suffix", "false:Suffix"},
			[]Option{Suffix(":Suffix")},
		},
	}
	var expectedValue string
	for _, v := range testData {
		for i, iV := range v[0].([]interface{}) {
			expectedValue = v[1].([]interface{})[i].(string)
			options := v[2].([]Option)
			value := Of(iV, options...).String()
			if value.V() != expectedValue {
				t.Errorf("Of(%v).String(), %s", iV, errNull{
					expectedValue, false, nil,
					value.V(), false, nil,
				})
			}
		}
	}
}

func TestConcat(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{"Hello", "World", "!"},
			"HelloWorld!",
			[]Option{},
		},
		{
			[]interface{}{"Hello", "World", "!"},
			"Hello;World;!",
			[]Option{Delimiter(";")},
		},
		{
			[]interface{}{float32(MaxFloat32), "someString", int(1)},
			"3.4028235e+38someString1",
			[]Option{},
		},
	}
	for _, v := range testData {
		expectedValue := v[1].(string)
		options := v[2].([]Option)
		if s := Concat(v[0].([]interface{}), options...); s.V() != expectedValue {
			t.Errorf("Concat(%v), %s", v[0].([]interface{}), errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with option error
		if s := Concat(v[0].([]interface{}), Base(99)); s.V() != "" || s.Error != ErrBaseInvalid {
			t.Errorf("Concat(%v), %s", v[0].([]interface{}), errNull{
				"", false, ErrBaseInvalid,
				s.V(), false, s.Error,
			})
		}
		// with convert error
		vls := []interface{}{nil}
		if s := Concat([]interface{}{nil}, options...); s.Error == nil {
			t.Errorf("Concat(%v), %s", vls, errNull{
				"", false, ErrUnexpectedValue,
				s.V(), false, s.Error,
			})
		}
	}
}

func TestStringEmpty(t *testing.T) {
	testData := [][]interface{}{
		{
			0,
			"",
		},
		{
			MaxInt,
			"9223372036854775807",
		},
	}
	for _, v := range testData {
		expectedValue := v[1].(string)
		if s := Of(v[0]).StringEmpty(); s.V() != expectedValue {
			t.Errorf("Of(%v).StringEmpty(), %s", v[0], errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with error
		nv := NewType(v[0], errPassed).StringEmpty()
		if nv.Error != errPassed || nv.Valid() {
			t.Errorf("Of(%T(%+[1]v)).StringEmpty(), %s", v[0], errNull{
				nil, false, errPassed,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestStringDefault(t *testing.T) {
	testData := [][]interface{}{
		{
			0,
			"default",
			"default",
		},
		{
			MaxInt,
			"default",
			"9223372036854775807",
		},
	}
	for _, v := range testData {
		defaultValue := v[1].(string)
		expectedValue := v[2].(string)
		if s := Of(v[0]).StringDefault(defaultValue); s.V() != expectedValue {
			t.Errorf("Of(%v).StringDefault(%v), %s", v[0], defaultValue, errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with error
		nv := NewType(v[0], errPassed).StringDefault(defaultValue)
		if nv.Error != errPassed || nv.Valid() {
			t.Errorf("Of(%T(%+[1]v)).StringDefault(%v), %s", v[0], defaultValue, errNull{
				nil, false, errPassed,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestStringInt(t *testing.T) {
	testData := [][]interface{}{
		{
			"",
			"0",
			true,
		},
		{
			MaxInt,
			"9223372036854775807",
			true,
		},
		{
			"92233720368547758079223372036854775807",
			"",
			false,
		},
		{
			float32(MaxInt16),
			"32767",
			true,
		},
	}
	for _, v := range testData {
		s := Of(v[0]).StringInt()
		expectedValue := v[1].(string)
		failed := !v[2].(bool)
		if s.V() != expectedValue && ((s.Error == nil) != failed) {
			t.Errorf("Of(%v).StringInt(), %s", v[0], errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with error
		nv := NewType(v[0], errPassed).StringInt()
		if nv.Error != errPassed || nv.Valid() {
			t.Errorf("Of(%T(%+[1]v)).StringInt(), %s", v[0], errNull{
				nil, false, errPassed,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestStringBool(t *testing.T) {
	testData := [][]interface{}{
		{
			"",
			"false",
		},
		{
			MaxInt,
			"true",
		},
		{
			"92233720368547758079223372036854775807",
			"true",
		},
		{
			float32(MaxInt16),
			"true",
		},
	}
	for _, v := range testData {
		expectedValue := v[1].(string)
		if s := Of(v[0]).StringBool(); s.V() != expectedValue {
			t.Errorf("Of(%v).StringBool(), %s", v[0], errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with error
		nv := NewType(v[0], errPassed).StringBool()
		if nv.Error != errPassed || nv.Valid() {
			t.Errorf("Of(%T(%+[1]v)).StringBool(), %s", v[0], errNull{
				nil, false, errPassed,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestStringFloat(t *testing.T) {
	testData := [][]interface{}{
		{
			"",
			"0e+00",
			true,
		},
		{
			float32(MaxInt16),
			"3.2767e+04",
			true,
		},
	}
	for _, v := range testData {
		s := Of(v[0]).StringFloat()
		expectedValue := v[1].(string)
		failed := !v[2].(bool)
		if s.V() != expectedValue && ((s.Error == nil) != failed) {
			t.Errorf("Of(%v).StringFloat(), %s", v[0], errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with error
		nv := NewType(v[0], errPassed).StringFloat()
		if nv.Error != errPassed || nv.Valid() {
			t.Errorf("Of(%T(%+[1]v)).StringFloat(), %s", v[0], errNull{
				nil, false, errPassed,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestStringComplex(t *testing.T) {
	testData := [][]interface{}{
		{
			"",
			"(0+0i)",
			true,
		},
		{
			float32(MaxInt16),
			"(32767+0i)",
			true,
		},
	}
	for _, v := range testData {
		s := Of(v[0]).StringComplex()
		expectedValue := v[1].(string)
		failed := !v[2].(bool)
		if s.V() != expectedValue && ((s.Error == nil) != failed) {
			t.Errorf("Of(%v).StringComplex(), %s", v[0], errNull{
				expectedValue, false, nil,
				s.V(), false, nil,
			})
		}
		// with error
		nv := NewType(v[0], errPassed).StringComplex()
		if nv.Error != errPassed || nv.Valid() {
			t.Errorf("Of(%T(%+[1]v)).StringComplex(), %s", v[0], errNull{
				nil, false, errPassed,
				nv.V(), nv.Valid(), nv.Error,
			})
		}
	}
}

func TestBoolToString(t *testing.T) {
	testData := [][]interface{}{
		{
			true,
			"true",
		},
		{
			false,
			"false",
		},
	}
	for _, v := range testData {
		expectedValue := v[1].(string)
		if s := BoolString(v[0].(bool)); s.V() != expectedValue {
			t.Errorf("BoolHumanize(%v) = %v failed, expectedValue %v", v[0], s, expectedValue)
		}
	}
}

func TestIntString(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{MinInt64},
			"-9223372036854775808",
		},
		{
			[]interface{}{MaxInt64},
			"9223372036854775807",
		},
		{
			[]interface{}{int64(0)},
			"0",
		},
		{
			[]interface{}{int64(0), IntStringBase(10), IntStringDefault("default")},
			"default",
		},
		{
			[]interface{}{int64(3), IntStringBase(2)},
			"11",
		},
	}
	for _, v := range testData {
		args := v[0].([]interface{})
		expectedValue := v[1].(string)
		res := rFnCall(IntString, args)
		if errMsg := testNativeCheckRes(res, false, expectedValue); errMsg != "" {
			t.Errorf("IntString(%v), %s", args, errMsg)
		}
		// with option error
		if s := IntString(args[0].(int64), IntStringBase(99)); s.Error != ErrBaseInvalid {
			t.Errorf("IntString(%v), %s", v[0].([]interface{}), errNull{
				"", false, ErrBaseInvalid,
				"", false, s.Error,
			})
		}
	}
}

func TestUintString(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{MaxUint64},
			"18446744073709551615",
		},
		{
			[]interface{}{uint64(0)},
			"0",
		},
		{
			[]interface{}{uint64(0), UintStringBase(10), UintStringDefault("default")},
			"default",
		},
		{
			[]interface{}{uint64(3), UintStringBase(2)},
			"11",
		},
	}
	for _, v := range testData {
		args := v[0].([]interface{})
		expectedValue := v[1].(string)
		res := rFnCall(UintString, args)
		if errMsg := testNativeCheckRes(res, false, expectedValue); errMsg != "" {
			t.Errorf("UintString(%v), %s", args, errMsg)
		}
		// with option error
		if s := UintString(args[0].(uint64), UintStringBase(99)); s.Error != ErrBaseInvalid {
			t.Errorf("UintString(%v), %s", v[0].([]interface{}), errNull{
				"", false, ErrBaseInvalid,
				"", false, s.Error,
			})
		}
	}
}

func TestFloatString(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{float64(0), FloatStringPrecision(5)},
			"0.00000e+00",
		},
		{
			[]interface{}{float64(MaxFloat64), FloatStringBitSize(32)},
			"+Inf",
		},
		{
			[]interface{}{float64(0), FloatStringFmtByte('g')},
			"0",
		},
		{
			[]interface{}{float64(0), FloatStringDefault("default")},
			"default",
		},
		{
			[]interface{}{float64(MaxFloat64)},
			"1.7976931348623157e+308",
		},
	}
	for _, v := range testData {
		args := v[0].([]interface{})
		expectedValue := v[1].(string)
		res := rFnCall(FloatString, args)
		if errMsg := testNativeCheckRes(res, false, expectedValue); errMsg != "" {
			t.Errorf("FloatString(%v), %s", args, errMsg)
		}
		// with option error
		if s := FloatString(args[0].(float64), FloatStringFmtByte('z')); s.Error != ErrFmtByteInvalid {
			t.Errorf("FloatString(%v), %s", v[0].([]interface{}), errNull{
				"", false, ErrFmtByteInvalid,
				"", false, s.Error,
			})
		}
	}
}

func TestComplexString(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{complex(float64(0), 0)},
			"(0+0i)",
		},
		{
			[]interface{}{complex(float64(0), 0), ComplexStringDefault("default")},
			"default",
		},
		{
			[]interface{}{complex(float64(MaxFloat64), 0)},
			"(1.7976931348623157e+308+0i)",
		},
	}
	for _, v := range testData {
		args := v[0].([]interface{})
		expectedValue := v[1].(string)
		res := rFnCall(ComplexString, args)
		if errMsg := testNativeCheckRes(res, false, expectedValue); errMsg != "" {
			t.Errorf("ComplexString(%v), %s", args, errMsg)
		}
	}
}

func TestString(t *testing.T) {
	testData := [][]interface{}{
		{
			[]interface{}{dTypes[reflect.Invalid]},
			[]interface{}{""},
		},
		{
			[]interface{}{true, false},
			[]interface{}{"true", "false"},
		},
		{
			[]interface{}{"something"},
			[]interface{}{"something"},
		},
		{
			[]interface{}{[]int{1, 2, 3}},
			[]interface{}{"[1 2 3]"},
		},
	}
	var expectedValue string
	for _, v := range testData {
		for i, iV := range v[0].([]interface{}) {
			expectedValue = v[1].([]interface{})[i].(string)
			value := Of(iV).String()
			if value.V() != expectedValue {
				t.Errorf("Of(%v).String(), %s", iV, errNull{
					expectedValue, false, nil,
					value.V(), false, nil,
				})
			}
		}
	}
}

// TODO: Benchmark
