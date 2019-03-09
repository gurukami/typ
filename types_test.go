package typ

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

type JsonToken struct {
	Value interface{}
	Typ   reflect.Type
	Token []byte
	Err   error
}

var (
	jsonTokenReflectType  = reflect.TypeOf(JsonToken{})
	jsonTokenReflectTypes = []reflect.Type{jsonTokenReflectType}
)

func init() {
	// Test Data
	dm := map[string]interface{}{"key": "Value"}
	ds := []interface{}{1, 2, 3}
	matrixSuite.Register(reflect.TypeOf(JsonToken{}), []dataItem{
		{reflect.ValueOf(JsonToken{dm, reflect.TypeOf(dm), []byte(`{"key":"Value"}`), nil}), nil},
		{reflect.ValueOf(JsonToken{ds, reflect.TypeOf(ds), []byte(`[1,2,3]`), nil}), nil},
		{reflect.ValueOf(JsonToken{nil, nil, []byte(`{`), nil}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(SqlValueType{}), []dataItem{
		{reflect.ValueOf(SqlValueType{}), nil},
		{reflect.ValueOf(SqlValueType{struct{}{}, struct{}{}}), nil},
	})
	// Converters
	// - from JsonToken to &Null*{}
	matrixSuite.SetConverters(jsonTokenReflectTypes, nullReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		jt := from.(JsonToken)
		var v interface{}
		if err := json.Unmarshal(jt.Token, &v); err != nil {
			return nil, false
		}
		v, c, _ := matrixSuite.Convert(v, to)
		return v, c
	})
	// - from SqlValueType to &Null*{}
	matrixSuite.SetConverters(sqlValueReflectTypes, nullReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		sv := from.(SqlValueType)
		if to == reflect.TypeOf(&NullInterface{}) {
			if reflect.TypeOf(from) != to {
				return nil, false
			}
		}
		v, c, _ := matrixSuite.Convert(sv.Value, to)
		return v, c
	})
	// For other types
	matrixSuite.SetConverter(nil, jsonTokenReflectType, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		b, err := json.Marshal(from)
		jt := JsonToken{from, rv.Type(), b, err}
		return jt, err == nil
	})
	matrixSuite.SetConverters(interfaceReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

var overFlowValue = string(bytes.Repeat([]byte("1"), 310))

func TestEquals(t *testing.T) {
	testData := [][]interface{}{
		{
			MaxInt64,
			[]interface{}{"9223372036854775807"},
			true,
			true,
			nil,
		},
		{
			MaxUint64,
			[]interface{}{"18446744073709551615"},
			true,
			true,
			nil,
		},
		{
			float64(MaxInt16),
			[]interface{}{32767},
			true,
			true,
			nil,
		},
		{
			complex(float64(0), 0),
			[]interface{}{"(0+0i)"},
			true,
			true,
			nil,
		},
		{
			true,
			[]interface{}{"true"},
			true,
			true,
			nil,
		},
		{
			"1",
			[]interface{}{1},
			true,
			true,
			nil,
		},
		{
			map[string]interface{}{"k": 1},
			[]interface{}{1},
			false,
			true,
			nil,
		},
	}
	var (
		val           interface{}
		args          []interface{}
		expectedValue bool
		expectedValid bool
		expectedError error
		value         NullBool
	)
	for _, v := range testData {
		val = v[0].(interface{})
		args = v[1].([]interface{})
		expectedValue = v[2].(bool)
		expectedValid = v[3].(bool)
		expectedError, _ = v[4].(error)
		typ := Of(val)
		res := rFnCall(typ.Equals, args)
		value = res[0].(NullBool)
		if value.V() != expectedValue || value.Valid() != expectedValid || value.Error != expectedError {
			t.Errorf("Of(%v).Equals(%v) = %v failed, expects %v, Error: %v Valid: %v", val, args, res[0], expectedValue, value.Error, value.Valid())
		}
	}
	// Nil tests
	value = Of(nil).Equals(nil)
	if value.V() || value.Valid() || value.Error == nil {
		t.Errorf("Of(%v).Equals(%v) = %v failed, expects %v, Error: %v Valid: %v", nil, nil, value.V(), false, value.Error, value.Valid())
	}
}

func TestIdentical(t *testing.T) {
	testData := [][]interface{}{
		{
			MaxInt64,
			[]interface{}{"9223372036854775807"},
			false,
			true,
			nil,
		},
		{
			MaxUint64,
			[]interface{}{"18446744073709551615"},
			false,
			true,
			nil,
		},
		{
			float64(MaxInt16),
			[]interface{}{32767},
			false,
			true,
			nil,
		},
		{
			complex(float64(0), 0),
			[]interface{}{"(0+0i)"},
			false,
			true,
			nil,
		},
		{
			true,
			[]interface{}{"true"},
			false,
			true,
			nil,
		},
		{
			"1",
			[]interface{}{1},
			false,
			true,
			nil,
		},
		{
			map[string]interface{}{"k": 1},
			[]interface{}{1},
			false,
			true,
			nil,
		},
		{
			MaxInt64,
			[]interface{}{MaxInt64},
			true,
			true,
			nil,
		},
		{
			MaxUint64,
			[]interface{}{MaxUint64},
			true,
			true,
			nil,
		},
		{
			float64(MaxInt16),
			[]interface{}{float64(MaxInt16)},
			true,
			true,
			nil,
		},
		{
			complex(float64(0), 0),
			[]interface{}{complex(float64(0), 0)},
			true,
			true,
			nil,
		},
		{
			true,
			[]interface{}{true},
			true,
			true,
			nil,
		},
		{
			1,
			[]interface{}{1},
			true,
			true,
			nil,
		},
	}
	var (
		val           interface{}
		args          []interface{}
		expectedValue bool
		expectedValid bool
		expectedError error
		value         NullBool
	)
	for _, v := range testData {
		val = v[0].(interface{})
		args = v[1].([]interface{})
		expectedValue = v[2].(bool)
		expectedValid = v[3].(bool)
		expectedError, _ = v[4].(error)
		typ := Of(val)
		res := rFnCall(typ.Identical, args)
		value = res[0].(NullBool)
		if value.V() != expectedValue || value.Valid() != expectedValid || value.Error != expectedError {
			t.Errorf("Of(%v).Identical(%v) = %v failed, expects %v, Error: %v Valid: %v", val, args, res[0], expectedValue, value.Error, value.Valid())
		}
	}
	// Nil tests
	value = Of(nil).Identical(nil)
	if value.V() || value.Valid() || value.Error == nil {
		t.Errorf("Of(%v).Identical(%v) = %v failed, expects %v, Error: %v Valid: %v", nil, nil, value.V(), false, value.Error, value.Valid())
	}
}

func TestInterface(t *testing.T) {
	for k, v := range dTypes {
		if Of(v).Interface().V() == nil && k != reflect.Invalid && k != reflect.Interface {
			t.Errorf("Of(%v).Interface() as %T failed", v, v)
		}
	}
}

func TestEmpty(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	testData := [][]interface{}{
		{
			[]interface{}{nil},
			[]bool{true},
			[]bool{false},
			[]error{ErrUnexpectedValue},
		},
		{
			[]interface{}{true, false},
			[]bool{false, true},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{int(0), MinInt, MaxInt},
			[]bool{true, false, false},
			[]bool{true, true, true},
			[]error{nil, nil, nil},
		},
		{
			[]interface{}{int8(0), MinInt8, MaxInt8},
			[]bool{true, false, false},
			[]bool{true, true, true},
			[]error{nil, nil, nil},
		},
		{
			[]interface{}{int16(0), MinInt16, MaxInt16},
			[]bool{true, false, false},
			[]bool{true, true, true},
			[]error{nil, nil, nil},
		},
		{
			[]interface{}{int32(0), MinInt32, MaxInt32},
			[]bool{true, false, false},
			[]bool{true, true, true},
			[]error{nil, nil, nil},
		},
		{
			[]interface{}{int64(0), MinInt64, MaxInt64},
			[]bool{true, false, false},
			[]bool{true, true, true},
			[]error{nil, nil, nil},
		},
		{
			[]interface{}{uint(0), MaxUint},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{uint8(0), MaxUint8},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{uint16(0), MaxUint16},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{uint32(0), MaxUint32},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{uint64(0), MaxUint64},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{uintptr(0), uintptr(unsafePointer)},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				float32(0),
				MinFloat32,
				MaxFloat32,
				MinSafeIntFloat32,
				MaxSafeIntFloat32,
				math.SmallestNonzeroFloat32,
				float32(math.NaN()),
				float32(math.Inf(-1)),
				float32(math.Inf(1)),
			},
			[]bool{true, false, false, false, false, false, false, false, false},
			[]bool{true, true, true, true, true, true, true, true, true},
			[]error{nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			[]interface{}{
				float64(0),
				MinFloat64,
				MaxFloat64,
				MinSafeIntFloat64,
				MaxSafeIntFloat64,
				math.SmallestNonzeroFloat64,
				float64(math.NaN()),
				float64(math.Inf(-1)),
				float64(math.Inf(1)),
			},
			[]bool{true, false, false, false, false, false, false, false, false},
			[]bool{true, true, true, true, true, true, true, true, true},
			[]error{nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			[]interface{}{
				complex64(complex(float32(0), float32(0))),
				complex64(complex(MinFloat32, MinFloat32)),
				complex64(complex(MaxFloat32, MaxFloat32)),
				complex64(complex(MinSafeIntFloat32, MinSafeIntFloat32)),
				complex64(complex(MaxSafeIntFloat32, MaxSafeIntFloat32)),
				complex64(complex(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32)),
				complex64(complex(math.NaN(), math.NaN())),
				complex64(complex(math.Inf(-1), math.Inf(-1))),
				complex64(complex(math.Inf(1), math.Inf(1))),
			},
			[]bool{true, false, false, false, false, false, false, false, false},
			[]bool{true, true, true, true, true, true, true, true, true},
			[]error{nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			[]interface{}{
				complex128(complex(float64(0), float64(0))),
				complex128(complex(MinFloat64, MinFloat64)),
				complex128(complex(MaxFloat64, MaxFloat64)),
				complex128(complex(MinSafeIntFloat64, MinSafeIntFloat64)),
				complex128(complex(MaxSafeIntFloat64, MaxSafeIntFloat64)),
				complex128(complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64)),
				complex128(complex(math.NaN(), math.NaN())),
				complex128(complex(math.Inf(-1), math.Inf(-1))),
				complex128(complex(math.Inf(1), math.Inf(1))),
			},
			[]bool{true, false, false, false, false, false, false, false, false},
			[]bool{true, true, true, true, true, true, true, true, true},
			[]error{nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
		{
			[]interface{}{
				[1]int{}, [1]int{1},
			},
			[]bool{false, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				make(chan int), ch,
			},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				dTypes[reflect.Func],
			},
			[]bool{false},
			[]bool{true},
			[]error{nil},
		},
		{
			[]interface{}{
				[]int{}, []int{1},
			},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				"", "nonempty",
			},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				dTypes[reflect.Struct],
			},
			[]bool{false},
			[]bool{true},
			[]error{nil},
		},
		{
			[]interface{}{
				unsafe.Pointer(uintptr(0)),
				dTypes[reflect.UnsafePointer],
			},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				make(map[int]int), map[int]int{0: 1},
			},
			[]bool{true, false},
			[]bool{true, true},
			[]error{nil, nil},
		},
		{
			[]interface{}{
				dTypes[reflect.Ptr],
			},
			[]bool{false},
			[]bool{true},
			[]error{nil},
		},
		{
			[]interface{}{
				dTypes[reflect.Interface],
			},
			[]bool{false},
			[]bool{true},
			[]error{nil},
		},
	}
	var (
		expectedValue bool
		expectedValid bool
		expectedError error
		value         NullBool
	)
	for _, v := range testData {
		for k, iV := range v[0].([]interface{}) {
			expectedValue = v[1].([]bool)[k]
			expectedValid = v[2].([]bool)[k]
			expectedError = v[3].([]error)[k]
			value = Of(iV).Empty()
			if value.V() != expectedValue || value.Valid() != expectedValid || value.Error != expectedError {
				t.Errorf("Of(%v).Empty() as %T type failed, expects %v, actual %v, Error: %v Valid: %v", iV, iV, expectedValue, value.V(), value.Error, value.Valid())
			}
		}
	}
	// Nil tests
	value = Of(nil).Empty()
	if !value.V() || value.Valid() || value.Error == nil {
		t.Errorf("Of(%v).Empty() = %v failed, expects %v, Error: %v Valid: %v", nil, value.V(), false, value.Error, value.Valid())
	}
}

func testNativeCheckRes(testResults []interface{}, expectedSafe bool, defaultValue interface{}) string {
	value, _, valid, err := testGetNullIfaceValue(reflect.Indirect(reflect.ValueOf(testResults[0])).Interface())
	deepEqual := reflect.DeepEqual(value, defaultValue)
	if (expectedSafe && !valid) || (!expectedSafe && !deepEqual && defaultValue != nil) {
		return errNull{
			defaultValue, expectedSafe, nil,
			value, valid, err,
		}.String()
	}
	return ""
}

func TestGeneric(t *testing.T) {
	typ := Of(1.1, Precision(7), FmtByte('G'))
	typ = Of(typ)
	if typ.OptionFmtByte() != 'G' || typ.OptionPrecision() != 7 {
		t.Error("Of(Of(nil), 7, 2, 'g') failed, copy of struct expected")
	}
	if Err := Of(nil).Int().Error; Err == nil || Err.Error() == "" {
		t.Error("Of(nil).Int() failed, expects non empty error")
	}
}

func TestSetFloatPrecision(t *testing.T) {
	typ := Of(nil)
	if typ.OptionPrecision() != -1 {
		t.Error("Of(nil).OptionPrecision() failed, default precision is changed")
	}
	typ = Of(1.1, Precision(7))
	if typ.OptionPrecision() != 7 {
		t.Error("Of(nil, 7).OptionPrecision() failed, expects 7")
	}
}

func TestSetGetFloatFmt(t *testing.T) {
	typ := Of(nil)
	if typ.OptionFmtByte() != 'e' {
		t.Error("Of(nil).OptionFmtByte() failed, default float fmt is changed")
	}
	typ = Of(nil, FmtByte('g'))
	if typ.OptionFmtByte() != 'g' {
		t.Error("Of(nil, 'g').OptionFmtByte() failed, expects `g`")
	}
}

func TestSetGetBase(t *testing.T) {
	typ := Of(nil)
	if typ.OptionBase() != 10 {
		t.Error("Of(nil).OptionBase() failed, default base is changed")
	}
	typ = Of(nil, Base(2))
	if typ.OptionBase() != 2 {
		t.Error("Of(nil, 2).OptionBase() failed, expects 2")
	}
}

type SqlValueType struct {
	SqlValue driver.Value
	Value    interface{}
}

func testGetNullIfaceValue(nv interface{}) (value interface{}, nkind reflect.Kind, valid bool, err error) {
	if nv == nil {
		return nil, reflect.Invalid, false, nil
	}
	nv = reflect.Indirect(reflect.ValueOf(nv)).Interface()
	switch v := nv.(type) {
	case NullBool:
		return v.V(), reflect.Bool, v.Valid(), v.Error
	case NullComplex64:
		return v.V(), reflect.Complex64, v.Valid(), v.Error
	case NullComplex:
		return v.V(), reflect.Complex128, v.Valid(), v.Error
	case NullInt:
		return v.V(), reflect.Int, v.Valid(), v.Error
	case NullInt8:
		return v.V(), reflect.Int8, v.Valid(), v.Error
	case NullInt16:
		return v.V(), reflect.Int16, v.Valid(), v.Error
	case NullInt32:
		return v.V(), reflect.Int32, v.Valid(), v.Error
	case NullInt64:
		return v.V(), reflect.Int64, v.Valid(), v.Error
	case NullUint:
		return v.V(), reflect.Uint, v.Valid(), v.Error
	case NullUint8:
		return v.V(), reflect.Uint8, v.Valid(), v.Error
	case NullUint16:
		return v.V(), reflect.Uint16, v.Valid(), v.Error
	case NullUint32:
		return v.V(), reflect.Uint32, v.Valid(), v.Error
	case NullUint64:
		return v.V(), reflect.Uint64, v.Valid(), v.Error
	case NullFloat32:
		return v.V(), reflect.Float32, v.Valid(), v.Error
	case NullFloat:
		return v.V(), reflect.Float64, v.Valid(), v.Error
	case NullString:
		return v.V(), reflect.String, v.Valid(), v.Error
	case NullTime:
		return v.V(), reflect.Struct, v.Valid(), v.Error
	case NullInterface:
		return v.V(), reflect.ValueOf(v.V()).Kind(), v.Valid(), v.Error
	}
	return nil, reflect.Invalid, false, nil
}

func testOfDefault(t *testing.T, v interface{}, methodTypeName string, defaultValue interface{}) {
	rt := reflect.ValueOf(Of(v))
	rm := rt.MethodByName(methodTypeName)
	var rres []reflect.Value
	if defaultValue != nil {
		rres = rm.Call([]reflect.Value{reflect.ValueOf(defaultValue)})
	} else {
		rres = rm.Call([]reflect.Value{})
	}
	rd := reflect.ValueOf(defaultValue)
	actualValue, ar, actualValid, actualError := testGetNullIfaceValue(reflect.Indirect(rres[0]).Interface())
	eValue, eValid, eError := matrixSuite.Test(v, getDefaultType(rd.Kind()))
	if eValue == nil {
		return
	}
	if v != nil && reflect.TypeOf(v).Kind() == ar {
		return
	}
	if !eValid && defaultValue != nil {
		eValue = defaultValue
	}
	if !matrixSuite.Compare(actualValue, eValue) || actualValid != eValid {
		t.Errorf("Of(%T(%+[1]v)).%s(%T(%[3]v)), %s", v, methodTypeName, defaultValue, errNull{
			eValue, eValid, eError,
			actualValue, actualValid, actualError,
		})
	}
}

func testOfDefaultNil(t *testing.T, methodTypeName string) {
	rt := reflect.ValueOf(Of(nil))
	rm := rt.MethodByName(methodTypeName)
	res := rm.Call([]reflect.Value{})
	_, _, _, actualError := testGetNullIfaceValue(reflect.Indirect(res[0]).Interface())
	if _, ok := actualError.(ErrorConvert); !ok {
		t.Errorf("Of(nil).%s(), must returns 'ErrorConvert' error instead of '%T'", methodTypeName, actualError)
	}
}

func testNative(t *testing.T, fn interface{}, args []interface{}) {
	rf := reflect.ValueOf(fn)
	var rres, rargs []reflect.Value
	for _, arg := range args {
		rargs = append(rargs, reflect.ValueOf(arg))
	}
	var defaultValue interface{}
	if len(rargs) > 1 {
		defaultValue = rargs[1].Interface()
	}
	rres = rf.Call(rargs)
	actualValue, _, actualValid, actualError := testGetNullIfaceValue(reflect.Indirect(rres[0]).Interface())
	eValue, eValid, eError := matrixSuite.Test(rargs[0].Interface(), getDefaultType(rargs[1].Kind()))
	if eValue == nil {
		return
	}
	if !actualValid && defaultValue != nil {
		eValue = defaultValue
	}
	if !matrixSuite.Compare(actualValue, eValue) || actualValid != eValid {
		t.Errorf("%s(%v), %s", rf.String(), args, errNull{
			eValue, eValid, eError,
			actualValue, actualValid, actualError,
		})
	}
}
