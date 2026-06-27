package problem

func TwoSumIIArraySorted(numbers []int, target int) []int {
	// left := 0
	// right := len(numbers) - 1
	// for left <= right {
	// 	if numbers[left]+numbers[right] < target {
	// 		left++
	// 	} else if numbers[left]+numbers[right] > target {
	// 		right--
	// 	} else {
	// 		return []int{left + 1, right + 1}
	// 	}
	// }

	// return make([]int, 2)

	for i := 0; i < len(numbers); i++ {
		complement := target - numbers[i]
		indexRight := bsRecursive(numbers, i+1, len(numbers)-1, complement)
		if indexRight != -1 {
			return []int{i + 1, indexRight + 1}
		}
	}

	return []int{0, 0}

}

func bsRecursive(numbers []int, start, end, target int) int {
	if start <= end {
		mid := start + (end-start)/2
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] < target {
			return bsRecursive(numbers, mid+1, end, target)
		} else {
			return bsRecursive(numbers, start, mid-1, target)
		}
	}

	return -1
}
