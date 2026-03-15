package leetcode

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	// {
	// 	str := strconv.Itoa(x)
	// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
	// 		if str[i] != str[j] {
	// 			return false
	// 		}
	// 	}

	// 	return true
	// }

	// digital manipulation number
	// 121/10 12
	// 121%10 = 1
	// last := x%10 -> 1
	// 121/100 = 1
	// 21890/10_000 = 2
	{
		if x < 0 {
			return false
		}

		divisor := 1
		for x/divisor >= 10 {
			divisor *= 10
		}

		for x > 0 {
			first := x / divisor
			last := x % 10

			if first != last {
				return false
			}

			x = (x % divisor) / 10
			divisor /= 100
		}

		return true
	}
}

// @lc code=end
