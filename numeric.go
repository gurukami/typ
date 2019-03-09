package typ

import (
	"fmt"
	"strconv"
)

const (
	MinFloat32        = float32(-3.40282346638528859811704183484516925440e+38)
	MaxFloat32        = float32(3.40282346638528859811704183484516925440e+38)
	MinSafeIntFloat32 = float32(-1.6777215e+07)
	MaxSafeIntFloat32 = float32(1.6777215e+07)
	MinFloat64        = float64(-1.797693134862315708145274237317043567981e+308)
	MaxFloat64        = float64(1.797693134862315708145274237317043567981e+308)
	MinSafeIntFloat64 = float64(-9.007199254740991e+15)
	MaxSafeIntFloat64 = float64(9.007199254740991e+15)
	MaxInt8           = int8(127)
	MinInt8           = int8(-128)
	MaxInt16          = int16(32767)
	MinInt16          = int16(-32768)
	MaxInt32          = int32(2147483647)
	MinInt32          = int32(-2147483648)
	MaxInt64          = int64(9223372036854775807)
	MinInt64          = int64(-9223372036854775808)
	MaxInt            = int(9223372036854775807)
	MinInt            = int(-9223372036854775808)
	MaxUint8          = uint8(255)
	MaxUint16         = uint16(65535)
	MaxUint32         = uint32(4294967295)
	MaxUint64         = uint64(18446744073709551615)
	MaxUint           = uint64(18446744073709551615)
)

// Convert from any numeric type to string in the given base
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
// The format fmt is one of
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise), or
// 'G' ('E' for large exponents, 'f' otherwise).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func NumericToString(v interface{}, base int, fmtByte byte, prec int) string {
	var str string
	switch v.(type) {
	case float32:
		str = strconv.FormatFloat(float64(v.(float32)), fmtByte, prec, 32)
	case float64:
		str = strconv.FormatFloat(v.(float64), fmtByte, prec, 64)
	case uint:
		str = strconv.FormatUint(uint64(v.(uint)), base)
	case uint8:
		str = strconv.FormatUint(uint64(v.(uint8)), base)
	case uint16:
		str = strconv.FormatUint(uint64(v.(uint16)), base)
	case uint32:
		str = strconv.FormatUint(uint64(v.(uint32)), base)
	case uint64:
		str = strconv.FormatUint(v.(uint64), base)
	case int:
		str = strconv.FormatInt(int64(v.(int)), base)
	case int8:
		str = strconv.FormatInt(int64(v.(int8)), base)
	case int16:
		str = strconv.FormatInt(int64(v.(int16)), base)
	case int32:
		str = strconv.FormatInt(int64(v.(int32)), base)
	case int64:
		str = strconv.FormatInt(v.(int64), base)
	case complex64:
		str = fmt.Sprintf("%v", v.(complex64))
	case complex128:
		str = fmt.Sprintf("%v", v.(complex128))
	}
	return str
}
