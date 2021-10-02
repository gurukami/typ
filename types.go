package typ

import (
	"errors"
	"reflect"
	"unsafe"
)

var (
	// ErrBaseInvalid is returned when a base option has invalid range
	ErrBaseInvalid = ErrorInvalidArgument(errors.New("base option must be in range 2 <= base <= 36"))
	// ErrFmtByteInvalid is returned when a fmt byte has invalid value
	ErrFmtByteInvalid = ErrorInvalidArgument(errors.New("fmtByte option must be one of 'b', 'e', 'E', 'f', 'g', 'G'"))
)

type (
	intStrOpts struct {
		defaultValue *string
		base         int
	}
	uintStrOpts struct {
		defaultValue *string
		base         int
	}
	floatStrOpts struct {
		defaultValue *string
		fmtByte      byte
		precision    int
		bitSize      int
	}
	complexStrOpts struct {
		defaultValue *string
	}

	// Option is interface function used as argument value for type conversion configuration
	Option func(*opts) error
	// IntStringOption is interface function used as argument value for type conversion configuration from int to string
	IntStringOption func(*intStrOpts) error
	// UintStringOption is interface function used as argument value for type conversion configuration from uint to string
	UintStringOption func(*uintStrOpts) error
	// FloatStringOption is interface function used as argument value for type conversion configuration from float to string
	FloatStringOption func(*floatStrOpts) error
	// ComplexStringOption is interface function used as argument value for type conversion configuration from complex to string
	ComplexStringOption func(*complexStrOpts) error
)

type opts struct {
	fmtByte                   *byte
	base, precision           *int
	suffix, prefix, delimiter *string
}

// IntStringDefault set default string value for int conversion to string.
func IntStringDefault(value string) IntStringOption {
	return func(t *intStrOpts) error {
		t.defaultValue = &value
		return nil
	}
}

// IntStringBase set base for int conversion to string.
// The base  must be 2 <= base <= 36
// for digit values >= 10.
func IntStringBase(value int) IntStringOption {
	return func(t *intStrOpts) error {
		if value < 2 || value > 36 {
			return ErrBaseInvalid
		}
		t.base = value
		return nil
	}
}

// UintStringDefault set default string value for uint conversion to string.
func UintStringDefault(value string) UintStringOption {
	return func(t *uintStrOpts) error {
		t.defaultValue = &value
		return nil
	}
}

// UintStringBase set base for uint conversion to string.
// The base  must be 2 <= base <= 36
// for digit values >= 10.
func UintStringBase(value int) UintStringOption {
	return func(t *uintStrOpts) error {
		if value < 2 || value > 36 {
			return ErrBaseInvalid
		}
		t.base = value
		return nil
	}
}

// FloatStringDefault set default string value for float conversion to string.
func FloatStringDefault(value string) FloatStringOption {
	return func(t *floatStrOpts) error {
		t.defaultValue = &value
		return nil
	}
}

// FloatStringFmtByte set format for float conversion to string.
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', and 'G' formats.
// For 'e', 'E', and 'f' it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func FloatStringFmtByte(value byte) FloatStringOption {
	return func(t *floatStrOpts) error {
		switch value {
		case 'b', 'e', 'E', 'f', 'g', 'G':
		default:
			return ErrFmtByteInvalid
		}
		t.fmtByte = value
		return nil
	}
}

// FloatStringPrecision set precision for float conversion to string.
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', and 'G' formats.
// For 'e', 'E', and 'f' it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func FloatStringPrecision(value int) FloatStringOption {
	return func(t *floatStrOpts) error {
		t.precision = value
		return nil
	}
}

// FloatStringBitSize set bit size for float conversion to string.
// The bitSize  must be 32 or 64
func FloatStringBitSize(value int) FloatStringOption {
	return func(t *floatStrOpts) error {
		t.bitSize = value
		return nil
	}
}

// ComplexStringDefault set default string value for complex conversion to string.
func ComplexStringDefault(value string) ComplexStringOption {
	return func(t *complexStrOpts) error {
		t.defaultValue = &value
		return nil
	}
}

// Precision set precision for float conversion.
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', and 'G' formats.
// For 'e', 'E', and 'f' it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func Precision(value int) Option {
	return func(t *opts) error {
		t.precision = &value
		return nil
	}
}

// Base set base for int conversion.
// The base  must be 2 <= base <= 36
// for digit values >= 10.
func Base(value int) Option {
	return func(t *opts) error {
		if value < 2 || value > 36 {
			return ErrBaseInvalid
		}
		t.base = &value
		return nil
	}
}

// FmtByte set format for float conversion.
// The format fmt is one of
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise), or
// 'G' ('E' for large exponents, 'f' otherwise).
func FmtByte(value byte) Option {
	return func(t *opts) error {
		switch value {
		case 'b', 'e', 'E', 'f', 'g', 'G':
		default:
			return ErrFmtByteInvalid
		}
		t.fmtByte = &value
		return nil
	}
}

// Suffix set suffix for string manipulation
func Suffix(value string) Option {
	return func(t *opts) error {
		t.suffix = &value
		return nil
	}
}

// Prefix set prefix for string manipulation
func Prefix(value string) Option {
	return func(t *opts) error {
		t.prefix = &value
		return nil
	}
}

// Delimiter set delimiter for string manipulation
func Delimiter(value string) Option {
	return func(t *opts) error {
		t.delimiter = &value
		return nil
	}
}

// Type stores all information about underlying value.
type Type struct {
	rv   reflect.Value
	kind reflect.Kind
	opts opts
	err  error
}

// Convert "value" to any convertible primitive types
func (t *Type) to(typeTo reflect.Kind) InterfaceAccessor {
	nv := &NullInterface{}
	switch {
	case isUint(typeTo):
		v := t.toUint(typeTo)
		nv.P, nv.Error = v.V(), v.Err()
		return nv
	case isInt(typeTo):
		v := t.toInt(typeTo)
		nv.P, nv.Error = v.V(), v.Err()
		return nv
	case isFloat(typeTo):
		v := t.toFloat(typeTo)
		nv.P, nv.Error = v.V(), v.Err()
		return nv
	case isComplex(typeTo):
		v := t.toComplex(typeTo)
		nv.P, nv.Error = v.V(), v.Err()
		return nv
	case typeTo == reflect.Bool:
		v := t.Bool()
		nv.P, nv.Error = v.V(), v.Err()
		return nv
	case typeTo == reflect.String:
		v := t.String()
		nv.P, nv.Error = v.V(), v.Err()
		return nv
	}
	nv.Error = ErrConvert
	return nv
}

// Empty determine whether a variable is zero
func (t *Type) Empty() BoolAccessor {
	nv := &NullBool{}
	from := t.rv.Kind()
	switch {
	case t.IsUint(true):
		v := t.rv.Uint() == 0
		nv.P = &v
		return nv
	case t.IsInt(true):
		v := t.rv.Int() == 0
		nv.P = &v
		return nv
	case t.IsFloat(true):
		v := t.rv.Float() == 0
		nv.P = &v
		return nv
	case t.IsComplex(true):
		v := t.rv.Complex() == 0
		nv.P = &v
		return nv
	case t.IsComposite(true) || from == reflect.Chan || from == reflect.String:
		v := t.rv.Len() == 0
		nv.P = &v
		return nv
	case t.IsBool(true):
		v := !t.rv.Bool()
		nv.P = &v
		return nv
	case from == reflect.Uintptr:
		v := t.rv.Interface().(uintptr) == 0
		nv.P = &v
		return nv
	case from == reflect.UnsafePointer:
		v := uintptr(t.rv.Interface().(unsafe.Pointer)) == 0
		nv.P = &v
		return nv
	}
	v := !t.rv.IsValid()
	nv.P = &v
	if v {
		nv.Error = ErrUnexpectedValue
	}
	return nv
}

// Equals determine whether a variable is equals with current "value" (same value, but can have different primitives types)
// Primitives type is: int, uint, float, complex, bool
func (t *Type) Equals(value interface{}) BoolAccessor {
	if vp := Of(value).to(t.rv.Kind()); vp.Valid() {
		value = vp.V()
	}
	return t.Identical(value)
}

// Identical determine whether a variable is identical with current "value" (same type and same value)
func (t *Type) Identical(src interface{}) BoolAccessor {
	nv := &NullBool{}
	if !t.rv.IsValid() {
		nv.Error = ErrUnexpectedValue
		return nv
	}
	v := reflect.DeepEqual(t.rv.Interface(), src)
	nv.P = &v
	return nv
}

// Interface returns value as interface.
// Returns nil if value can't safely represents as interface
func (t *Type) Interface() InterfaceAccessor {
	nv := &NullInterface{InterfaceCommon{Error: t.err}}
	if !t.rv.IsValid() {
		return nv
	}
	nv.P = t.rv.Interface()
	return nv
}

// OptionBase returns the base for numeric conversion to string
func (t *Type) OptionBase() int {
	return *t.opts.base
}

// OptionFmtByte returns float format option for float conversion to string
func (t *Type) OptionFmtByte() byte {
	return *t.opts.fmtByte
}

// OptionPrecision returns float precision for float conversion to string
func (t *Type) OptionPrecision() int {
	return *t.opts.precision
}

// Error returns Underlying error.
func (t *Type) Error() error {
	return t.err
}

// Of create type converter from interface value.
// This function recursive dereference value by a reference if value is a pointer
func Of(value interface{}, options ...Option) *Type {
	return NewType(value, nil, options...)
}

var (
	dBase      = 10
	dFmtByte   = byte('f')
	dPrecision = -1
)

// NewType create a Type instance with value.
// This function recursive dereference value by a reference if value is a pointer
func NewType(value interface{}, err error, options ...Option) *Type {
	nt := &Type{}
	switch v := value.(type) {
	case *Type:
		nt.rv, nt.kind = v.rv, v.kind
		nt.opts.fmtByte, nt.opts.base, nt.opts.precision = v.opts.fmtByte, v.opts.base, v.opts.precision
		if v.err != nil && err == nil {
			nt.err = v.err
		}
		return nt
	default:
		nt.rv = reflect.ValueOf(value)
		nt.kind = nt.rv.Kind()
		nt.err = err
	next:
		for {
			switch nt.rv.Kind() {
			case reflect.Interface:
				nt.rv = nt.rv.Elem()
			case reflect.Ptr:
				nt.rv = nt.rv.Elem()
			case reflect.UnsafePointer, reflect.Uintptr, reflect.Invalid:
				fallthrough
			default:
				break next
			}
		}
		if len(options) > 0 {
			for _, v := range options {
				if optErr := v(&nt.opts); optErr != nil {
					nt.err = optErr
					break
				}
			}
		}
		if nt.opts.base == nil {
			nt.opts.base = &dBase
		}
		if nt.opts.fmtByte == nil {
			nt.opts.fmtByte = &dFmtByte
		}
		if nt.opts.precision == nil {
			nt.opts.precision = &dPrecision
		}
		return nt
	}
}
