package leetcode

func twoSum(sums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range sums {
		if idx, found := indexMap[target-currNum]; found {
			return []int{idx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{0, 0}
}
