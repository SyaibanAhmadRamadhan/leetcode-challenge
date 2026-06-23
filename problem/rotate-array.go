package problem

func rotate(nums []int, k int) {
	k %= len(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[j], nums[i] = nums[i], nums[j]
	}

	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[j], nums[i] = nums[i], nums[j]
	}

	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[j], nums[i] = nums[i], nums[j]
	}
}
