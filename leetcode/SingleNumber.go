package leetcode

import "golang.org/x/exp/maps"

func singleNumber(nums []int) []int {

	set := make(map[int]bool)

	for _, num := range nums {
		if _, found := set[num]; found {
			delete(set, num)
		} else {
			set[num] = true
		}
	}
	return maps.Keys(set)
}
