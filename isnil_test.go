// isnil_test.go - Tests for the isnil package

package isnil

import (
	"testing"
)

// Test cases for IsNil function
func TestIsNil(t *testing.T) {
	type testCase struct {
		input interface{}
		expected bool
	}

	// Grouped test cases for clarity
	var testCases = []testCase{
		{input: nil, expected: true},
		{input: "", expected: false},
		{input: 0, expected: false},
		{input: []int{}, expected: false},
		{input: map[string]int{}, expected: false},
	}

	// Execute each test case
	for _, tc := range testCases {
		result := IsNil(tc.input)
		if result != tc.expected {
			t.Fatalf("Expected IsNil(%v) to be %v, got %v", tc.input, tc.expected, result)
		}
	}
}