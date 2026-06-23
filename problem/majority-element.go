package problem

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, v := range nums {
		if count < 0 {
			candidate = v
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}

	return candidate
}
