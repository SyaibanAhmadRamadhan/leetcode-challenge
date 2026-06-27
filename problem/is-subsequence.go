package problem

func isSubsequence(s string, t string) bool {
	tt := map[rune][]int{}
	for i, v := range t {
		_, ok := tt[v]
		if !ok {
			tt[v] = []int{i}
		} else {
			tt[v] = append(tt[v], i)
		}
	}

	latestIndex := -1
	for _, v := range s {
		indexex, ok := tt[v]
		if !ok {
			return false
		}

		index, exists := findHighIndexFromLatestIndex(indexex, latestIndex)
		if !exists {
			return false
		}

		latestIndex = index
	}

	return true
}

func findHighIndexFromLatestIndex(indexex []int, latestIndex int) (int, bool) {
	left := 0
	right := len(indexex) - 1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		if indexex[mid] > latestIndex {
			index = indexex[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if index == -1 {
		return index, false
	}

	return index, true
}
