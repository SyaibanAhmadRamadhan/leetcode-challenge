package problem

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1
	for left < right {
		l := s[left]
		r := s[right]
		canCompare := true
		if (l < 'a' || l > 'z') && (l < '0' || l > '9') {
			left++
			canCompare = false
		}
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') {
			right--
			canCompare = false
		}
		if !canCompare {
			continue
		}
		if l != r {
			return false
		}
		left++
		right--
	}

	return true
}
