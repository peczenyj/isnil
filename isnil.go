package isnil

import "reflect"

// IsNiler public interface.
type IsNiler interface {
	// IsNil reports whether its instance is nil.
	IsNil() bool
}

// IsNil reports whether its argument is nil.
// It first checks using untyped nil, then try convert to a [IsNiler] interface.
// Will use [reflect.ValueOf] as fallback, if the value is nilable via [IsNilable] function.
func IsNil(argument any) bool {
	if argument == nil {
		return true
	}

	if checker, ok := argument.(IsNiler); ok {
		return checker.IsNil()
	}

	if value := reflect.ValueOf(argument); IsNilable(value) {
		return value.IsNil()
	}

	return false
}

// IsNilable reports if it is safe call [reflect.Value.IsNil] method on a reflect value object.
// The argument must be a chan, func, interface, map, pointer, or slice value; if it is not, it returns false.
func IsNilable(obj reflect.Value) bool {
	switch obj.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return true
	default:
		return false
	}
}
