package isnil_test

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
	"unsafe"

	"github.com/peczenyj/isnil"

	"github.com/stretchr/testify/assert"
)

// Example demonstrates the problem isnil solves: typed nil interfaces.
// When a pointer to a type that implements an interface is nil, the interface
// itself is not nil because the interface contains type information.
//nolint:govet,staticcheck // this is intentional - demonstrates the typed nil problem
func Example() {
	var err error = (*fs.PathError)(nil)
	if err == nil {
		fmt.Println("I expected this to be true, but")
	} else {
		fmt.Println("this check fails, since err != nil")
	}

	if isnil.IsNil(err) {
		fmt.Println("the solution is to use isnil package")
	}

	// Output:
	// this check fails, since err != nil
	// the solution is to use isnil package
}

// TestIsNil verifies that IsNil correctly identifies nil values across various types.
func TestIsNil(t *testing.T) {
	t.Parallel()

	// Test cases organized by category for clarity and maintainability
testCases := []struct {
		label    string
		value    any
		expected bool
	}{
		// Untyped and typed nil values
		{"Untyped nil", nil, true},
		{"Typed nil interface (Stringer)", fmt.Stringer(nil), true},
		{"Nil interface (error)", nilPointerOf[error](), true},

		// Zero values (should be false)
		{"Integer zero value: 0", 0, false},
		{"String zero value: empty string", "", false},

		// IsNiler interface implementation
		{"IsNiler interface: nil pointer (should detect as nil)", (*nillable)(nil), true},
		{"IsNiler interface: non-nil pointer", &nillable{}, false},
		{"IsNiler interface: struct value", nillable{}, false},

		// Channels
		{"Channel: nil", nilPointerOf[chan struct{}](), true},
		{"Channel: non-nil (empty channel)", make(chan struct{}), false},

		// Functions
		{"Function: nil", nilPointerOf[func()](), true},
		{"Function: non-nil closure", func() {}, false},

		// Pointers
		{"Pointer: nil *int", nilPointerOf[*int](), true},
		{"Pointer: non-nil *int", pointerOf(1), false},
		{"Pointer: nil **int (pointer to pointer)", nilPointerOf[**int](), true},
		{"Pointer: non-nil **int (pointer to pointer)", pointerOf(pointerOf(1)), false},

		// Unsafe pointers
		{"Unsafe pointer: nil", unsafe.Pointer(nil), true},
		{"Unsafe pointer: to nil *int", unsafe.Pointer(nilPointerOf[*int]()), true},
		{"Unsafe pointer: to valid *int", unsafe.Pointer(pointerOf(1)), false},

		// Interfaces
		{"Interface: non-nil error", assert.AnError, false},

		// Slices
		{"Slice: nil slice", nilPointerOf[[]string](), true},
		{"Slice: non-nil empty slice", []string{}, false},
		{"Slice: non-nil slice with element", []string{""}, false},

		// Edge cases
		{"Custom struct: no IsNiler interface", &struct{}{}, false},
		{"Interface with nil element: non-nil slice containing nil", []any{nil}, false},
	}

	for _, tt := range testCases {
		tt := tt // capture for parallel execution
		t.Run(tt.label, func(t *testing.T) {
			t.Parallel()

			result := isnil.IsNil(tt.value)
			assert.Equalf(t, tt.expected, result,
				"IsNil(%v) should return %v, but got %v",
				tt.value, tt.expected, result)
		})
	}
}

// TestIsNilable verifies that IsNilable correctly identifies types that can be nil.
func TestIsNilable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		label    string
		value    reflect.Value
		expected bool
	}{
		// Non-nilable types
		{"Invalid (nil interface value)", reflect.ValueOf(fmt.Stringer(nil)), false},
		{"Bool", reflect.ValueOf(true), false},
		{"Int", reflect.ValueOf(int(1)), false},
		{"Int8", reflect.ValueOf(int8(1)), false},
		{"Int16", reflect.ValueOf(int16(1)), false},
		{"Int32", reflect.ValueOf(int32(1)), false},
		{"Int64", reflect.ValueOf(int64(1)), false},
		{"Uint", reflect.ValueOf(uint(1)), false},
		{"Uint8", reflect.ValueOf(uint8(1)), false},
		{"Uint16", reflect.ValueOf(uint16(1)), false},
		{"Uint32", reflect.ValueOf(uint32(1)), false},
		{"Uint64", reflect.ValueOf(uint64(1)), false},
		{"Uintptr", reflect.ValueOf(uintptr(unsafe.Pointer(nil))), false},
		{"Float32", reflect.ValueOf(float32(1.0)), false},
		{"Float64", reflect.ValueOf(float64(1.0)), false},
		{"Complex64", reflect.ValueOf(complex64(complex(1, 1))), false},
		{"Complex128", reflect.ValueOf(complex(1, 1)), false},
		{"Array", reflect.ValueOf([1]any{1}), false},
		{"String", reflect.ValueOf("string"), false},
		{"Struct", reflect.ValueOf(struct{}{}), false},

		// Nilable types
		{"Chan", reflect.ValueOf(make(chan struct{})), true},
		{"Func", reflect.ValueOf(func() {}), true},
		{"Interface", reflect.ValueOf(assert.AnError), true},
		{"Map", reflect.ValueOf(map[int]any{}), true},
		{"Pointer", reflect.ValueOf(pointerOf("pointer")), true},
		{"Slice", reflect.ValueOf([]any{}), true},
		{"UnsafePointer", reflect.ValueOf(unsafe.Pointer(nil)), true},
	}

	for _, tt := range testCases {
		tt := tt // capture for parallel execution
		t.Run(tt.label, func(t *testing.T) {
			t.Parallel()

			result := isnil.IsNilable(tt.value)
			assert.Equalf(t, tt.expected, result,
				"IsNilable(%v) should return %v, but got %v",
				tt.value, tt.expected, result)
		})
	}
}

// Helper function to create a nil pointer of any type.
// This is used to test nil values for pointer types.
func nilPointerOf[T any]() *T {
	return nil
}

// Helper function to create a non-nil pointer to a value.
// This is used to ensure we have valid non-nil pointers for testing.
func pointerOf[T any](v T) *T {
	return &v
}

// nillable is a test type that implements the IsNiler interface.
var _ isnil.IsNiler = (*nillable)(nil)

type nillable struct{}

func (n *nillable) IsNil() bool { return n == nil }