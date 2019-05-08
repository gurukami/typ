package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	complexNullReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullComplex{}),
		reflect.TypeOf(&NotNullComplex{}),
		reflect.TypeOf(&NullComplex64{}),
		reflect.TypeOf(&NotNullComplex64{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullComplex{}), []dataItem{
		{reflect.ValueOf(&NullComplex{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullComplex{}), []dataItem{
		{reflect.ValueOf(&NotNullComplex{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullComplex64{}), []dataItem{
		{reflect.ValueOf(&NullComplex64{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullComplex64{}), []dataItem{
		{reflect.ValueOf(&NotNullComplex64{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(complexNullReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullComplex:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullComplex:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullComplex64:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullComplex64:
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
	matrixSuite.SetConverters(complexNullReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullComplex:
			v, null = tv.V(), !tv.Valid()
		case *NotNullComplex:
			v, null = tv.V(), !tv.Valid()
		case *NullComplex64:
			v, null = tv.V(), !tv.Valid()
		case *NotNullComplex64:
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
	matrixSuite.SetConverters(interfaceReflectTypes, complexNullReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullComplex(t *testing.T) {
	for _, nv := range []interface{}{
		&NullComplex64{}, &NullComplex{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NullComplex64{}, NComplex64)
	testNType(t, &NullComplex{}, NComplex)
}

func TestNotNullComplex(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullComplex64{}, &NotNullComplex{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NotNullComplex64{}, NNComplex64)
	testNType(t, &NotNullComplex{}, NNComplex)
}

func TestNullComplexSlice(t *testing.T) {
	ns := []ComplexAccessor{
		NComplex(0),
		NComplex(1),
		&NullComplex{ComplexCommon{Error: ErrDefaultValue}},
	}
	sl := ComplexSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, false), slice length not equal", ns)
	}
	sl = ComplexSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullComplex64Slice(t *testing.T) {
	ns := []Complex64Accessor{
		NComplex64(0),
		NComplex64(1),
		&NullComplex64{Complex64Common{Error: ErrDefaultValue}},
	}
	sl := Complex64Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplex64Slice(%v, false), slice length not equal", ns)
	}
	sl = Complex64Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplex64Slice(%v, true), slice length not equal", ns)
	}
}
