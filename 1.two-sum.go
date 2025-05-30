package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		//complement to find the pair of num itself
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return nil
}

// @lc code=end
