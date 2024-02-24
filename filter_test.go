package totext

import (
	"testing"
)

// TestFilterNonReadableCharacter tests FilterNonReadableCharacter
func TestFilterNonReadableCharacter(t *testing.T) {
	// use testcases and loop through them
	testcases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"\r\n", "\n"},
		{"\n\n\r", "\n"},
		{"\n\n\n", "\n"},
		{"\n\n\n\n", "\n"},
		{"\n\n\n\n\n", "\n"},
		{"\n\n\n\n\n\n", "\n"},
		{"\n\n\n\n\n\n\n", "\n"},
		{"\n\n\n\n\n\n\n\n", "\n"},
		{"\n\n\n\n\n\n\n\n\n", "\n"},
	}

	for _, tc := range testcases {
		actual := FilterNonReadableCharacter(tc.input)
		if actual != tc.expected {
			t.Errorf("FilterNonReadableCharacter(%v): expected %v, got %v", tc.input, tc.expected, actual)
		}
	}
}
