package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	nullFloatReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullFloat32{}),
		reflect.TypeOf(&NotNullFloat32{}),
		reflect.TypeOf(&NullFloat{}),
		reflect.TypeOf(&NotNullFloat{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullFloat32{}), []dataItem{
		{reflect.ValueOf(&NullFloat32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullFloat32{}), []dataItem{
		{reflect.ValueOf(&NotNullFloat32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullFloat{}), []dataItem{
		{reflect.ValueOf(&NullFloat{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullFloat{}), []dataItem{
		{reflect.ValueOf(&NotNullFloat{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullFloatReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullFloat32:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullFloat32:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullFloat:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullFloat:
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
	matrixSuite.SetConverters(nullFloatReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullFloat32:
			v, null = tv.V(), !tv.Valid()
		case *NotNullFloat32:
			v, null = tv.V(), !tv.Valid()
		case *NullFloat:
			v, null = tv.V(), !tv.Valid()
		case *NotNullFloat:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullFloatReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullFloat(t *testing.T) {
	for _, nv := range []interface{}{
		&NullFloat32{}, &NullFloat{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NullFloat32{}, NFloat32)
	testNType(t, &NullFloat{}, NFloat)
}

func TestNotNullFloat(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullFloat32{}, &NotNullFloat{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NotNullFloat32{}, NNFloat32)
	testNType(t, &NotNullFloat{}, NNFloat)
}

func TestNullFloatSlice(t *testing.T) {
	ns := []FloatAccessor{
		NFloat(0),
		NFloat(1),
		&NullFloat{FloatCommon{Error: ErrDefaultValue}},
	}
	sl := FloatSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullFloatSlice(%v, false), slice length not equal", ns)
	}
	sl = FloatSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullFloatSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullFloat32Slice(t *testing.T) {
	ns := []Float32Accessor{
		NFloat32(0),
		NFloat32(1),
		&NullFloat32{Float32Common{Error: ErrDefaultValue}},
	}
	sl := Float32Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullFloat32Slice(%v, false), slice length not equal", ns)
	}
	sl = Float32Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullFloat32Slice(%v, true), slice length not equal", ns)
	}
}
