package problem

// salah -> pakai i tanpa break
// 5,1,4,2,3
// 3,1,4,2,5
// 3,1,2,4,5

// salah 1 -> pakai i pake break
// 5,1,4,2,3
// 1,5,4,2,3
// 1,4,2,3,5

// benar -> minIdx selalu ke update
// 5,1,4,2,3
// 1,2,3,4,5
func HIndex(citations []int) int {
	for i := 0; i < len(citations); i++ {
		minIdx := i
		for j := i + 1; j < len(citations); j++ {
			if citations[j] >= citations[minIdx] {
				minIdx = j
			}
		}

		citations[i], citations[minIdx] = citations[minIdx], citations[i]
	}

	h := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] > h {
			h++
		}
	}

	return h
}
