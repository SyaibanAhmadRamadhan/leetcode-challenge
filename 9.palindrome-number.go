package leetcode

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	// cara 1
	// result
	// 11511/11511 cases passed (8 ms)
	// Your runtime beats 18.48 % of golang submissions
	// Your memory usage beats 56.09 % of golang submissions (6.2 MB)
	// {
	// 	str := strconv.Itoa(x)
	// 	b := []byte(str)
	// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
	// 		if b[i] != b[j] {
	// 			return false
	// 		}
	// 	}
	// 	return true
	// }

	// digit manipulation number
	// cara 2
	// 11511/11511 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 79 % of golang submissions (6 MB)
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
