func plusOne(digits []int) []int {
    for i := range digits {
		if digits[len(digits)-1-i] == 9 {
			digits[len(digits)-1-i] = 0
		} else {
			digits[len(digits)-1-i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}