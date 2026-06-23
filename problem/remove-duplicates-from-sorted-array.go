package problem

func RemoveDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
