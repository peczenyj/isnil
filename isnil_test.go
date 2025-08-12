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

//nolint:govet,staticcheck //this is on purpose
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

func TestIsNil(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		label    string
		value    any
		expected bool
	}{
		{"untyped nil", nil, true},
		{"typed nil (interface)", fmt.Stringer(nil), true},
		{"integer (zero value)", 0, false},
		{"string (zero value)", "", false},
		{"isNiler interface, positive case", (*nillable)(nil), true},
		{"isNiler interface, negative case (pointer)", &nillable{}, false},
		{"isNiler interface, negative case (struct)", nillable{}, false},
		{"nil channel", nilPointerOf[chan struct{}](), true},
		{"non-nil channel", make(chan struct{}), false},
		{"nil func", nilPointerOf[func()](), true},
		{"non-nil func", func() {}, false},
		{"nil pointer", nilPointerOf[*int](), true},
		{"non-nil pointer", pointerOf(1), false},
		{"nil pointer to pointer", nilPointerOf[**int](), true},
		{"non-nil pointer to pointer", pointerOf(pointerOf(1)), false},
		{"unsafe pointer to untyped nil", unsafe.Pointer(nil), true},
		{"unsafe pointer to *int nil", unsafe.Pointer(nilPointerOf[*int]()), true},
		{"unsafe pointer to int", unsafe.Pointer(pointerOf(1)), false},
		{"nil interface ", nilPointerOf[error](), true},
		{"non-nil interface ", assert.AnError, false},
		{"nil slice ", nilPointerOf[[]string](), true},
		{"non-nil empty slice ", []string{}, false},
		{"non-nil slice ", []string{""}, false},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, isnil.IsNil(tt.value))
		})
	}
}

func TestIsNilable(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		label    string
		value    reflect.Value
		expected bool
	}{
		{"Invalid", reflect.ValueOf(fmt.Stringer(nil)), false},
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
		{"Chan", reflect.ValueOf(make(chan struct{})), true},
		{"Func", reflect.ValueOf(func() {}), true},
		{"Interface", reflect.ValueOf(assert.AnError), true},
		{"Map", reflect.ValueOf(map[int]any{}), true},
		{"Pointer", reflect.ValueOf(pointerOf("pointer")), true},
		{"Slice", reflect.ValueOf([]any{}), true},
		{"String", reflect.ValueOf("string"), false},
		{"Struct", reflect.ValueOf(struct{}{}), false},
		{"UnsafePointer", reflect.ValueOf(unsafe.Pointer(nil)), true},
	}

	for _, tt := range testcases {
		t.Run(tt.label, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, isnil.IsNilable(tt.value))
		})
	}
}

func nilPointerOf[T any]() *T {
	return nil
}

func pointerOf[T any](v T) *T {
	return &v
}

var _ isnil.IsNiler = (*nillable)(nil)

type nillable struct{}

func (n *nillable) IsNil() bool { return n == nil }
