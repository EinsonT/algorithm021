func merge(nums1 []int, m int, nums2 []int, n int) {
	total := m + n - 1
	index1, index2 := m-1, n-1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] < nums2[index2] {
			nums1[total] = nums2[index2]
			index2--
		} else {
			nums1[total] = nums1[index1]
			index1--
		}
		total--
	}
	for index1 >= 0 {
		nums1[total] = nums1[index1]
		total--
		index1--
	}
	for index2 >= 0 {
		nums1[total] = nums2[index2]
		total--
		index2--
	}
}