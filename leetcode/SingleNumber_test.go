package leetcode

import (
	"testing"
)

func TestSingleNumber1(t *testing.T) {
	nums := []int{2, 2, 1}
	target := 1

	result := singleNumber(nums)
	if result != target {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}

func TestSingleNumber2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	target := 4

	result := singleNumber(nums)
	if result != target {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}

func TestSingleNumber3(t *testing.T) {
	nums := []int{1}
	target := 1

	result := singleNumber(nums)
	if result != target {
		t.Errorf("Expected %v, but got %v", target, result)
	}
}
