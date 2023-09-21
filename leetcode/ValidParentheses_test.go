package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testExamples := []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "(]",
			expected: false,
		},
		{
			s:        "{[]}",
			expected: true,
		},
		{
			s:        "((",
			expected: false,
		},
		{
			s:        "[",
			expected: false,
		},
		{
			s:        "][",
			expected: false,
		},
		{
			s:        "(){}}{",
			expected: false,
		},
		{
			s:        "{}[{}]((){})(){}",
			expected: true,
		},
	}
	for _, example := range testExamples {
		result := isValid(example.s)
		if result != example.expected {
			t.Errorf("Expected true but got %v", result)
		}
	}
}
