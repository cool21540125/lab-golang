package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	if result[0] != 0 || result[1] != 1 {
		t.Errorf("Expected [0,1] but got %v", result)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)
	if result[0] != 1 || result[1] != 2 {
		t.Errorf("Expected [1,2] but got %v", result)
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	result := twoSum(nums, target)
	if result[0] != 0 || result[1] != 1 {
		t.Errorf("Expected [0,1] but got %v", result)
	}
}

func TestTwoSum4(t *testing.T) {
	nums := []int{3, 3}
	target := 8

	result := twoSum(nums, target)
	if result[0] != 0 || result[1] != 0 {
		t.Errorf("Expected [0,0] but got %v", result)
	}
}

func TestTwoSum5(t *testing.T) {
	nums := []int{3, 6, 9, 12, 15}
	target := 27

	result := twoSum(nums, target)
	if result[0] != 3 && result[1] != 4 {
		t.Errorf("Expected [3,4] but got %v", result)
	}
}
