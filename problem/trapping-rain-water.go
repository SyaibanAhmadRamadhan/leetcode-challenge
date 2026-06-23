package problem

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	left_max := height[left]
	right_max := height[right]
	trap := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] > left_max {
				left_max = height[left]
			}
			trap += left_max - height[left]
			left++
		} else {
			if height[right] > right_max {
				right_max = height[right]
			}
			trap += right_max - height[right]
			right--
		}
	}

	return trap
}
