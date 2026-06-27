package problem

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left <= right {
		width := right - left
		minHeight := height[left]
		if height[right] < minHeight {
			minHeight = height[right]
			right--
		} else {
			left++
		}

		if width*minHeight > max {
			max = width * minHeight
		}
	}

	return max
}
