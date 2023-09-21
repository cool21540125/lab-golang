package leetcode

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testExamples := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}
	for _, example := range testExamples {
		result := singleNumber(example.nums)
		if result != example.expected {
			t.Errorf("Expected %v, but got %v", example.expected, result)
		}
	}
}
