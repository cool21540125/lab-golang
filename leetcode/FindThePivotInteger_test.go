package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/find-the-pivot-integer/?envType=daily-question&envId=2024-03-13
func TestPivotInteger(t *testing.T) {
	testExamples := []struct {
		nn    int
		pivot int
	}{
		{
			nn:    8,
			pivot: 6,
		},
		{
			nn:    1,
			pivot: 1,
		},
		{
			nn:    4,
			pivot: -1,
		},
		{
			nn:    3,
			pivot: -1,
		},
		{
			nn:    49,
			pivot: 35,
		},
	}

	for _, example := range testExamples {
		result := pivotInteger(example.nn)
		if result != example.pivot {
			t.Errorf("Expected %v, bot got %v", example.pivot, result)
		}
	}
}
