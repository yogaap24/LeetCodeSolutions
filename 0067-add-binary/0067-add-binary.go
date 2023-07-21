func addBinary(a string, b string) string {
    if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}
	
	if a[len(a)-1] == '1' && b[len(b)-1] == '1' {
		return addBinary(addBinary(a[:len(a)-1], b[:len(b)-1]), "1") + "0"
	} else if a[len(a)-1] == '0' && b[len(b)-1] == '0' {
		return addBinary(a[:len(a)-1], b[:len(b)-1]) + "0"
	} else {
		return addBinary(a[:len(a)-1], b[:len(b)-1]) + "1"
	}
}