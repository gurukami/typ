package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	nullBoolReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullBool{}),
		reflect.TypeOf(&NotNullBool{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullBool{}), []dataItem{
		{reflect.ValueOf(&NullBool{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullBool{}), []dataItem{
		{reflect.ValueOf(&NotNullBool{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullBoolReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullBool:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullBool:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		}
		rv := reflect.ValueOf(v)
		b, err := json.Marshal(v)
		if null || !present {
			b, err = []byte("null"), nil
		}
		if !rv.IsValid() {
			return nil, false
		}
		jt := JSONToken{from, rv.Type(), b, err}
		return jt, err == nil
	})
	// - from &Null*{} to SQLValueType
	matrixSuite.SetConverters(nullBoolReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullBool:
			v, null = tv.V(), !tv.Valid()
		case *NotNullBool:
			v, null = tv.V(), !tv.Valid()
		}
		if null || v == nil {
			return SQLValueType{}, true
		}
		rv := reflect.ValueOf(v)
		cv, valid, _ := matrixSuite.Convert(rv.Interface(), to)
		if !valid {
			return SQLValueType{}, false
		}
		return cv, driver.IsValue(cv.(SQLValueType).SQLValue)
	})
	// For other types
	matrixSuite.SetConverters(interfaceReflectTypes, nullBoolReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullBool(t *testing.T) {
	testMarshalJSON(t, &NullBool{})
	testUnmarshalJSON(t, &NullBool{})
	testScanSQL(t, &NullBool{})
	testValueSQL(t, &NullBool{})
	testTyp(t, &NullBool{})
	testClone(t, &NullBool{})
	testSet(t, &NullBool{})
	testNType(t, &NullBool{}, NBool)
}

func TestNotNullBool(t *testing.T) {
	testMarshalJSON(t, &NotNullBool{})
	testUnmarshalJSON(t, &NotNullBool{})
	testScanSQL(t, &NotNullBool{})
	testValueSQL(t, &NotNullBool{})
	testTyp(t, &NotNullBool{})
	testClone(t, &NotNullBool{})
	testSet(t, &NotNullBool{})
	testNType(t, &NotNullBool{}, NNBool)
}

func TestNullBoolSlice(t *testing.T) {
	ns := []BoolAccessor{
		NBool(true),
		NBool(false),
		&NullBool{BoolCommon{Error: ErrDefaultValue}},
	}
	sl := BoolSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullBoolSlice(%v, false), slice length not equal", ns)
	}
	sl = BoolSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullBoolSlice(%v, true), slice length not equal", ns)
	}
}
