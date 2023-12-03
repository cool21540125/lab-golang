package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/maximum-average-subarray-i/description/?envType=study-plan-v2&envId=leetcode-75
func TestMaximumAverageSubarrayI(t *testing.T) {
	testExamples := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: 4,
		},
		{
			nums:     []int{-1, -2, -3, -4, -5},
			k:        3,
			expected: -2,
		},
	}

	for _, example := range testExamples {
		result := findMaxAverage(example.nums, example.k)
		if result != example.expected {
			t.Errorf("Expected %v, bot got %v", example.expected, result)
		}
	}
}
