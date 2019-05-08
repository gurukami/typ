package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var (
	nullTimeReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullTime{}),
		reflect.TypeOf(&NotNullTime{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullTime{}), []dataItem{
		{reflect.ValueOf(&NullTime{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullTime{}), []dataItem{
		{reflect.ValueOf(&NotNullTime{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(time.Time{}), []dataItem{
		{reflect.ValueOf(time.Now()), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullTimeReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullTime:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullTime:
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
	matrixSuite.SetConverters([]reflect.Type{reflect.TypeOf(time.Time{})}, nullTimeReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		v := from.(time.Time)
		if to == reflect.TypeOf(&NullTime{}) {
			return &NullTime{TimeCommon{P: &v}}, true
		}
		if to == reflect.TypeOf(&NotNullTime{}) {
			return &NotNullTime{TimeCommon{P: &v}}, true
		}
		return nil, false
	})
	matrixSuite.SetConverter(reflect.TypeOf(time.Time{}), sqlValueReflectType, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		sv := from.(time.Time)
		return SQLValueType{sv, from}, true
	})
	// - from &Null*{} to SQLValueType
	matrixSuite.SetConverters(nullTimeReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullTime:
			v, null = tv.V(), !tv.Valid()
		case *NotNullTime:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullTimeReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullTime(t *testing.T) {
	for _, nv := range []interface{}{
		&NullTime{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
		testNType(t, nv, NTime)
	}
}

func TestNotNullTime(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullTime{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
		testNType(t, nv, NNTime)
	}
}

func TestNullTimeSlice(t *testing.T) {
	ns := []TimeAccessor{
		NTime(time.Now()),
		NTime(time.Now()),
		&NullTime{TimeCommon{Error: ErrDefaultValue}},
	}
	sl := TimeSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullTimeSlice(%v, false), slice length not equal", ns)
	}
	sl = TimeSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullTimeSlice(%v, true), slice length not equal", ns)
	}
}
