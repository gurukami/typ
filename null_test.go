package typ

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

type NullStringBytes struct{}

var (
	nullReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullBool{}),
		reflect.TypeOf(&NullComplex{}),
		reflect.TypeOf(&NullComplex64{}),
		reflect.TypeOf(&NullFloat32{}),
		reflect.TypeOf(&NullFloat{}),
		reflect.TypeOf(&NullInt{}),
		reflect.TypeOf(&NullInt8{}),
		reflect.TypeOf(&NullInt16{}),
		reflect.TypeOf(&NullInt32{}),
		reflect.TypeOf(&NullInt64{}),
		reflect.TypeOf(&NullUint{}),
		reflect.TypeOf(&NullUint8{}),
		reflect.TypeOf(&NullUint16{}),
		reflect.TypeOf(&NullUint32{}),
		reflect.TypeOf(&NullUint64{}),
		reflect.TypeOf(&NullString{}),
		reflect.TypeOf(&NullInterface{}),
		reflect.TypeOf(&NullTime{}),
	}
	nullStringReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullString{}),
	}
	nullIntReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullInt{}),
		reflect.TypeOf(&NullInt8{}),
		reflect.TypeOf(&NullInt16{}),
		reflect.TypeOf(&NullInt32{}),
		reflect.TypeOf(&NullInt64{}),
	}
	nullUintReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullUint{}),
		reflect.TypeOf(&NullUint8{}),
		reflect.TypeOf(&NullUint16{}),
		reflect.TypeOf(&NullUint32{}),
		reflect.TypeOf(&NullUint64{}),
	}
	nullFloatReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullFloat32{}),
		reflect.TypeOf(&NullFloat{}),
	}
	nullComplexReflectTypes = []reflect.Type{
		reflect.TypeOf(&NullComplex{}),
		reflect.TypeOf(&NullComplex64{}),
	}
	sqlValueReflectTypes = []reflect.Type{
		reflect.TypeOf(SQLValueType{}),
	}
	sqlValueReflectType = reflect.TypeOf(SQLValueType{})
)

func init() {
	// Test Data
	matrixSuite.Register(reflect.TypeOf(&NullBool{}), []dataItem{
		{reflect.ValueOf(&NullBool{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullComplex{}), []dataItem{
		{reflect.ValueOf(&NullComplex{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullComplex64{}), []dataItem{
		{reflect.ValueOf(&NullComplex64{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullFloat32{}), []dataItem{
		{reflect.ValueOf(&NullFloat32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullFloat{}), []dataItem{
		{reflect.ValueOf(&NullFloat{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt{}), []dataItem{
		{reflect.ValueOf(&NullInt{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt8{}), []dataItem{
		{reflect.ValueOf(&NullInt8{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt16{}), []dataItem{
		{reflect.ValueOf(&NullInt16{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt32{}), []dataItem{
		{reflect.ValueOf(&NullInt32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInt64{}), []dataItem{
		{reflect.ValueOf(&NullInt64{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint{}), []dataItem{
		{reflect.ValueOf(&NullUint{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint8{}), []dataItem{
		{reflect.ValueOf(&NullUint8{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint16{}), []dataItem{
		{reflect.ValueOf(&NullUint16{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint32{}), []dataItem{
		{reflect.ValueOf(&NullUint32{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullUint64{}), []dataItem{
		{reflect.ValueOf(&NullUint64{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullString{}), []dataItem{
		{reflect.ValueOf(&NullString{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullInterface{}), []dataItem{
		{reflect.ValueOf(&NullInterface{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(&NullTime{}), []dataItem{
		{reflect.ValueOf(&NullTime{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(time.Time{}), []dataItem{
		{reflect.ValueOf(time.Now()), nil},
	})
	// Converters
	// - from &Null*{} to JSONToken
	matrixSuite.SetConverters(nullReflectTypes, jsonTokenReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v       interface{}
			null    bool
			present bool
		)
		switch tv := from.(type) {
		case *NullBool:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullComplex:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullComplex64:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullFloat32:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullFloat:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullInt:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullInt8:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullInt16:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullInt32:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullInt64:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullUint:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullUint8:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullUint16:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullUint32:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullUint64:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullString:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullInterface:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
		case *NullTime:
			v, null, present = tv.V(), !tv.Valid(), tv.Present()
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
	matrixSuite.SetConverter(reflect.TypeOf(time.Time{}), reflect.TypeOf(&NullTime{}), func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		v := from.(time.Time)
		return &NullTime{P: &v}, true
	})
	matrixSuite.SetConverter(reflect.TypeOf(time.Time{}), sqlValueReflectType, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		sv := from.(time.Time)
		return SQLValueType{sv, from}, true
	})
	// - from &Null*{} to SQLValueType
	matrixSuite.SetConverters(nullReflectTypes, sqlValueReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		var (
			v    interface{}
			null bool
		)
		switch tv := from.(type) {
		case *NullBool:
			v, null = tv.V(), !tv.Valid()
		case *NullComplex:
			v, null = tv.V(), !tv.Valid()
		case *NullComplex64:
			v, null = tv.V(), !tv.Valid()
		case *NullFloat32:
			v, null = tv.V(), !tv.Valid()
		case *NullFloat:
			v, null = tv.V(), !tv.Valid()
		case *NullInt:
			v, null = tv.V(), !tv.Valid()
		case *NullInt8:
			v, null = tv.V(), !tv.Valid()
		case *NullInt16:
			v, null = tv.V(), !tv.Valid()
		case *NullInt32:
			v, null = tv.V(), !tv.Valid()
		case *NullInt64:
			v, null = tv.V(), !tv.Valid()
		case *NullUint:
			v, null = tv.V(), !tv.Valid()
		case *NullUint8:
			v, null = tv.V(), !tv.Valid()
		case *NullUint16:
			v, null = tv.V(), !tv.Valid()
		case *NullUint32:
			v, null = tv.V(), !tv.Valid()
		case *NullUint64:
			v, null = tv.V(), !tv.Valid()
		case *NullString:
			v, null = tv.V(), !tv.Valid()
		case *NullInterface:
			v, null = tv.V(), !tv.Valid()
		case *NullTime:
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
	matrixSuite.SetConverters(interfaceReflectTypes, nullReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		if to == reflect.TypeOf(&NullInterface{}) {
			return &NullInterface{P: from}, true
		}
		return nil, false
	})
}

func TestNullBool(t *testing.T) {
	testMarshalJSON(t, &NullBool{})
	testUnmarshalJSON(t, &NullBool{})
	testScanSQL(t, &NullBool{})
	testValueSQL(t, &NullBool{})
	testTyp(t, &NullBool{})
	testSet(t, &NullBool{})
	testNType(t, &NullBool{}, NBool)
}

func TestNullBoolSlice(t *testing.T) {
	ns := []NullBool{
		NBool(true),
		NBool(false),
		{Error: ErrDefaultValue},
	}
	sl := NullBoolSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullBoolSlice(%v, false), slice length not equal", ns)
	}
	sl = NullBoolSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullBoolSlice(%v, true), slice length not equal", ns)
	}
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
		testSet(t, nv)
	}
	testNType(t, &NullComplex64{}, NComplex64)
	testNType(t, &NullComplex{}, NComplex)
}

func TestNullComplexSlice(t *testing.T) {
	ns := []NullComplex{
		NComplex(0),
		NComplex(1),
		{Error: ErrDefaultValue},
	}
	sl := NullComplexSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, false), slice length not equal", ns)
	}
	sl = NullComplexSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullComplex64Slice(t *testing.T) {
	ns := []NullComplex64{
		NComplex64(0),
		NComplex64(1),
		{Error: ErrDefaultValue},
	}
	sl := NullComplex64Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplex64Slice(%v, false), slice length not equal", ns)
	}
	sl = NullComplex64Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplex64Slice(%v, true), slice length not equal", ns)
	}
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
		testSet(t, nv)
	}
	testNType(t, &NullInt{}, NInt)
	testNType(t, &NullInt64{}, NInt64)
	testNType(t, &NullInt32{}, NInt32)
	testNType(t, &NullInt16{}, NInt16)
	testNType(t, &NullInt8{}, NInt8)
}

func TestNullIntSlice(t *testing.T) {
	ns := []NullInt{
		NInt(0),
		NInt(1),
		{Error: ErrDefaultValue},
	}
	sl := NullIntSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, false), slice length not equal", ns)
	}
	sl = NullIntSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt64Slice(t *testing.T) {
	ns := []NullInt64{
		NInt64(0),
		NInt64(1),
		{Error: ErrDefaultValue},
	}
	sl := NullInt64Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt64Slice(%v, false), slice length not equal", ns)
	}
	sl = NullInt64Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt64Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt32Slice(t *testing.T) {
	ns := []NullInt32{
		NInt32(0),
		NInt32(1),
		{Error: ErrDefaultValue},
	}
	sl := NullInt32Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt32Slice(%v, false), slice length not equal", ns)
	}
	sl = NullInt32Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt32Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt16Slice(t *testing.T) {
	ns := []NullInt16{
		NInt16(0),
		NInt16(1),
		{Error: ErrDefaultValue},
	}
	sl := NullInt16Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt16Slice(%v, false), slice length not equal", ns)
	}
	sl = NullInt16Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt16Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullInt8Slice(t *testing.T) {
	ns := []NullInt8{
		NInt8(0),
		NInt8(1),
		{Error: ErrDefaultValue},
	}
	sl := NullInt8Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInt8Slice(%v, false), slice length not equal", ns)
	}
	sl = NullInt8Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInt8Slice(%v, true), slice length not equal", ns)
	}
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
		testSet(t, nv)
	}
	testNType(t, &NullUint{}, NUint)
	testNType(t, &NullUint64{}, NUint64)
	testNType(t, &NullUint32{}, NUint32)
	testNType(t, &NullUint16{}, NUint16)
	testNType(t, &NullUint8{}, NUint8)
}

func TestNullUintSlice(t *testing.T) {
	ns := []NullUint{
		NUint(0),
		NUint(1),
		{Error: ErrDefaultValue},
	}
	sl := NullUintSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, false), slice length not equal", ns)
	}
	sl = NullUintSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullComplexSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint64Slice(t *testing.T) {
	ns := []NullUint64{
		NUint64(0),
		NUint64(1),
		{Error: ErrDefaultValue},
	}
	sl := NullUint64Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint64Slice(%v, false), slice length not equal", ns)
	}
	sl = NullUint64Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint64Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint32Slice(t *testing.T) {
	ns := []NullUint32{
		NUint32(0),
		NUint32(1),
		{Error: ErrDefaultValue},
	}
	sl := NullUint32Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint32Slice(%v, false), slice length not equal", ns)
	}
	sl = NullUint32Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint32Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint16Slice(t *testing.T) {
	ns := []NullUint16{
		NUint16(0),
		NUint16(1),
		{Error: ErrDefaultValue},
	}
	sl := NullUint16Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint16Slice(%v, false), slice length not equal", ns)
	}
	sl = NullUint16Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint16Slice(%v, true), slice length not equal", ns)
	}
}

func TestNullUint8Slice(t *testing.T) {
	ns := []NullUint8{
		NUint8(0),
		NUint8(1),
		{Error: ErrDefaultValue},
	}
	sl := NullUint8Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullUint8Slice(%v, false), slice length not equal", ns)
	}
	sl = NullUint8Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullUint8Slice(%v, true), slice length not equal", ns)
	}
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
		testSet(t, nv)
	}
	testNType(t, &NullFloat32{}, NFloat32)
	testNType(t, &NullFloat{}, NFloat)
}

func TestNullFloatSlice(t *testing.T) {
	ns := []NullFloat{
		NFloat(0),
		NFloat(1),
		{Error: ErrDefaultValue},
	}
	sl := NullFloatSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullFloatSlice(%v, false), slice length not equal", ns)
	}
	sl = NullFloatSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullFloatSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullFloat32Slice(t *testing.T) {
	ns := []NullFloat32{
		NFloat32(0),
		NFloat32(1),
		{Error: ErrDefaultValue},
	}
	sl := NullFloat32Slice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullFloat32Slice(%v, false), slice length not equal", ns)
	}
	sl = NullFloat32Slice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullFloat32Slice(%v, true), slice length not equal", ns)
	}
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
		testNType(t, nv, NString)
	}
}

func TestNullStringSlice(t *testing.T) {
	ns := []NullString{
		NString("t"),
		NString("f"),
		{Error: ErrDefaultValue},
	}
	sl := NullStringSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullStringSlice(%v, false), slice length not equal", ns)
	}
	sl = NullStringSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullStringSlice(%v, true), slice length not equal", ns)
	}
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
		testNType(t, nv, NInterface)
	}
}

func TestNullInterfaceSlice(t *testing.T) {
	ns := []NullInterface{
		NInterface(0),
		NInterface(1),
		{Error: ErrDefaultValue},
	}
	sl := NullInterfaceSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullInterfaceSlice(%v, false), slice length not equal", ns)
	}
	sl = NullInterfaceSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullInterfaceSlice(%v, true), slice length not equal", ns)
	}
}

func TestNullTime(t *testing.T) {
	for _, nv := range []interface{}{
		&NullTime{},
	} {
		testMarshalJSON(t, nv)
		testUnmarshalJSON(t, nv)
		testScanSQL(t, nv)
		testValueSQL(t, nv)
		testSet(t, nv)
		testNType(t, nv, NTime)
	}
}

func TestNullTimeSlice(t *testing.T) {
	ns := []NullTime{
		NTime(time.Now()),
		NTime(time.Now()),
		{Error: ErrDefaultValue},
	}
	sl := NullTimeSlice(ns, false)
	if len(sl) != len(ns) || cap(sl) != cap(ns) {
		t.Errorf("NullTimeSlice(%v, false), slice length not equal", ns)
	}
	sl = NullTimeSlice(ns, true)
	if len(sl) != len(ns)-1 || cap(sl) != cap(ns) {
		t.Errorf("NullTimeSlice(%v, true), slice length not equal", ns)
	}
}

func testScanSQL(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(SQLValueType{}))
	testSuite := func(sv SQLValueType, cnv interface{}, valid bool, err error) {
		aErr := nv.(sql.Scanner).Scan(sv.SQLValue)
		expected := testGetNullIfaceValue(cnv)
		actual := testGetNullIfaceValue(nv)
		if !valid {
			return
		}
		if !matrixSuite.Compare(actual.value, expected.value) || actual.valid != expected.valid || actual.err != expected.err {
			t.Errorf("%T{}.Scan(%T([%[2]v])) failed, expected value by reference %s",
				nv, sv.SQLValue,
				errNull{
					expected.value, expected.valid, expected.err,
					actual.value, actual.valid, aErr,
				})
		}
	}
	for _, di := range testData {
		sv := di.value.Interface().(SQLValueType)
		cnv, valid, _ := matrixSuite.Convert(sv, reflect.TypeOf(nv))
		testSuite(sv, cnv, valid, nil)
	}
}

func testValueSQL(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		v := di.value.Interface()
		eValue, valid, _ := matrixSuite.Convert(v, sqlValueReflectType)
		sv := eValue.(SQLValueType)
		actualValue, actualErr := v.(driver.Valuer).Value()
		if !valid {
			if _, ok := actualErr.(ErrorConvert); !ok {
				t.Errorf("%T{%+[1]v}.Value() must returns 'ErrorConvert' error instead of '%T'", v, actualErr)
			}
		} else {
			from := testGetNullIfaceValue(v)
			if from.present && !matrixSuite.Compare(actualValue, sv.SQLValue) && actualErr == nil {
				t.Errorf("%T{%+[1]v}.Value() failed, expected (expected == actual) %v == %v, error %v",
					v, sv.SQLValue, actualValue, actualErr,
				)
			}
		}
	}
}

func testTyp(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		rv := reflect.ValueOf(di.value.Interface())
		rMethod := rv.MethodByName("Typ")
		res := rMethod.Call([]reflect.Value{})[0].MethodByName("Kind").Call([]reflect.Value{})
		expected := testGetNullIfaceValue(di.value.Interface())
		if !expected.valid {
			expected.nkind = reflect.Invalid
		}
		if res[0].Interface().(reflect.Kind) != expected.nkind {
			t.Errorf("%T{%+[1]v}.Typ().Kind() failed, expected (expected == actual) %s == %s", di.value.Interface(), expected.nkind, res[0].Interface())
		}
		// typ error
		fl := rv.Elem().FieldByName("Error")
		fl.Set(reflect.ValueOf(errPassed))
		res = rMethod.Call([]reflect.Value{})[0].MethodByName("Kind").Call([]reflect.Value{})
		if res[0].Interface().(reflect.Kind) != reflect.Invalid {
			t.Errorf("%T{%+[1]v}.Typ().Kind() failed, expected (expected == actual) %s == %s", di.value.Interface(), expected.nkind, res[0].Interface())
		}
		res = rMethod.Call([]reflect.Value{})[0].MethodByName("Error").Call([]reflect.Value{})
		if res[0].Interface().(error) != errPassed {
			t.Errorf("%T{%+[1]v}.Typ().Error() failed, expected (expected == actual) %v == %v", di.value.Interface(), errPassed, res[0].Interface())
		}
	}
}

func testSet(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		rd := reflect.ValueOf(nv)
		from := testGetNullIfaceValue(reflect.ValueOf(di.value.Interface()).Interface())
		rv := reflect.ValueOf(from.value)
		_, ok := from.value.(NullInterface)
		if from.value == nil || ok || rv.Kind() == reflect.Func {
			continue
		}
		rd.MethodByName("Set").Call([]reflect.Value{rv})
		actual := testGetNullIfaceValue(rd.Interface())
		if !matrixSuite.CompareSafe(actual.value, from.value, true) {
			t.Errorf("%T{}.Set(%T([%[2]v])) failed, expected value from V() %s",
				nv, from.value,
				errNull{
					from.value, false, nil,
					actual.value, actual.valid, actual.err,
				})
		}
	}
}

func testMarshalJSON(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		v := di.value.Interface()
		aValue, vv, _ := matrixSuite.Convert(v, jsonTokenReflectType)
		if aValue == nil || !vv {
			continue
		}
		jt := aValue.(JSONToken)
		actualBytes, actualErr := json.Marshal(jt.Value)
		if !bytes.Equal(actualBytes, jt.Token) && actualErr == nil {
			t.Errorf("%T{%+[1]v}.MarshalJSON() failed, expected (expected == actual) []byte (%s == %s), error %v",
				v, jt.Token, actualBytes, actualErr,
			)
		}
	}
}

func testUnmarshalJSON(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(JSONToken{}))
	for _, di := range testData {
		jt := di.value.Interface().(JSONToken)
		cnv, valid, _ := matrixSuite.Convert(jt, reflect.TypeOf(nv))
		aErr := nv.(json.Unmarshaler).UnmarshalJSON(jt.Token)
		expected := testGetNullIfaceValue(cnv)
		actual := testGetNullIfaceValue(nv)
		if !valid {
			continue
		}
		if !matrixSuite.Compare(actual.value, expected.value) || actual.valid != expected.valid || aErr != expected.err {
			t.Errorf("%T{}.UnmarshalJSON([]byte(%s)) failed, expected value by reference %s",
				nv, jt.Token,
				errNull{
					expected.value, expected.valid, expected.err,
					actual.value, actual.valid, aErr,
				})
		}
	}
}

func testNType(t *testing.T, nv interface{}, fn interface{}) {
	rf := reflect.ValueOf(fn)
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		from := testGetNullIfaceValue(reflect.ValueOf(di.value.Interface()).Interface())
		_, ok := from.value.(NullInterface)
		if from.value == nil || ok || from.nkind == reflect.Func {
			continue
		}
		res := rf.Call([]reflect.Value{reflect.ValueOf(from.value)})
		expected := testGetNullIfaceValue(nv)
		actual := testGetNullIfaceValue(res[0].Interface())
		if !matrixSuite.CompareSafe(actual.value, from.value, true) {
			t.Errorf("%s(%v), %s",
				rf.String(), from.value,
				errNull{
					expected.value, expected.valid, expected.err,
					actual.value, actual.valid, actual.err,
				})
		}
	}
}

// TODO: Benchmark
