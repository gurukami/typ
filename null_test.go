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
		rv := reflect.ValueOf(v)
		b, err := json.Marshal(v)
		if null {
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
		return SQLValueType{from.(time.Time), from}, true
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
	}
}

func testScanSQL(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(SQLValueType{}))
	for _, di := range testData {
		sv := di.value.Interface().(SQLValueType)
		cnv, valid, _ := matrixSuite.Convert(sv, reflect.TypeOf(nv))
		aErr := nv.(sql.Scanner).Scan(sv.SQLValue)
		eValue, _, eValid, _, eErr := testGetNullIfaceValue(cnv)
		aValue, _, aValid, _, _ := testGetNullIfaceValue(nv)
		if !valid {
			continue
		}
		if !matrixSuite.Compare(aValue, eValue) || aValid != eValid || aErr != eErr {
			t.Errorf("%T{}.Scan(%T([%[2]v])) failed, expected value by reference %s",
				nv, sv.SQLValue,
				errNull{
					eValue, eValid, eErr,
					aValue, aValid, aErr,
				})
		}
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
			_, _, _, present, _ := testGetNullIfaceValue(v)
			if present && !matrixSuite.Compare(actualValue, sv.SQLValue) && actualErr == nil {

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
		_, eType, valid, _, _ := testGetNullIfaceValue(di.value.Interface())
		if !valid {
			eType = reflect.Invalid
		}
		if res[0].Interface().(reflect.Kind) != eType {
			t.Errorf("%T{%+[1]v}.Typ().Kind() failed, expected (expected == actual) %s == %s", di.value.Interface(), eType, res[0].Interface())
		}
	}
}

func testSet(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		rd := reflect.ValueOf(nv)
		sValue, _, _, _, _ := testGetNullIfaceValue(reflect.ValueOf(di.value.Interface()).Interface())
		rv := reflect.ValueOf(sValue)
		_, ok := sValue.(NullInterface)
		if sValue == nil || ok || rv.Kind() == reflect.Func {
			continue
		}
		rd.MethodByName("Set").Call([]reflect.Value{rv})
		aValue, _, aValid, _, aErr := testGetNullIfaceValue(rd.Interface())
		if !matrixSuite.CompareSafe(aValue, sValue, true) {
			t.Errorf("%T{}.Set(%T([%[2]v])) failed, expected value from V() %s",
				nv, sValue,
				errNull{
					sValue, false, nil,
					aValue, aValid, aErr,
				})
		}
	}
}

func testMarshalJSON(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		v := di.value.Interface()
		eValue, vv, _ := matrixSuite.Convert(v, jsonTokenReflectType)
		if eValue == nil || !vv {
			continue
		}
		jt := eValue.(JSONToken)
		actualBytes, actualErr := v.(json.Marshaler).MarshalJSON()
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
		eValue, _, eValid, _, eErr := testGetNullIfaceValue(cnv)
		aValue, _, aValid, _, _ := testGetNullIfaceValue(nv)
		if !valid {
			continue
		}
		if !matrixSuite.Compare(aValue, eValue) || aValid != eValid || aErr != eErr {
			t.Errorf("%T{}.UnmarshalJSON([]byte(%s)) failed, expected value by reference %s",
				nv, jt.Token,
				errNull{
					eValue, eValid, eErr,
					aValue, aValid, aErr,
				})
		}
	}
}

// TODO: Benchmark
// TODO: Test N*
// TODO: Test Null*Slice
// TODO: Test Present
