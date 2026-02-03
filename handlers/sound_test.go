package handlers

import (
	"fmt"
	"testing"
)

// Test numeric values for ID
func TestStringIsInt(t *testing.T) {

	// Test data
	testCases := []struct {
		name string
		id   string
		want int
	}{
		{"Normal int", "123", 1},
		{"Single int", "1", 1},
		{"Large int", "123456789", 1},
		{"Larger int (Long)", "123456789123456789", -1},
		{"text", "hello", -1},
		{"word and int mixed", "12ab", -1},
		{"Float", "12.34", -1},
	}

	// The test itself
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Testing: %s", tc.id), func(t *testing.T) {
			if stringIsInt(tc.id) != tc.want {
				t.Errorf(`StringIsInt("%s") = %v, want "1", error`, tc.id, tc.want)
			}
		})
	}

}
