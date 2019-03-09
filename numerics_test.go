package typ

import (
	"math"
	"testing"
)

func init() {
	// Test Data
	// Converters
}

func TestNumericToString(t *testing.T) {
	// TODO: Matrix tests
	testData := [][]interface{}{
		{
			[]interface{}{
				float32(0),
				float32(MaxFloat32),
				float32(math.NaN()),
				float32(math.Inf(-1)),
				float32(math.Inf(1)),
			},
			[]string{"0", "3.4028235e+38", "NaN", "-Inf", "+Inf"},
		},
		{
			[]interface{}{
				float64(0),
				float64(MaxFloat64),
				float64(math.NaN()),
				float64(math.Inf(-1)),
				float64(math.Inf(1)),
			},
			[]string{"0", "1.7976931348623157e+308", "NaN", "-Inf", "+Inf"},
		},
		{
			[]interface{}{uint(0),},
			[]string{"0"},
		},
		{
			[]interface{}{uint8(0),},
			[]string{"0"},
		},
		{
			[]interface{}{uint16(0),},
			[]string{"0"},
		},
		{
			[]interface{}{uint32(0),},
			[]string{"0"},
		},
		{
			[]interface{}{uint64(0),},
			[]string{"0"},
		},
		{
			[]interface{}{int(0),},
			[]string{"0"},
		},
		{
			[]interface{}{int8(0),},
			[]string{"0"},
		},
		{
			[]interface{}{int16(0),},
			[]string{"0"},
		},
		{
			[]interface{}{int32(0),},
			[]string{"0"},
		},
		{
			[]interface{}{int64(0),},
			[]string{"0"},
		},
		{
			[]interface{}{
				complex64(0),
				complex(float32(MaxFloat32), 0),
				complex(float32(math.NaN()), 0),
				complex(float32(math.Inf(-1)), 0),
				complex(float32(math.Inf(1)), 0),
			},
			[]string{"(0+0i)", "(3.4028235e+38+0i)", "(NaN+0i)", "(-Inf+0i)", "(+Inf+0i)"},
		},
		{
			[]interface{}{
				complex128(0),
				complex(float64(MaxFloat64), 0),
				complex(float64(math.NaN()), 0),
				complex(float64(math.Inf(-1)), 0),
				complex(float64(math.Inf(1)), 0),
			},
			[]string{"(0+0i)", "(1.7976931348623157e+308+0i)", "(NaN+0i)", "(-Inf+0i)", "(+Inf+0i)"},
		},
	}
	var expectedValue string
	for _, v := range testData {
		for k, iV := range v[0].([]interface{}) {
			expectedValue = v[1].([]string)[k]
			if NumericToString(iV, 10, 'g', -1) != expectedValue {
				t.Errorf("NumericToString(%v, 10, 'g', -1) as %T type failed, expectedValue %v", iV, iV, expectedValue)
			}
		}
	}
}
