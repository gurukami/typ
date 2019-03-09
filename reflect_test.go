package typ

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

var interfaceReflectTypes = []reflect.Type{
	nil,
}

func init() {
	// Test Data
	emptyCh := make(chan struct{})
	chBuff := make(chan struct{}, 1)
	chBuff <- struct{}{}
	closedCh := make(chan struct{})
	close(closedCh)
	matrixSuite.Register(reflect.TypeOf(emptyCh), []dataItem{
		{reflect.ValueOf(emptyCh), nil},
		{reflect.ValueOf(chBuff), nil},
		{reflect.ValueOf(closedCh), nil},
	})
	var pv interface{} = nil
	ptr := &pv
	up := unsafe.Pointer(ptr)
	uptr := uintptr(up)
	matrixSuite.Register(reflect.TypeOf(up), []dataItem{
		{reflect.ValueOf(up), nil},
	})
	matrixSuite.Register(reflect.TypeOf(uptr), []dataItem{
		{reflect.ValueOf(uptr), nil},
	})
	matrixSuite.Register(reflect.TypeOf(struct{}{}), []dataItem{
		{reflect.ValueOf(struct{}{}), nil},
	})
	matrixSuite.Register(reflect.TypeOf([1]interface{}{}), []dataItem{
		{reflect.ValueOf([1]interface{}{}), nil},
		{reflect.ValueOf([1]interface{}{nil}), nil},
	})
	matrixSuite.Register(reflect.TypeOf([]interface{}{}), []dataItem{
		{reflect.ValueOf([]interface{}{}), nil},
		{reflect.ValueOf([]interface{}{nil}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(map[interface{}]interface{}{}), []dataItem{
		{reflect.ValueOf(map[interface{}]interface{}{}), nil},
		{reflect.ValueOf(map[interface{}]interface{}{nil: nil}), nil},
	})
	matrixSuite.Register(reflect.TypeOf(func() {}), []dataItem{
		{reflect.ValueOf(func() {}), nil},
	})
	// Converters
	defaultNil := func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		return nil, false
	}
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Chan), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Uintptr), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.UnsafePointer), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Struct), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Array), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Func), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Slice), defaultNil)
	matrixSuite.SetConverter(nil, getDefaultType(reflect.Map), defaultNil)
	matrixSuite.SetConverters([]reflect.Type{
		getDefaultType(reflect.Chan),
		getDefaultType(reflect.Array),
		getDefaultType(reflect.Map),
		getDefaultType(reflect.Slice),
	}, boolReflectTypes, func(from interface{}, to reflect.Type, opts ...interface{}) (interface{}, bool) {
		rv := reflect.ValueOf(from)
		return rv.Len() > 0, true
	})
}

const magicNumber = 42

var (
	pointerv      = true
	pointer       = &pointerv
	unsafePointer = unsafe.Pointer(pointer)
	dTypes        = map[reflect.Kind]interface{}{
		reflect.Invalid:       nil,
		reflect.Bool:          false,
		reflect.Int:           int(0),
		reflect.Int8:          int8(0),
		reflect.Int16:         int16(0),
		reflect.Int32:         int32(0),
		reflect.Int64:         int64(0),
		reflect.Uint:          uint(0),
		reflect.Uint8:         uint8(0),
		reflect.Uint16:        uint16(0),
		reflect.Uint32:        uint32(0),
		reflect.Uint64:        uint64(0),
		reflect.Float32:       float32(0),
		reflect.Float64:       float64(0),
		reflect.Complex64:     complex64(0),
		reflect.Complex128:    complex128(0),
		reflect.Uintptr:       uintptr(unsafePointer),
		reflect.Struct:        struct{}{},
		reflect.Array:         [1]int{},
		reflect.Chan:          make(chan struct{}),
		reflect.Func:          func() {},
		reflect.Interface:     (interface{})(struct{}{}),
		reflect.Ptr:           pointer,
		reflect.Slice:         []int{},
		reflect.String:        "",
		reflect.UnsafePointer: unsafePointer,
		reflect.Map:           make(map[int]int),
	}
)

type errNull struct {
	EValue interface{}
	EValid bool
	EError error
	AValue interface{}
	AValid bool
	AError error
}

func (e errNull) String() string {
	return fmt.Sprintf("(expected == actual) | Value (%T(%[1]v) == %T(%[2]v)), Valid (%T(%[3]v) == %T(%[4]v)), Error (%T(%[5]v) == %T(%[6]v))",
		e.EValue, e.AValue,
		e.EValid, e.AValid,
		e.EError, e.AError,
	)
}

func getDefaultType(kind reflect.Kind) reflect.Type {
	return reflect.TypeOf(dTypes[kind])
}

func rType(v interface{}) reflect.Kind {
	kind := reflect.Invalid
	typeOf := reflect.TypeOf(v)
	if typeOf != nil {
		kind = typeOf.Kind()
	}
	return kind
}

func rFnCall(fn interface{}, args []interface{}) []interface{} {
	rFn := reflect.ValueOf(fn)
	var fnArgs, fnRes []reflect.Value
	for _, arg := range args {
		fnArgs = append(fnArgs, reflect.ValueOf(arg))
	}
	fnRes = rFn.Call(fnArgs)
	var res []interface{}
	for _, ret := range fnRes {
		res = append(res, ret.Interface())
	}
	return res
}

func TestIsSafeFloat(t *testing.T) {
	floatBitSize := []int{32, 64}
	testData := [][]interface{}{
		{
			float64(0),
			[]bool{true, true},
		},
		{
			float64(MinFloat32),
			[]bool{true, true},
		},
		{
			MinFloat64,
			[]bool{false, true},
		},
		{
			float64(MaxFloat32),
			[]bool{true, true},
		},
		{
			MaxFloat64,
			[]bool{false, true},
		},
		{
			float64(MinSafeIntFloat32),
			[]bool{true, true},
		},
		{
			MinSafeIntFloat64,
			[]bool{false, true},
		},
		{
			float64(MaxSafeIntFloat32),
			[]bool{true, true},
		},
		{
			MaxSafeIntFloat64,
			[]bool{false, true},
		},
		{
			math.NaN(),
			[]bool{true, true},
		},
		{
			math.Inf(1),
			[]bool{true, true},
		},
		{
			math.Inf(-1),
			[]bool{true, true},
		},
		{
			float64(math.SmallestNonzeroFloat32),
			[]bool{true, true},
		},
		{
			float64(math.SmallestNonzeroFloat64),
			[]bool{false, true},
		},
	}
	for _, v := range testData {
		for i, bitSize := range floatBitSize {
			from := v[0].(float64)
			expects := v[1].([]bool)[i]
			if isSafeFloat(from, bitSize) != expects {
				t.Errorf("isSafeFloat(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeFloatToInt(t *testing.T) {
	intBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		{
			float64(0),
			[]bool{true, true, true, true},
			32,
		},
		{
			float64(MinFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			MinFloat64,
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(MaxFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			MaxFloat64,
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(MinSafeIntFloat32),
			[]bool{false, false, true, true},
			32,
		},
		{
			MinSafeIntFloat64,
			[]bool{false, false, false, true},
			64,
		},
		{
			float64(MaxSafeIntFloat32),
			[]bool{false, false, true, true},
			32,
		},
		{
			MaxSafeIntFloat64,
			[]bool{false, false, false, true},
			64,
		},
		{
			math.NaN(),
			[]bool{false, false, false, false},
			64,
		},
		{
			math.Inf(1),
			[]bool{false, false, false, false},
			64,
		},
		{
			math.Inf(-1),
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(math.SmallestNonzeroFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			float64(math.SmallestNonzeroFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(1.1),
			[]bool{false, false, false, false},
			64,
		},
	}
	for _, v := range testData {
		for i, bitSize := range intBitSize {
			from := v[0].(float64)
			expects := v[1].([]bool)[i]
			floatBitSize := v[2].(int)
			if isSafeFloatToInt(from, floatBitSize, bitSize) != expects {
				t.Errorf("isSafeFloatToInt(%v, %v, %v) failed, expects %v", from, floatBitSize, bitSize, expects)
			}
		}
	}
}

func TestIsSafeFloatToUint(t *testing.T) {
	uintBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		{
			float64(0),
			[]bool{true, true, true, true},
			32,
		},
		{
			float64(MinFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			MinFloat64,
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(MaxFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			MaxFloat64,
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(MinSafeIntFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			MinSafeIntFloat64,
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(MaxSafeIntFloat32),
			[]bool{false, false, true, true},
			32,
		},
		{
			MaxSafeIntFloat64,
			[]bool{false, false, false, true},
			64,
		},
		{
			math.NaN(),
			[]bool{false, false, false, false},
			64,
		},
		{
			math.Inf(1),
			[]bool{false, false, false, false},
			64,
		},
		{
			math.Inf(-1),
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(math.SmallestNonzeroFloat32),
			[]bool{false, false, false, false},
			32,
		},
		{
			float64(math.SmallestNonzeroFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			float64(1.1),
			[]bool{false, false, false, false},
			64,
		},
	}
	for _, v := range testData {
		for i, bitSize := range uintBitSize {
			from := v[0].(float64)
			expects := v[1].([]bool)[i]
			floatBitSize := v[2].(int)
			if isSafeFloatToUint(from, floatBitSize, bitSize) != expects {
				t.Errorf("isSafeFloatToUint(%v, %v, %v) failed, expects %v", from, floatBitSize, bitSize, expects)
			}
		}
	}
}

func TestIsSafeComplex(t *testing.T) {
	floatBitSize := []int{32, 64}
	testData := [][]interface{}{
		{
			complex(float64(0), float64(0)),
			[]bool{true, true},
		},
		{
			complex(float64(MinFloat32), float64(MinFloat32)),
			[]bool{true, true},
		},
		{
			complex(MinFloat64, MinFloat64),
			[]bool{false, true},
		},
		{
			complex(float64(MaxFloat32), float64(MaxFloat32)),
			[]bool{true, true},
		},
		{
			complex(MaxFloat64, MaxFloat64),
			[]bool{false, true},
		},
		{
			complex(float64(MinSafeIntFloat32), float64(MinSafeIntFloat32)),
			[]bool{true, true},
		},
		{
			complex(MinSafeIntFloat64, MinSafeIntFloat64),
			[]bool{false, true},
		},
		{
			complex(float64(MaxSafeIntFloat32), float64(MaxSafeIntFloat32)),
			[]bool{true, true},
		},
		{
			complex(MaxSafeIntFloat64, MaxSafeIntFloat64),
			[]bool{false, true},
		},
		{
			complex(math.NaN(), math.NaN()),
			[]bool{true, true},
		},
		{
			complex(math.Inf(1), math.Inf(1)),
			[]bool{true, true},
		},
		{
			complex(math.Inf(-1), math.Inf(-1)),
			[]bool{true, true},
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), float64(math.SmallestNonzeroFloat32)),
			[]bool{true, true},
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), float64(math.SmallestNonzeroFloat64)),
			[]bool{false, true},
		},
	}
	for _, v := range testData {
		for i, bitSize := range floatBitSize {
			from := v[0].(complex128)
			expects := v[1].([]bool)[i]
			if isSafeComplex(from, bitSize) != expects {
				t.Errorf("isSafeComplex(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeComplexToFloat(t *testing.T) {
	floatBitSize := []int{32, 64}
	testData := [][]interface{}{
		// Without imaginary part
		{
			complex(0, 0),
			[]bool{true, true},
		},
		{
			complex(float64(MinFloat32), 0),
			[]bool{true, true},
		},
		{
			complex(MinFloat64, 0),
			[]bool{false, true},
		},
		{
			complex(float64(MaxFloat32), 0),
			[]bool{true, true},
		},
		{
			complex(MaxFloat64, 0),
			[]bool{false, true},
		},
		{
			complex(float64(MinSafeIntFloat32), 0),
			[]bool{true, true},
		},
		{
			complex(MinSafeIntFloat64, 0),
			[]bool{false, true},
		},
		{
			complex(float64(MaxSafeIntFloat32), 0),
			[]bool{true, true},
		},
		{
			complex(MaxSafeIntFloat64, 0),
			[]bool{false, true},
		},
		{
			complex(math.NaN(), 0),
			[]bool{true, true},
		},
		{
			complex(math.Inf(1), 0),
			[]bool{true, true},
		},
		{
			complex(math.Inf(-1), 0),
			[]bool{true, true},
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), 0),
			[]bool{true, true},
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), 0),
			[]bool{false, true},
		},
		// With imaginary part
		{
			complex(float64(MinFloat32), float64(MinFloat32)),
			[]bool{false, false},
		},
		{
			complex(MinFloat64, MinFloat64),
			[]bool{false, false},
		},
		{
			complex(float64(MaxFloat32), float64(MaxFloat32)),
			[]bool{false, false},
		},
		{
			complex(MaxFloat64, MaxFloat64),
			[]bool{false, false},
		},
		{
			complex(float64(MinSafeIntFloat32), float64(MinSafeIntFloat32)),
			[]bool{false, false},
		},
		{
			complex(MinSafeIntFloat64, MinSafeIntFloat64),
			[]bool{false, false},
		},
		{
			complex(float64(MaxSafeIntFloat32), float64(MaxSafeIntFloat32)),
			[]bool{false, false},
		},
		{
			complex(MaxSafeIntFloat64, MaxSafeIntFloat64),
			[]bool{false, false},
		},
		{
			complex(math.NaN(), math.NaN()),
			[]bool{false, false},
		},
		{
			complex(math.Inf(1), math.Inf(1)),
			[]bool{false, false},
		},
		{
			complex(math.Inf(-1), math.Inf(-1)),
			[]bool{false, false},
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), float64(math.SmallestNonzeroFloat32)),
			[]bool{false, false},
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), float64(math.SmallestNonzeroFloat64)),
			[]bool{false, false},
		},
	}
	for _, v := range testData {
		for i, bitSize := range floatBitSize {
			from := v[0].(complex128)
			expects := v[1].([]bool)[i]
			if isSafeComplexToFloat(from, bitSize) != expects {
				t.Errorf("isSafeComplexToFloat(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeComplexToInt(t *testing.T) {
	intBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		// Without imaginary part
		{
			complex(float64(0), 0),
			[]bool{true, true, true, true},
			32,
		},
		{
			complex(float64(MinFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MinFloat64, 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MaxFloat64, 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MinSafeIntFloat32), 0),
			[]bool{false, false, true, true},
			32,
		},
		{
			complex(MinSafeIntFloat64, 0),
			[]bool{false, false, false, true},
			64,
		},
		{
			complex(float64(MaxSafeIntFloat32), 0),
			[]bool{false, false, true, true},
			32,
		},
		{
			complex(MaxSafeIntFloat64, 0),
			[]bool{false, false, false, true},
			64,
		},
		{
			complex(math.NaN(), 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(1), 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(-1), 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), 0),
			[]bool{false, false, false, false},
			64,
		},
		// With imaginary part
		{
			complex(float64(MinFloat32), float64(MinFloat32)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(MinFloat64, MinFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxFloat32), float64(MaxFloat32)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MaxFloat64, MaxFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MinSafeIntFloat32), float64(MinSafeIntFloat32)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(MinSafeIntFloat64, MinSafeIntFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxSafeIntFloat32), float64(MaxSafeIntFloat32)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MaxSafeIntFloat64, MaxSafeIntFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.NaN(), math.NaN()),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(1), math.Inf(1)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(-1), math.Inf(-1)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), float64(math.SmallestNonzeroFloat32)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), float64(math.SmallestNonzeroFloat64)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(1.1), float64(0)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(1.1), float64(0)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(1.1), float64(1.1)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(1.1), float64(1.1)),
			[]bool{false, false, false, false},
			64,
		},
	}
	for _, v := range testData {
		for i, bitSize := range intBitSize {
			from := v[0].(complex128)
			expects := v[1].([]bool)[i]
			floatBitSize := v[2].(int)
			if isSafeComplexToInt(from, floatBitSize, bitSize) != expects {
				t.Errorf("isSafeComplexToInt(%v, %v, %v) failed, expects %v", from, floatBitSize, bitSize, expects)
			}
		}
	}
}

func TestIsSafeComplexToUint(t *testing.T) {
	uintBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		// Without imaginary part
		{
			complex(float64(0), 0),
			[]bool{true, true, true, true},
			32,
		},
		{
			complex(float64(MinFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MinFloat64, 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MaxFloat64, 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MinSafeIntFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MinSafeIntFloat64, 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxSafeIntFloat32), 0),
			[]bool{false, false, true, true},
			32,
		},
		{
			complex(MaxSafeIntFloat64, 0),
			[]bool{false, false, false, true},
			64,
		},
		{
			complex(math.NaN(), 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(1), 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(-1), 0),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), 0),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), 0),
			[]bool{false, false, false, false},
			64,
		},
		// With imaginary part
		{
			complex(float64(MinFloat32), float64(MinFloat32)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(MinFloat64, MinFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxFloat32), float64(MaxFloat32)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MaxFloat64, MaxFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MinSafeIntFloat32), float64(MinSafeIntFloat32)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(MinSafeIntFloat64, MinSafeIntFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(MaxSafeIntFloat32), float64(MaxSafeIntFloat32)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(MaxSafeIntFloat64, MaxSafeIntFloat64),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.NaN(), math.NaN()),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(1), math.Inf(1)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(math.Inf(-1), math.Inf(-1)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(math.SmallestNonzeroFloat32), float64(math.SmallestNonzeroFloat32)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(math.SmallestNonzeroFloat64), float64(math.SmallestNonzeroFloat64)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(1.1), float64(0)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(1.1), float64(0)),
			[]bool{false, false, false, false},
			64,
		},
		{
			complex(float64(1.1), float64(1.1)),
			[]bool{false, false, false, false},
			32,
		},
		{
			complex(float64(1.1), float64(1.1)),
			[]bool{false, false, false, false},
			64,
		},
	}
	for _, v := range testData {
		for i, bitSize := range uintBitSize {
			from := v[0].(complex128)
			expects := v[1].([]bool)[i]
			floatBitSize := v[2].(int)
			if isSafeComplexToUint(from, floatBitSize, bitSize) != expects {
				t.Errorf("isSafeComplexToUint(%v, %v, %v) failed, expects %v", from, floatBitSize, bitSize, expects)
			}
		}
	}
}

func TestIsSafeInt(t *testing.T) {
	intBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		{
			int64(0),
			[]bool{true, true, true, true},
		},
		{
			int64(MinInt8),
			[]bool{true, true, true, true},
		},
		{
			int64(MinInt16),
			[]bool{false, true, true, true},
		},
		{
			int64(MinInt32),
			[]bool{false, false, true, true},
		},
		{
			int64(MinInt64),
			[]bool{false, false, false, true},
		},
		{
			int64(MaxInt8),
			[]bool{true, true, true, true},
		},
		{
			int64(MaxInt16),
			[]bool{false, true, true, true},
		},
		{
			int64(MaxInt32),
			[]bool{false, false, true, true},
		},
		{
			int64(MaxInt64),
			[]bool{false, false, false, true},
		},
	}
	for _, v := range testData {
		for i, bitSize := range intBitSize {
			from := v[0].(int64)
			expects := v[1].([]bool)[i]
			if isSafeInt(from, bitSize) != expects {
				t.Errorf("isSafeInt(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeIntToFloat(t *testing.T) {
	floatBitSize := []int{32, 64}
	testData := [][]interface{}{
		{
			int64(0),
			[]bool{true, true},
		},
		{
			int64(MinInt8),
			[]bool{true, true},
		},
		{
			int64(MinInt16),
			[]bool{true, true},
		},
		{
			int64(MinInt32),
			[]bool{false, true},
		},
		{
			int64(MinInt64),
			[]bool{false, false},
		},
		{
			int64(MaxInt8),
			[]bool{true, true},
		},
		{
			int64(MaxInt16),
			[]bool{true, true},
		},
		{
			int64(MaxInt32),
			[]bool{false, true},
		},
		{
			int64(MaxInt64),
			[]bool{false, false},
		},
	}
	for _, v := range testData {
		for i, bitSize := range floatBitSize {
			from := v[0].(int64)
			expects := v[1].([]bool)[i]
			if isSafeIntToFloat(from, bitSize) != expects {
				t.Errorf("isSafeIntToFloat(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeIntToUint(t *testing.T) {
	uintBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		{
			int64(0),
			[]bool{true, true, true, true},
		},
		{
			int64(MinInt8),
			[]bool{false, false, false, false},
		},
		{
			int64(MinInt16),
			[]bool{false, false, false, false},
		},
		{
			int64(MinInt32),
			[]bool{false, false, false, false},
		},
		{
			int64(MinInt64),
			[]bool{false, false, false, false},
		},
		{
			int64(MaxInt8),
			[]bool{true, true, true, true},
		},
		{
			int64(MaxInt16),
			[]bool{false, true, true, true},
		},
		{
			int64(MaxInt32),
			[]bool{false, false, true, true},
		},
		{
			int64(MaxInt64),
			[]bool{false, false, false, true},
		},
	}
	for _, v := range testData {
		for i, bitSize := range uintBitSize {
			from := v[0].(int64)
			expects := v[1].([]bool)[i]
			if isSafeIntToUint(from, bitSize) != expects {
				t.Errorf("isSafeIntToUint(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeUint(t *testing.T) {
	uintBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		{
			uint64(0),
			[]bool{true, true, true, true},
		},
		{
			uint64(MaxUint8),
			[]bool{true, true, true, true},
		},
		{
			uint64(MaxUint16),
			[]bool{false, true, true, true},
		},
		{
			uint64(MaxUint32),
			[]bool{false, false, true, true},
		},
		{
			MaxUint64,
			[]bool{false, false, false, true},
		},
	}
	for _, v := range testData {
		for i, bitSize := range uintBitSize {
			from := v[0].(uint64)
			expects := v[1].([]bool)[i]
			if isSafeUint(from, bitSize) != expects {
				t.Errorf("isSafeUint(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeUintToFloat(t *testing.T) {
	floatBitSize := []int{32, 64}
	testData := [][]interface{}{
		{
			uint64(0),
			[]bool{true, true},
		},
		{
			uint64(MaxUint8),
			[]bool{true, true},
		},
		{
			uint64(MaxUint16),
			[]bool{true, true},
		},
		{
			uint64(MaxUint32),
			[]bool{false, true},
		},
		{
			MaxUint64,
			[]bool{false, false},
		},
	}
	for _, v := range testData {
		for i, bitSize := range floatBitSize {
			from := v[0].(uint64)
			expects := v[1].([]bool)[i]
			if isSafeUintToFloat(from, bitSize) != expects {
				t.Errorf("isSafeUintToFloat(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsSafeUintToInt(t *testing.T) {
	intBitSize := []int{8, 16, 32, 64}
	testData := [][]interface{}{
		{
			uint64(0),
			[]bool{true, true, true, true},
		},
		{
			uint64(MaxUint8),
			[]bool{false, true, true, true},
		},
		{
			uint64(MaxUint16),
			[]bool{false, false, true, true},
		},
		{
			uint64(MaxUint32),
			[]bool{false, false, false, true},
		},
		{
			MaxUint64,
			[]bool{false, false, false, false},
		},
	}
	for _, v := range testData {
		for i, bitSize := range intBitSize {
			from := v[0].(uint64)
			expects := v[1].([]bool)[i]
			if isSafeUintToInt(from, bitSize) != expects {
				t.Errorf("isSafeUintToInt(%v, %v) failed, expects %v", from, bitSize, expects)
			}
		}
	}
}

func TestIsType(t *testing.T) {
	testData := map[string][]interface{}{
		"primitives": {
			[]reflect.Kind{
				reflect.Bool,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
			},
			func(v interface{}) bool {
				return isPrimitives(rType(v)) && Of(v).IsPrimitives() && Of(&v).IsPrimitives(true)
			},
		},
		"composite": {
			[]reflect.Kind{
				reflect.Array,
				reflect.Slice,
				reflect.Map,
			},
			func(v interface{}) bool {
				return isComposite(rType(v)) && Of(v).IsComposite() && Of(&v).IsComposite(true)
			},
		},
		"int": {
			[]reflect.Kind{
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			},
			func(v interface{}) bool {
				return isInt(rType(v)) && Of(v).IsInt() && Of(&v).IsInt(true)
			},
		},
		"uint": {
			[]reflect.Kind{
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			},
			func(v interface{}) bool {
				return isUint(rType(v)) && Of(v).IsUint() && Of(&v).IsUint(true)
			},
		},
		"float": {
			[]reflect.Kind{
				reflect.Float32, reflect.Float64,
			},
			func(v interface{}) bool {
				return isFloat(rType(v)) && Of(v).IsFloat() && Of(&v).IsFloat(true)
			},
		},
		"complex": {
			[]reflect.Kind{
				reflect.Complex64, reflect.Complex128,
			},
			func(v interface{}) bool {
				return isComplex(rType(v)) && Of(v).IsComplex()
			},
		},
		"numeric": {
			[]reflect.Kind{
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
			},
			func(v interface{}) bool {
				return isNumeric(rType(v)) && Of(v).IsNumeric() && Of(&v).IsNumeric(true)
			},
		},
		"pointer": {
			[]reflect.Kind{
				reflect.Ptr,
				reflect.Uintptr,
				reflect.UnsafePointer,
			},
			func(v interface{}) bool {
				return isPointer(rType(v)) && Of(v).IsPointer()
			},
		},
		"unsafe_pointer": {
			[]reflect.Kind{
				reflect.Uintptr,
				reflect.UnsafePointer,
			},
			func(v interface{}) bool {
				return isPointer(rType(v)) && Of(v).IsPointer() && Of(&v).IsPointer(true)
			},
		},
		"indirect_pointer": {
			[]reflect.Kind{
				reflect.Ptr,
			},
			func(v interface{}) bool {
				return isPointer(rType(v)) && Of(v).IsPointer() && Of(v).IsBool(true)
			},
		},
		"bool": {
			[]reflect.Kind{
				reflect.Bool,
			},
			func(v interface{}) bool {
				return isBool(rType(v)) && Of(v).IsBool() && Of(&v).IsBool(true)
			},
		},
		"string": {
			[]reflect.Kind{
				reflect.String,
			},
			func(v interface{}) bool {
				return isString(rType(v)) && Of(v).IsString() && Of(&v).IsString(true)
			},
		},
		"float32": {
			[]reflect.Kind{
				reflect.Float32,
			},
			func(v interface{}) bool {
				return isFloat32(rType(v)) && Of(v).IsFloat32() && Of(&v).IsFloat32(true)
			},
		},
		"float64": {
			[]reflect.Kind{
				reflect.Float64,
			},
			func(v interface{}) bool {
				return isFloat64(rType(v)) && Of(v).IsFloat64() && Of(&v).IsFloat64(true)
			},
		},
		"complex64": {
			[]reflect.Kind{
				reflect.Complex64,
			},
			func(v interface{}) bool {
				return isComplex64(rType(v)) && Of(v).IsComplex64() && Of(&v).IsComplex64(true)
			},
		},
		"complex128": {
			[]reflect.Kind{
				reflect.Complex128,
			},
			func(v interface{}) bool {
				return isComplex128(rType(v)) && Of(v).IsComplex128() && Of(&v).IsComplex128(true)
			},
		},
	}
	expectIs := func(kind reflect.Kind, successList []reflect.Kind) bool {
		for _, v := range successList {
			if kind == v {
				return true
			}
		}
		return false
	}
	var (
		expects bool
		fnTest  func(v interface{}) bool
	)
	for k, v := range testData {
		successList := v[0].([]reflect.Kind)
		fnTest = v[1].(func(v interface{}) bool)
		for dK, dV := range dTypes {
			expects = expectIs(dK, successList)
			if fnTest(dV) != expects {
				t.Errorf("Is*(%v) as %T type by fn name %v failed, expects %v", dV, dV, k, expects)
			}
			if fnTest((interface{})(dV)) != expects {
				t.Errorf("Is*(%v) under iface as %T type by fn name %v failed, expects %v", dV, dV, k, expects)
			}
		}
	}
}

func TestKind(t *testing.T) {
	testData := [][]interface{}{
		{
			dTypes[reflect.Invalid],
			reflect.Invalid,
		},
		{
			dTypes[reflect.Bool],
			reflect.Bool,
		},
		{
			dTypes[reflect.Int],
			reflect.Int,
		},
		{
			dTypes[reflect.Int8],
			reflect.Int8,
		},
		{
			dTypes[reflect.Int16],
			reflect.Int16,
		},
		{
			dTypes[reflect.Int32],
			reflect.Int32,
		},
		{
			dTypes[reflect.Int64],
			reflect.Int64,
		},
		{
			dTypes[reflect.Uint],
			reflect.Uint,
		},
		{
			dTypes[reflect.Uint8],
			reflect.Uint8,
		},
		{
			dTypes[reflect.Uint16],
			reflect.Uint16,
		},
		{
			dTypes[reflect.Uint32],
			reflect.Uint32,
		},
		{
			dTypes[reflect.Uint64],
			reflect.Uint64,
		},
		{
			dTypes[reflect.Uintptr],
			reflect.Uintptr,
		},
		{
			dTypes[reflect.Float32],
			reflect.Float32,
		},
		{
			dTypes[reflect.Float64],
			reflect.Float64,
		},
		{
			dTypes[reflect.Complex64],
			reflect.Complex64,
		},
		{
			dTypes[reflect.Complex128],
			reflect.Complex128,
		},
		{
			dTypes[reflect.Array],
			reflect.Array,
		},
		{
			dTypes[reflect.Chan],
			reflect.Chan,
		},
		{
			dTypes[reflect.Func],
			reflect.Func,
		},
		{
			dTypes[reflect.Slice],
			reflect.Slice,
		},
		{
			dTypes[reflect.String],
			reflect.String,
		},
		{
			dTypes[reflect.Struct],
			reflect.Struct,
		},
		{
			dTypes[reflect.UnsafePointer],
			reflect.UnsafePointer,
		},
		{
			dTypes[reflect.Map],
			reflect.Map,
		},
		{
			dTypes[reflect.Ptr],
			reflect.Ptr,
		},
	}
	var expects reflect.Kind
	for _, v := range testData {
		expects = v[1].(reflect.Kind)
		if Of(v[0]).Kind() != expects {
			t.Errorf("Of(%T).Kind() failed, expects %v, actual %v", v[0], expects, Of(v[0]).Kind())
		}
		indirKind := Of(&v[0]).Kind(true)
		if indirKind != expects && (expects != reflect.Ptr && indirKind == reflect.Bool) {
			t.Errorf("Of(%T).Kind() via link failed, expects %v", v[0], expects)
		}
	}
}
