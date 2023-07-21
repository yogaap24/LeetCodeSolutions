func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	var i int
	for j := range nums {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}