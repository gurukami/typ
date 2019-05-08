package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	nullInterfaceReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullInterface{}),
		reflect.TypeOf(&NotNullInterface{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullInterface{}), []dataItem{
		{reflect.ValueOf(&NullInterface{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullInterface{}), []dataItem{
		{reflect.ValueOf(&NotNullInterface{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullInterfaceReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullInterface:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullInterface:
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
	matrixSuite.SetConverters(nullInterfaceReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullInterface:
			v, null = tv.V(), !tv.Valid()
		case *NotNullInterface:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullInterfaceReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		if to == reflect.TypeOf(&NullInterface{}) {
			return &NullInterface{InterfaceCommon{P: from}}, true
		}
		if to == reflect.TypeOf(&NotNullInterface{}) {
			return &NotNullInterface{InterfaceCommon{P: from}}, true
		}
		return nil, false
	})
}

func TestNullInterface(t *testing.T) {
	for _, nv := range []interface{}{
		&NullInterface{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testSet(t, nv)
		testClone(t, nv)
		testNType(t, nv, NInterface)
	}
}

func TestNotNullInterface(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullInterface{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testSet(t, nv)
		testClone(t, nv)
		testNType(t, nv, NNInterface)
	}
}

func TestNullInterfaceSlice(t *testing.T) {
	ns := []InterfaceAccessor{
		NInterface(0),
		NInterface(1),
		&NullInterface{InterfaceCommon{Error: ErrDefaultValue}},
	}
	sl := InterfaceSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInterfaceSlice(%v, false), slice length not equal", ns)
	}
	sl = InterfaceSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInterfaceSlice(%v, true), slice length not equal", ns)
	}
}
