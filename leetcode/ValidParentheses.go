package leetcode

import "strings"

func isValid(s string) bool {
	left := ""

	for _, char := range s {
		cc := string(char)
		if strings.Contains("[{(", cc) {
			left += cc
		} else if left == "" {
			return false
		} else if (cc == "]" && string(left[len(left)-1:]) == "[") ||
			(cc == "}" && string(left[len(left)-1:]) == "{") ||
			(cc == ")" && string(left[len(left)-1:]) == "(") {
			left = string(left[0 : len(left)-1])
		}
	}
	return left == ""
}
