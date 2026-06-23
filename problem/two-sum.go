package problem

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, 0)

	for i, v := range nums {
		complement := target - v
		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}

		seen[v] = i
	}

	return nil

}
