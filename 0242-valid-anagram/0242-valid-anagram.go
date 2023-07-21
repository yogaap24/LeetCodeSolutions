func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	var MapS = make(map[rune]int)
	var MapT = make(map[rune]int)
	for _, valueS := range s {
		MapS[valueS]++
	}
	for _, valueT := range t {
		MapT[valueT]++
	}
	for i, valMapS := range MapS {
		if MapT[i] != valMapS {
			return false
		}
	}
	return true
}