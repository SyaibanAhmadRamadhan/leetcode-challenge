package problem

import (
	"slices"
)

// [-100,-70,-60,110,120,130,160]
// [-100,-70,-60,110,120,130,160]

// [0,0,0,0]

// [0,0,0]

// [0,1,1]

// [-1,0,1,2,-1,-4]
// [-4,-1,-1,0,1,2]

// [1,2,-2,-1]

// [2,-3,0,-2,-5,-5,-4,1,2,-2,2,0,2,-4,5,5,-10]
// [-10, -5, -5, -4, -4, -3, -2, -2, 0, 0, 1, 2, 2, 2, 2, 5, 5]
// [-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]

// [-100,-70,-60,110,120,130,160]
// [-60,-70,-100,110,120,130,160]

// [0,0,0,0]

// [0,0,0]

// [0,1,1]

// [-1,0,1,2,-1,-4]
// [-4,-1,-1,0,1,2]

// [1,2,-2,-1]

// [2,-3,0,-2,-5,-5,-4,1,2,-2,2,0,2,-4,5,5,-10]
// [-10, -5, -5, -4, -4, -3, -2, -2, 0, 0, 1, 2, 2, 2, 2, 5, 5]
// [-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	output := [][]int{}
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				candidates := []int{nums[i], nums[left], nums[right]}
				slices.Sort(candidates)
				output = appendUnique(output, candidates)
				right--
				left++
				continue
			} else {
				if nums[left]+nums[right] > nums[i]*-1 {
					right--
				} else {
					left++
				}
			}
		}
	}
	return output
}

func appendUnique(output [][]int, candidates []int) [][]int {
	if len(output) == 0 {
		output = append(output, candidates)
	} else {
		if output[len(output)-1][0] != candidates[0] || output[len(output)-1][1] != candidates[1] || output[len(output)-1][2] != candidates[2] {
			output = append(output, candidates)
		}
	}

	return output
}
