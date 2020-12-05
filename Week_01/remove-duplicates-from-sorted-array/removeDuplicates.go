func removeDuplicates(nums []int) int {
	curIndex, index, l := 0, 1, len(nums)
	for ; index < l; index++ {
		if nums[index] != nums[curIndex] {
			curIndex++
			nums[curIndex] = nums[index]
		}
	}
	return curIndex + 1
}