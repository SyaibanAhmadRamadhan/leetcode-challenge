package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// complement := target-x
	seen := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}

		seen[nums[i]] = i
	}

	return nil
}

// @lc code=end
