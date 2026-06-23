package problem

// [2,3,1,1,4]
// [2,3,0,1,4]
// [0]
// [1,2,1,1,1]
// [1,2,0,1]
// [2,1,1,1,1]
// [4,1,1,3,1,1,1]
func JumpII(nums []int) int {
	currentEnd := 0
	jump := 0
	farthest := 0
	for i := range nums {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jump++
			currentEnd = farthest
			if currentEnd >= len(nums)-1 {
				return jump
			}
		}
	}

	return jump
}
