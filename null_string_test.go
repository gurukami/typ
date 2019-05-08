package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

type NullStringBytes struct{}

var (
	nullStringReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullString{}),
		reflect.TypeOf(&NotNullString{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullString{}), []dataItem{
		{reflect.ValueOf(&NullString{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullString{}), []dataItem{
		{reflect.ValueOf(&NotNullString{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullStringReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullString:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullString:
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
	matrixSuite.SetConverters(nullStringReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullString:
			v, null = tv.V(), !tv.Valid()
		case *NotNullString:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullStringReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullString(t *testing.T) {
	for _, nv := range []interface{}{
		&NullString{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testSet(t, nv)
		testClone(t, nv)
		testNType(t, nv, NString)
	}
}

func TestNotNullString(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullString{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testSet(t, nv)
		testClone(t, nv)
		testNType(t, nv, NNString)
	}
}

func TestNullStringSlice(t *testing.T) {
	ns := []StringAccessor{
		NString("t"),
		NString("f"),
		&NullString{StringCommon{Error: ErrDefaultValue}},
	}
	sl := StringSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullStringSlice(%v, false), slice length not equal", ns)
	}
	sl = StringSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullStringSlice(%v, true), slice length not equal", ns)
	}
}
