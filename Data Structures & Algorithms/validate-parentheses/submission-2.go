func isValid(s string) bool {
    parenthesis := map[byte]byte{')': '(', '}': '{', ']': '['}
	var stack []byte

	for i:= 0; i < len(s); i++ {
		char := s[i]
		if expected, ok := parenthesis[char]; ok {
			if len(stack) == 0 { return false }
			last := stack[len(stack)-1]
			if last != expected {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}