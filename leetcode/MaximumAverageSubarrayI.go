package leetcode

func findMaxAverage(nums []int, k int) float64 {
	var summ int
	for _, item := range nums[:k] {
		summ = summ + item
	}

	largestSumm := summ

	for idx := range nums[k:] {
		summ = summ - nums[idx] + nums[idx+k]
		if summ > largestSumm {
			largestSumm = summ
		}
	}
	return float64(largestSumm) / float64(k)
}
