package problem

func RemoveDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	write := 2
	for i := 2; i < len(nums); i++ {
		if nums[write-2] != nums[i] {
			nums[write] = nums[i]
			write++
		}
	}

	return write
}
