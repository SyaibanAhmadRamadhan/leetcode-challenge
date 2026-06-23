package problem

func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := (left + right) / 2
		valMid := nums[mid]
		if valMid == target {
			return mid
		} else if valMid > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
