func lengthOfLastWord(s string) int {
    count := 0
	for i := range s {
		if s[i] != ' ' {
			count++
		} else if i+1 < len(s) && s[i+1] != ' ' {
			count = 0
		}
	}
	return count
}