package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	nullIntReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullInt{}),
		reflect.TypeOf(&NullInt8{}),
		reflect.TypeOf(&NullInt16{}),
		reflect.TypeOf(&NullInt32{}),
		reflect.TypeOf(&NullInt64{}),
		reflect.TypeOf(&NotNullInt{}),
		reflect.TypeOf(&NotNullInt8{}),
		reflect.TypeOf(&NotNullInt16{}),
		reflect.TypeOf(&NotNullInt32{}),
		reflect.TypeOf(&NotNullInt64{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullInt{}), []dataItem{
		{reflect.ValueOf(&NullInt{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullInt{}), []dataItem{
		{reflect.ValueOf(&NotNullInt{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt8{}), []dataItem{
		{reflect.ValueOf(&NullInt8{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullInt8{}), []dataItem{
		{reflect.ValueOf(&NotNullInt8{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt16{}), []dataItem{
		{reflect.ValueOf(&NullInt16{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullInt16{}), []dataItem{
		{reflect.ValueOf(&NotNullInt16{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt32{}), []dataItem{
		{reflect.ValueOf(&NullInt32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullInt32{}), []dataItem{
		{reflect.ValueOf(&NotNullInt32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt64{}), []dataItem{
		{reflect.ValueOf(&NullInt64{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullInt64{}), []dataItem{
		{reflect.ValueOf(&NotNullInt64{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullIntReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullInt:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullInt:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullInt8:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullInt8:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullInt16:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullInt16:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullInt32:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullInt32:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullInt64:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullInt64:
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
	matrixSuite.SetConverters(nullIntReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullInt:
			v, null = tv.V(), !tv.Valid()
		case *NotNullInt:
			v, null = tv.V(), !tv.Valid()
		case *NullInt8:
			v, null = tv.V(), !tv.Valid()
		case *NotNullInt8:
			v, null = tv.V(), !tv.Valid()
		case *NullInt16:
			v, null = tv.V(), !tv.Valid()
		case *NotNullInt16:
			v, null = tv.V(), !tv.Valid()
		case *NullInt32:
			v, null = tv.V(), !tv.Valid()
		case *NotNullInt32:
			v, null = tv.V(), !tv.Valid()
		case *NullInt64:
			v, null = tv.V(), !tv.Valid()
		case *NotNullInt64:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullIntReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullInt(t *testing.T) {
	for _, nv := range []interface{}{
		&NullInt{}, &NullInt64{},
		&NullInt32{}, &NullInt16{},
		&NullInt8{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NullInt{}, NInt)
	testNType(t, &NullInt64{}, NInt64)
	testNType(t, &NullInt32{}, NInt32)
	testNType(t, &NullInt16{}, NInt16)
	testNType(t, &NullInt8{}, NInt8)
}

func TestNotNullInt(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullInt{}, &NotNullInt64{},
		&NotNullInt32{}, &NotNullInt16{},
		&NotNullInt8{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NotNullInt{}, NNInt)
	testNType(t, &NotNullInt64{}, NNInt64)
	testNType(t, &NotNullInt32{}, NNInt32)
	testNType(t, &NotNullInt16{}, NNInt16)
	testNType(t, &NotNullInt8{}, NNInt8)
}

func TestNullIntSlice(t *testing.T) {
	ns := []IntAccessor{
		NInt(0),
		NInt(1),
		&NullInt{IntCommon{Error: ErrDefaultValue}},
	}
	sl := IntSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, false), slice length not equal", ns)
	}
	sl = IntSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt64Slice(t *testing.T) {
	ns := []Int64Accessor{
		NInt64(0),
		NInt64(1),
		&NullInt64{Int64Common{Error: ErrDefaultValue}},
	}
	sl := Int64Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt64Slice(%v, false), slice length not equal", ns)
	}
	sl = Int64Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt64Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt32Slice(t *testing.T) {
	ns := []Int32Accessor{
		NInt32(0),
		NInt32(1),
		&NullInt32{Int32Common{Error: ErrDefaultValue}},
	}
	sl := Int32Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt32Slice(%v, false), slice length not equal", ns)
	}
	sl = Int32Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt32Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt16Slice(t *testing.T) {
	ns := []Int16Accessor{
		NInt16(0),
		NInt16(1),
		&NullInt16{Int16Common{Error: ErrDefaultValue}},
	}
	sl := Int16Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt16Slice(%v, false), slice length not equal", ns)
	}
	sl = Int16Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt16Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt8Slice(t *testing.T) {
	ns := []Int8Accessor{
		NInt8(0),
		NInt8(1),
		&NullInt8{Int8Common{Error: ErrDefaultValue}},
	}
	sl := Int8Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt8Slice(%v, false), slice length not equal", ns)
	}
	sl = Int8Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt8Slice(%v, true), slice length not equal", ns)
	}
}
