package typ

import "errors"

// ErrorConvert is returned when value can't safely convert
type ErrorConvert error
// ErrorPanic is returned when fatal error is occurred
type ErrorPanic error
// ErrorOutOfRange is returned when value out of range on given data
type ErrorOutOfRange error
// ErrorOutOfBounds is returned when value out of bounds on given data
type ErrorOutOfBounds error
// ErrorInvalidArgument is returned when invalid argument is present
type ErrorInvalidArgument error
// ErrorUnexpectedValue is returned when unexpected value given on data
type ErrorUnexpectedValue error

var (
	// ErrConvert is returned when value can't safely convert
	ErrConvert         = ErrorConvert(errors.New("value can't safely convert"))
	// ErrOutOfRange is returned when value out of range on given data
	ErrOutOfRange      = ErrorOutOfRange(errors.New("out of range on given data"))
	// ErrOutOfBounds is returned when value out of bounds on given data
	ErrOutOfBounds     = ErrorOutOfBounds(errors.New("out of bounds on given data"))
	// ErrUnexpectedValue is returned when unexpected value given on data
	ErrUnexpectedValue = ErrorUnexpectedValue(errors.New("unexpected value given on data"))
	// ErrInvalidArgument is returned when invalid argument is present
	ErrInvalidArgument = ErrorInvalidArgument(errors.New("invalid argument"))
	// ErrDefaultValue is returned when default value is ambiguous
	ErrDefaultValue    = ErrorInvalidArgument(errors.New("default value is ambiguous"))
)
