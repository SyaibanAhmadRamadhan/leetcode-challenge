package problem

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		nums2val := nums2[n-1]
		if m-1 < 0 {
			nums1[m+n-1] = nums2val
			n--
			continue
		}
		nums1val := nums1[m-1]
		if nums2val >= nums1val {
			nums1[m+n-1] = nums2val
			n--
		} else {
			nums1[m+n-1] = nums1val
			m--
		}
	}
}
