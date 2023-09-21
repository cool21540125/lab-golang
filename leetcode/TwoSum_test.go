package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testExamples := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}
	for _, example := range testExamples {
		result := twoSum(example.nums, example.target)
		if result[0] != example.expected[0] || result[1] != example.expected[1] {
			t.Errorf("Expected %v but got %v", example.expected, result)
		}
	}
}
