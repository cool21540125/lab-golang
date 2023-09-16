package leetcode

func singleNumber(nums []int) int {
	for _, num := range nums {
		repeated := count(nums, num)
		if repeated == 1 {
			return num
		}
	}
	return 0
}

func count(nums []int, target int) int {
	cnt := 0
	for _, num := range nums {
		if num == target {
			cnt += 1
		}
	}
	return cnt
}
