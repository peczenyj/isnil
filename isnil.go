package isnil

import "reflect"

type isNiler interface {
	IsNil() bool
}

// IsNil reports whether its argument is nil.
func IsNil(argument any) bool {
	if argument == nil {
		return true
	}

	if checker, ok := argument.(isNiler); ok {
		return checker.IsNil()
	}

	value := reflect.ValueOf(argument)

	switch value.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}
