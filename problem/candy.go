package problem

import "fmt"

// [5,4,3,2,1]
// [1,1,1,1,1] -> kiri - kanan
// [5,4,3,2,1] -> kanan - kiri
//
// [1,2,2]
// [1,2,1] -> kiri - kanan
// [1,2,1] -> kanan - kiri
//
// [1,0,2]
// [1,1,2] -> kiri - kanan
// [2,1,2] -> kanan - kiri
//
// [1,3,4,5,2]
// [1,2,3,4,1] -> kiri - kanan
// [1,2,3,4,1] -> kanan - kiri
//
// [1,2,87,87,87,2,1]
// [1,2,3,1,1,1,1] -> kiri - kanan
// [1,2,3,1,3,2,1] -> kanan - kiri
//
// [1,6,10,8,7,3,2]
// [1,2,3,1,1,1,1] -> kiri - kanan
// [1,2,5,4,3,2,1] -> kanan - kiri
func Candy(ratings []int) int {
	left := make([]int, len(ratings))

	// left to right
	for i := range ratings {
		if i == 0 {
			left[0] = 1
		} else {
			if ratings[i] > ratings[i-1] {
				left[i] = left[i-1] + 1
			} else {
				left[i] = 1
			}
		}
	}

	total := 0
	// right to left and count total
	for i := len(ratings) - 1; i >= 0; i-- {
		if i == len(ratings)-1 {
			total += left[i]
		} else {
			if ratings[i] > ratings[i+1] && left[i] <= left[i+1] {
				left[i] = left[i+1] + 1
				total += left[i]
			} else {
				total += left[i]
			}
		}
	}

	fmt.Println(total)
	fmt.Println(left)
	return total
}
