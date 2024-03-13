package leetcode

func pivotInteger(n int) int {
	if n < 1 || n > 1000 {
		return -1
	}

	total := (1 + n) * n / 2
	pivot_sum := 0

	for ii := n; ii >= 1; ii-- {
		pivot_sum += ii
		if pivot_sum == total {
			return ii
		} else if pivot_sum > total {
			return -1
		}
		total -= ii
	}
	return -1
}
