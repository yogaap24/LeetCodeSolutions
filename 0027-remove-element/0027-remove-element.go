func removeElement(nums []int, val int) int {
    i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i = i + 1
		}
	}
	return i
}