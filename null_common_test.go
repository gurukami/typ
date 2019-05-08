package typ

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

var (
	sqlValueReflectTypes = []reflect.Type{
		reflect.TypeOf(SQLValueType{}),
	}
	sqlValueReflectType = reflect.TypeOf(SQLValueType{})
)

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

func testClone(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		rv := reflect.ValueOf(di.value.Interface())
		rMethod := rv.MethodByName("Clone")
		parent := rv.MethodByName("V").Call([]reflect.Value{})
		child := rMethod.Call([]reflect.Value{})[0].MethodByName("V").Call([]reflect.Value{})
		var a, b interface{}
		a = parent[0].Interface()
		b = child[0].Interface()
		if parentIface := testGetNullIfaceValue(a); parentIface.nkind != reflect.Invalid {
			a = parentIface.value
		}
		if childIface := testGetNullIfaceValue(b); childIface.nkind != reflect.Invalid {
			b = childIface.value
		}
		if !matrixSuite.CompareSafe(a, b, true) {
			t.Errorf("%T{%+[1]v}.Clone() failed, expected (expected == actual) %v == %v", di.value.Interface(), a, b)
		}
	}
}

func testSet(t *testing.T, nv interface{}) {
	testData := matrixSuite.GenerateToTyp(matrixSuite.Generate(), reflect.TypeOf(nv))
	for _, di := range testData {
		rd := reflect.ValueOf(nv)
		from := testGetNullIfaceValue(reflect.ValueOf(di.value.Interface()).Interface())
		rv := reflect.ValueOf(from.value)
		_, nok := from.value.(NullInterface)
		_, nnok := from.value.(NotNullInterface)
		if from.value == nil || nok || nnok || rv.Kind() == reflect.Func {
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
		_, nok := from.value.(NullInterface)
		_, nnok := from.value.(NotNullInterface)
		if from.value == nil || nok || nnok || from.nkind == reflect.Func {
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
