package typ

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	nullUintReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullUint{}),
		reflect.TypeOf(&NotNullUint{}),
		reflect.TypeOf(&NullUint8{}),
		reflect.TypeOf(&NotNullUint8{}),
		reflect.TypeOf(&NullUint16{}),
		reflect.TypeOf(&NotNullUint16{}),
		reflect.TypeOf(&NullUint32{}),
		reflect.TypeOf(&NotNullUint32{}),
		reflect.TypeOf(&NullUint64{}),
		reflect.TypeOf(&NotNullUint64{}),
	}
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullUint{}), []dataItem{
		{reflect.ValueOf(&NullUint{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullUint{}), []dataItem{
		{reflect.ValueOf(&NotNullUint{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint8{}), []dataItem{
		{reflect.ValueOf(&NullUint8{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullUint8{}), []dataItem{
		{reflect.ValueOf(&NotNullUint8{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint16{}), []dataItem{
		{reflect.ValueOf(&NullUint16{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullUint16{}), []dataItem{
		{reflect.ValueOf(&NotNullUint16{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint32{}), []dataItem{
		{reflect.ValueOf(&NullUint32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullUint32{}), []dataItem{
		{reflect.ValueOf(&NotNullUint32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint64{}), []dataItem{
		{reflect.ValueOf(&NullUint64{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NotNullUint64{}), []dataItem{
		{reflect.ValueOf(&NotNullUint64{}), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullUintReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullUint:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullUint:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullUint8:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullUint8:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullUint16:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullUint16:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullUint32:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullUint32:
			v, null, _ = tv.V(), !tv.Valid(), tv.Present()
			present = true
		case *NullUint64:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NotNullUint64:
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
	matrixSuite.SetConverters(nullUintReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullUint:
			v, null = tv.V(), !tv.Valid()
		case *NotNullUint:
			v, null = tv.V(), !tv.Valid()
		case *NullUint8:
			v, null = tv.V(), !tv.Valid()
		case *NotNullUint8:
			v, null = tv.V(), !tv.Valid()
		case *NullUint16:
			v, null = tv.V(), !tv.Valid()
		case *NotNullUint16:
			v, null = tv.V(), !tv.Valid()
		case *NullUint32:
			v, null = tv.V(), !tv.Valid()
		case *NotNullUint32:
			v, null = tv.V(), !tv.Valid()
		case *NullUint64:
			v, null = tv.V(), !tv.Valid()
		case *NotNullUint64:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullUintReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	})
}

func TestNullUint(t *testing.T) {
	for _, nv := range []interface{}{
		&NullUint{}, &NullUint64{},
		&NullUint32{}, &NullUint16{},
		&NullUint8{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NullUint{}, NUint)
	testNType(t, &NullUint64{}, NUint64)
	testNType(t, &NullUint32{}, NUint32)
	testNType(t, &NullUint16{}, NUint16)
	testNType(t, &NullUint8{}, NUint8)
}

func TestNotNullUint(t *testing.T) {
	for _, nv := range []interface{}{
		&NotNullUint{}, &NotNullUint64{},
		&NotNullUint32{}, &NotNullUint16{},
		&NotNullUint8{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testTyp(t, nv)
		testClone(t, nv)
		testSet(t, nv)
	}
	testNType(t, &NotNullUint{}, NNUint)
	testNType(t, &NotNullUint64{}, NNUint64)
	testNType(t, &NotNullUint32{}, NNUint32)
	testNType(t, &NotNullUint16{}, NNUint16)
	testNType(t, &NotNullUint8{}, NNUint8)
}

func TestNullUintSlice(t *testing.T) {
	ns := []UintAccessor{
		NUint(0),
		NUint(1),
		&NullUint{UintCommon{Error: ErrDefaultValue}},
	}
	sl := UintSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, false), slice length not equal", ns)
	}
	sl = UintSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint64Slice(t *testing.T) {
	ns := []Uint64Accessor{
		NUint64(0),
		NUint64(1),
		&NullUint64{Uint64Common{Error: ErrDefaultValue}},
	}
	sl := Uint64Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint64Slice(%v, false), slice length not equal", ns)
	}
	sl = Uint64Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint64Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint32Slice(t *testing.T) {
	ns := []Uint32Accessor{
		NUint32(0),
		NUint32(1),
		&NullUint32{Uint32Common{Error: ErrDefaultValue}},
	}
	sl := Uint32Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint32Slice(%v, false), slice length not equal", ns)
	}
	sl = Uint32Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint32Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint16Slice(t *testing.T) {
	ns := []Uint16Accessor{
		NUint16(0),
		NUint16(1),
		&NullUint16{Uint16Common{Error: ErrDefaultValue}},
	}
	sl := Uint16Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint16Slice(%v, false), slice length not equal", ns)
	}
	sl = Uint16Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint16Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint8Slice(t *testing.T) {
	ns := []Uint8Accessor{
		NUint8(0),
		NUint8(1),
		&NullUint8{Uint8Common{Error: ErrDefaultValue}},
	}
	sl := Uint8Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint8Slice(%v, false), slice length not equal", ns)
	}
	sl = Uint8Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint8Slice(%v, true), slice length not equal", ns)
	}
}
