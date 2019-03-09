package typ

import (
	"reflect"
	"testing"
)

type (
	TypedArray [6]interface{}
	TypedSlice []interface{}
	StructType struct{}
)

var (
	MapData        map[string]interface{}
	ArrayData      [6]interface{}
	SliceData      []interface{}
	TypedArrayData TypedArray
	TypedSliceData TypedSlice
	StructTypeData StructType
)

func init() {
	// Test Data
	ArrayData = [6]interface{}{
		1, 2, nil,
	}
	SliceData = []interface{}{
		1, 2, nil,
	}
	TypedArrayData = TypedArray{1, 2, nil}
	TypedSliceData = TypedSlice{1, 2, nil}
	StructTypeData = StructType{}
	MapData = map[string]interface{}{
		"one": map[string]string{
			"sub_one": "1",
			"2":       "2",
		},
		"array":       ArrayData,
		"slice":       SliceData,
		"typed_array": TypedArrayData,
		"typed_slice": TypedSliceData,
		"struct":      StructTypeData,
	}
	// Converters
}

func TestCompositeGet(t *testing.T) {
	testData := [][]interface{}{
		{
			// exists
			MapData,
			[]interface{}{"one", "sub_one"},
			true,
			"1",
			nil,
		},
		{
			// invalid, same type
			MapData,
			[]interface{}{"one", "invalid"},
			false,
			nil,
			ErrOutOfBounds,
		},
		{
			// invalid, wrong type
			MapData,
			[]interface{}{"one", 2},
			false,
			nil,
			ErrInvalidArgument,
		},
		{
			// empty args
			MapData,
			[]interface{}{},
			false,
			nil,
			ErrInvalidArgument,
		},
		{
			// invalid data
			nil,
			[]interface{}{},
			false,
			nil,
			ErrInvalidArgument,
		},
		{
			// exists
			ArrayData,
			[]interface{}{1},
			true,
			2,
			nil,
		},
		{
			// invalid, nil Value
			ArrayData,
			[]interface{}{3},
			false,
			nil,
			ErrOutOfBounds,
		},
		{
			// invalid, wrong type
			ArrayData,
			[]interface{}{"1"},
			false,
			nil,
			ErrInvalidArgument,
		}, {
			// exists
			TypedArrayData,
			[]interface{}{1},
			true,
			2,
			nil,
		},
		{
			// invalid, nil Value
			TypedArrayData,
			[]interface{}{3},
			false,
			nil,
			ErrOutOfBounds,
		},
		{
			// invalid, wrong type
			TypedArrayData,
			[]interface{}{"1"},
			false,
			nil,
			ErrInvalidArgument,
		},
		{
			// exists
			SliceData,
			[]interface{}{1},
			true,
			2,
			nil,
		},
		{
			// invalid, out of range
			SliceData,
			[]interface{}{3},
			false,
			nil,
			ErrOutOfRange,
		},
		{
			// invalid, wrong type
			SliceData,
			[]interface{}{"1"},
			false,
			nil,
			ErrInvalidArgument,
		},
		{
			// invalid, nil Value given
			SliceData,
			[]interface{}{2},
			false,
			nil,
			ErrOutOfBounds,
		},
		{
			// exists
			TypedSliceData,
			[]interface{}{1},
			true,
			2,
			"",
		},
		{
			// invalid, out of range
			TypedSliceData,
			[]interface{}{3},
			false,
			nil,
			ErrOutOfRange,
		},
		{
			// invalid, wrong type
			TypedSliceData,
			[]interface{}{"1"},
			false,
			nil,
			ErrInvalidArgument,
		},
		{
			// invalid, nil Value given
			TypedSliceData,
			[]interface{}{2},
			false,
			nil,
			ErrOutOfBounds,
		},
		{
			// struct, wrong type data
			StructTypeData,
			[]interface{}{1},
			false,
			nil,
			ErrUnexpectedValue,
		},
	}
	var (
		test    bool
		expects interface{}
	)
	for _, v := range testData {
		data := v[0]
		args := v[1].([]interface{})
		test = v[2].(bool)
		expects = v[3]
		err, _ := v[4].(error)
		typ := Of(data).Get(args...)
		value := typ.Interface()
		testValid := value.Valid() != test
		testEquals := test && !reflect.DeepEqual(value.V(), expects)
		testError := !test && (value.Error != err)
		if testValid || testEquals || testError {
			t.Errorf("Of(%v[%[1]T]).Get(%v), %s", data, args, errNull{
				expects, value.Valid(), err,
				value.V(), value.Valid(), value.Error,
			})
		}
	}
}

// BenchmarkOfGet-8   	 3000000	       407 ns/op
func BenchmarkOfGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Of(MapData).Get("one", "sub_one")
	}
}
