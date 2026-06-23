package problem

func canJump(nums []int) bool {
	reach := 0
	for i := range nums {
		if i > reach {
			return false
		}
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
		if reach >= len(nums)-1 {
			return true
		}
	}

	return false
}
