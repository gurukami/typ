package typ

import "errors"

type ErrorConvert error
type ErrorPanic error
type ErrorOutOfRange error
type ErrorOutOfBounds error
type ErrorInvalidArgument error
type ErrorUnexpectedValue error

var (
	ErrConvert         = ErrorConvert(errors.New("value can't safely convert"))
	ErrOutOfRange      = ErrorOutOfRange(errors.New("out of range on given data"))
	ErrOutOfBounds     = ErrorOutOfBounds(errors.New("out of bounds on given data"))
	ErrUnexpectedValue = ErrorUnexpectedValue(errors.New("unexpected value on given data"))
	ErrInvalidArgument = ErrorInvalidArgument(errors.New("invalid argument"))
	ErrDefaultValue    = ErrorInvalidArgument(errors.New("default value is ambiguous"))
)
