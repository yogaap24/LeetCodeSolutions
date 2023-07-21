var m = map[byte]byte {
    '}': '{',
    ')': '(',
    ']': '[',
}

func isValid(s string) bool {
    str := []byte(s)

    stack := make([]byte, 0)

    for _, ch := range str {
        op, cl := m[ch]

        if !cl {
            stack = append(stack, ch)
            continue
        }

        if len(stack) == 0 {
            return false
        }

        lastOp := stack[len(stack) - 1]
        
        stack = stack[:len(stack) - 1] // "remove" it

        if lastOp != op {
            return false
        }
    }

    return len(stack) == 0
}