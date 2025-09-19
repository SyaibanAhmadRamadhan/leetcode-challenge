package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

// This function uses a hash map (hash table) as its data structure
// and applies the One-pass Hash Table strategy,
// which is a type of Linear Time Algorithm (O(n)).
//
// What is One-pass Hash Table (Linear Time Algorithm)?
// → It's an algorithmic strategy that:
//  1. Uses a hash table (hash map) to store previously seen values.
//  2. Iterates through the data only once (one pass) to:
//     - Find the solution (check if complement exists).
//     - Store each number for future lookups in the same iteration.
//
// Why is it called "one-pass"?
// → Because the entire array is scanned only once,
//
//	and both the searching and storing processes happen simultaneously.
//
// Relationship between the two:
// → One-pass Hash Table is a strategy that results in Linear Time Complexity (O(n)),
//
//	because we loop through the data only once, making the processing time
//	grow proportionally to the input size.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		//complement to find the pair of num itself
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{i, j}
		}

		// Store the current number with its index
		seen[num] = i
	}

	return nil
}

// @lc code=end
