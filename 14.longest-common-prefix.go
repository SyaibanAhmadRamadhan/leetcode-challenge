package leetcode

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	prefix := strs[0]

	for _, v := range strs[1:] {
		for len(prefix) > len(v) || (v[:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

// @lc code=end
