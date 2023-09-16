package leetcode

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestSingleNumberIII1(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}
	target := []int{3, 5}

	result := singleNumberIII(nums)
	if !slices.Equal(result, target) && !slices.Equal(result, []int{5, 3}) {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}

func TestSingleNumberIII2(t *testing.T) {
	nums := []int{-1, 0}
	target := []int{-1, 0}

	result := singleNumberIII(nums)
	if !slices.Equal(result, target) && !slices.Equal(result, []int{0, -1}) {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}

func TestSingleNumberIII3(t *testing.T) {
	nums := []int{0, 1}
	target := []int{0, 1}

	result := singleNumberIII(nums)
	if !slices.Equal(result, target) && !slices.Equal(result, []int{1, 0}) {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}

func TestSingleNumberIII4(t *testing.T) {
	nums := []int{-283, -174, 44, -174, 40, 40, 101, 101, 354, -283, 513, 513}
	target := []int{44, 354}

	result := singleNumberIII(nums)
	if !slices.Equal(result, target) && !slices.Equal(result, []int{354, 44}) {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}
