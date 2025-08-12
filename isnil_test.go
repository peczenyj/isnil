package isnil_test

import (
	"fmt"
	"io/fs"
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

func nilPointerOf[T any]() *T {
	return nil
}

func pointerOf[T any](v T) *T {
	return &v
}

type nillable struct{}

func (n *nillable) IsNil() bool { return n == nil }
