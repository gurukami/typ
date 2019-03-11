package typ

import (
	"math"
	"math/big"
	"reflect"
)

var complexFloatMap = map[reflect.Kind]reflect.Kind{
	reflect.Complex64:  reflect.Float32,
	reflect.Complex128: reflect.Float64,
}

var bitSizeMap = map[reflect.Kind]int{
	reflect.Int8:       8,
	reflect.Int16:      16,
	reflect.Int32:      32,
	reflect.Int64:      64,
	reflect.Int:        64,
	reflect.Uint8:      8,
	reflect.Uint16:     16,
	reflect.Uint32:     32,
	reflect.Uint64:     64,
	reflect.Uint:       64,
	reflect.Float32:    32,
	reflect.Float64:    64,
	reflect.Complex64:  32,
	reflect.Complex128: 64,
}

// Determine whether a value can safely convert between float
func isSafeFloat(from float64, floatBitSize int) bool {
	if math.IsInf(from, 0) || math.IsNaN(from) {
		return true
	}
	if floatBitSize <= 32 && from != 0 {
		_, accuracy := big.NewFloat(from).Float32()
		return accuracy == big.Exact
	}
	return true
}

// Determine whether a value can safely convert from float to int
func isSafeFloatToInt(from float64, floatBitSize, intBitSize int) bool {
	if math.IsNaN(from) {
		return false
	}
	_, exp := math.Frexp(math.Abs(from))
	safeFloat := big.NewFloat(from).IsInt() && ((floatBitSize <= 32 && exp <= 24) || (floatBitSize > 32 && exp <= 53))
	return safeFloat && isSafeInt(int64(from), intBitSize)
}

// Determine whether a value can safely convert from float to uint
func isSafeFloatToUint(from float64, floatBitSize, uintBitSize int) bool {
	if from < 0 || math.IsNaN(from) {
		return false
	}
	_, exp := math.Frexp(math.Abs(from))
	safeFloat := big.NewFloat(from).IsInt() && ((floatBitSize <= 32 && exp <= 24) || (floatBitSize > 32 && exp <= 53))
	return safeFloat && isSafeUint(uint64(from), uintBitSize)
}

// Determine whether a value can safely convert between complex
func isSafeComplex(from complex128, floatBitSize int) bool {
	return isSafeFloat(real(from), floatBitSize) && isSafeFloat(imag(from), floatBitSize)
}

// Determine whether a value can safely convert from complex to float
func isSafeComplexToFloat(from complex128, floatBitSize int) bool {
	return isSafeFloat(real(from), floatBitSize) && imag(from) == 0
}

// Determine whether a value can safely convert from complex to int
func isSafeComplexToInt(from complex128, floatBitSize, intBitSize int) bool {
	return isSafeFloatToInt(real(from), floatBitSize, intBitSize) && imag(from) == 0
}

// Determine whether a value can safely convert from complex to uint
func isSafeComplexToUint(from complex128, floatBitSize, uintBitSize int) bool {
	return isSafeFloatToUint(real(from), floatBitSize, uintBitSize) && imag(from) == 0
}

// Determine whether a value can safely convert between int
func isSafeInt(from int64, intBitSize int) bool {
	maxValue := int64(1<<uint(intBitSize-1) - 1)
	return from >= ^maxValue && from <= maxValue
}

// Determine whether a value can safely convert from int to float
func isSafeIntToFloat(from int64, floatBitSize int) bool {
	floatValue := float64(from)
	_, exp := math.Frexp(math.Abs(floatValue))
	return ((floatBitSize <= 32 && exp <= 24) || (floatBitSize > 32 && exp <= 53)) && isSafeInt(from, floatBitSize)
}

// Determine whether a value can safely convert from int to uint
func isSafeIntToUint(from int64, uintBitSize int) bool {
	return from >= 0 && uint64(from) <= uint64(1<<uint(uintBitSize)-1)
}

// Determine whether a value can safely convert between uint
func isSafeUint(from uint64, uintBitSize int) bool {
	maxValue := uint64(1<<uint(uintBitSize) - 1)
	return from <= maxValue
}

// Determine whether a value can safely convert from uint to float
func isSafeUintToFloat(from uint64, floatBitSize int) bool {
	floatValue := float64(from)
	_, exp := math.Frexp(math.Abs(floatValue))
	return ((floatBitSize <= 32 && exp <= 24) || (floatBitSize > 32 && exp <= 53)) && isSafeUint(from, floatBitSize)
}

// Determine whether a value can safely convert from uint to int
func isSafeUintToInt(from uint64, intBitSize int) bool {
	return from <= uint64(1<<uint(intBitSize-1)-1)
}

// Determine whether a reflect type is primitives type (int, uint, float, complex, bool)
func isPrimitives(kind reflect.Kind) bool {
	return isNumeric(kind) || kind == reflect.Bool
}

// Determine whether a reflect type is composite type (array, slice, map)
func isComposite(kind reflect.Kind) bool {
	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// Determine whether a reflect type is signed integer type
func isInt(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

// Determine whether a reflect type is unsigned integer type
func isUint(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

// Determine whether a reflect type is float type
func isFloat(kind reflect.Kind) bool {
	switch kind {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// Determine whether a reflect type is complex type
func isComplex(kind reflect.Kind) bool {
	switch kind {
	case reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

// Determine whether a reflect type is numeric type (int, uint, float, complex)
func isNumeric(kind reflect.Kind) bool {
	return isInt(kind) || isUint(kind) || isFloat(kind) || isComplex(kind)
}

// Determine whether a reflect type is pointer type
func isPointer(kind reflect.Kind) bool {
	switch kind {
	case reflect.Ptr, reflect.Uintptr, reflect.UnsafePointer:
		return true
	}
	return false
}

// Determine whether a value is bool type
func isBool(kind reflect.Kind) bool {
	return kind == reflect.Bool
}

// Determine whether a value is string type
func isString(kind reflect.Kind) bool {
	return kind == reflect.String
}

// Determine whether a value is float32 type
func isFloat32(kind reflect.Kind) bool {
	return kind == reflect.Float32
}

// Determine whether a value is float64 type
func isFloat64(kind reflect.Kind) bool {
	return kind == reflect.Float64
}

// Determine whether a value is complex64 type
func isComplex64(kind reflect.Kind) bool {
	return kind == reflect.Complex64
}

// Determine whether a value is complex128 type
func isComplex128(kind reflect.Kind) bool {
	return kind == reflect.Complex128
}

// IsPrimitives determine whether a value is primitives type (int, uint, float, complex, bool)
// If indirect argument specified, real reflect type returned
func (t *Type) IsPrimitives(indirect ...bool) bool {
	return isPrimitives(t.Kind(indirect...))
}

// IsComposite determine whether a value is composite type (array, slice, map)
// If indirect argument specified, real reflect type returned
func (t *Type) IsComposite(indirect ...bool) bool {
	return isComposite(t.Kind(indirect...))
}

// IsInt determine whether a value is signed integer type
// If indirect argument specified, real reflect type returned
func (t *Type) IsInt(indirect ...bool) bool {
	return isInt(t.Kind(indirect...))
}

// IsUint determine whether a value is unsigned integer type
// If indirect argument specified, real reflect type returned
func (t *Type) IsUint(indirect ...bool) bool {
	return isUint(t.Kind(indirect...))
}

// IsFloat determine whether a value is float type
// If indirect argument specified, real reflect type returned
func (t *Type) IsFloat(indirect ...bool) bool {
	return isFloat(t.Kind(indirect...))
}

// IsComplex determine whether a value is complex type
// If indirect argument specified, real reflect type returned
func (t *Type) IsComplex(indirect ...bool) bool {
	return isComplex(t.Kind(indirect...))
}

// IsNumeric determine whether a value is numeric type (int, uint, float, complex)
// If indirect argument specified, real reflect type returned
func (t *Type) IsNumeric(indirect ...bool) bool {
	return isNumeric(t.Kind(indirect...))
}

// IsPointer determine whether a value is pointer type (Ptr, UnsafePointer, Uintptr)
// If indirect argument specified, real reflect type returned
func (t *Type) IsPointer(indirect ...bool) bool {
	return isPointer(t.Kind(indirect...))
}

// IsBool determine whether a value is boolean type
// If indirect argument specified, real reflect type returned
func (t *Type) IsBool(indirect ...bool) bool {
	return isBool(t.Kind(indirect...))
}

// IsString determine whether a value is string type
// If indirect argument specified, real reflect type returned
func (t *Type) IsString(indirect ...bool) bool {
	return isString(t.Kind(indirect...))
}

// IsFloat32 determine whether a value is float32 type
// If indirect argument specified, real reflect type returned
func (t *Type) IsFloat32(indirect ...bool) bool {
	return isFloat32(t.Kind(indirect...))
}

// IsFloat64 determine whether a value is float64 type
// If indirect argument specified, real reflect type returned
func (t *Type) IsFloat64(indirect ...bool) bool {
	return isFloat64(t.Kind(indirect...))
}

// IsComplex64 determine whether a value is complex64 type
// If indirect argument specified, real reflect type returned
func (t *Type) IsComplex64(indirect ...bool) bool {
	return isComplex64(t.Kind(indirect...))
}

// IsComplex128 determine whether a value is complex128 type
// If indirect argument specified, real reflect type returned
func (t *Type) IsComplex128(indirect ...bool) bool {
	return isComplex128(t.Kind(indirect...))
}

// Kind returns reflect type
// If indirect argument specified, real reflect type returned
func (t *Type) Kind(indirect ...bool) reflect.Kind {
	if len(indirect) > 0 && indirect[0] {
		return t.rv.Kind()
	}
	return t.kind
}
