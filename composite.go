package typ

import (
	"reflect"
)

// Get value from composite type, argument values used as index keys
func (t *Type) Get(argIndexes ...interface{}) (typ *Type) {
	if !t.rv.IsValid() {
		return NewType(nil, ErrInvalidArgument)
	}
	var (
		p   = t.rv
		cnt = len(argIndexes)
		i   int
	)
	defer func() {
		if r := recover(); r != nil {
			typ = NewType(nil, ErrInvalidArgument)
		}
	}()
	for ; i < cnt; i++ {
		switch p.Kind() {
		case reflect.Slice, reflect.Array:
			index, ok := argIndexes[i].(int)
			if !ok {
				return NewType(nil, ErrInvalidArgument)
			}
			if index < 0 || index > p.Len()-1 {
				return NewType(nil, ErrOutOfRange)
			}
			if p = p.Index(index); p.Interface() == nil {
				return NewType(nil, ErrOutOfBounds)
			}
			if p.Kind() == reflect.Interface {
				p = p.Elem()
			}
		case reflect.Map:
			if p = p.MapIndex(reflect.ValueOf(argIndexes[i])); !p.IsValid() {
				return NewType(nil, ErrOutOfBounds)
			}
			if p.Kind() == reflect.Interface {
				p = p.Elem()
			}
		default:
			return NewType(nil, ErrUnexpectedValue)
		}
		if i == cnt-1 {
			return NewType(p.Interface(), nil)
		}
	}
	return NewType(nil, ErrInvalidArgument)
}
