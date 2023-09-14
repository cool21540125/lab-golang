package leetcode

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	/*
		( 40
		) 41
		[ 91
		] 93
		{ 123
		} 125
	*/
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	left := []rune{}

	for _, rr := range s {
		if _, found := pairs[rr]; found {
			left = append(left, rr)
		} else if len(left) == 0 || pairs[left[len(left)-1]] != rr {
			return false
		} else {
			left = left[:len(left)-1]
		}
	}
	return len(left) == 0
}
