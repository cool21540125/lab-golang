package leetcode

import (
	"testing"

	"golang.org/x/exp/slices"
)

func ReverseSlice(slc []int) []int {
	rev_slc := []int{}
	for i := range slc {
		rev_slc = append(rev_slc, slc[len(slc)-1-i])
	}
	return rev_slc
}

func TestSingleNumberIII(t *testing.T) {
	testExamples := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 1, 3, 2, 5},
			expected: []int{5, 3},
		},
		{
			nums:     []int{-1, 0},
			expected: []int{-1, 0},
		},
		{
			nums:     []int{0, 1},
			expected: []int{0, 1},
		},
		{
			nums:     []int{-283, -174, 44, -174, 40, 40, 101, 101, 354, -283, 513, 513},
			expected: []int{354, 44},
		},
	}

	for _, example := range testExamples {
		result := singleNumberIII(example.nums)
		if !slices.Equal(result, example.expected) && !slices.Equal(result, ReverseSlice(example.expected)) {
			t.Errorf("Expected %v, but got %v", example.expected, result)
		}
	}
}
