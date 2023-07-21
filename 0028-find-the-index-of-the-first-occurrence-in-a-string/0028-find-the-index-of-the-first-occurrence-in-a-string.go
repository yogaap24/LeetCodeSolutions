func strStr(haystack string, needle string) int {
    if needle == "" {
		return 0
	}
	for i := range haystack {
		if haystack[i] == needle[0] {
			if i+len(needle) > len(haystack) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}