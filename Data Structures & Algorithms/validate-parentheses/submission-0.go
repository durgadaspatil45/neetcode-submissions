func isValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		if open, exists := closeToOpen[c]; exists { //only checking at closing parentheses
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1] //poping top element which opening parentheses
				if top != open { // if this open and top(opening only) are not matched
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, c) //opening are adding to stack
		}
	}
	return len(stack) == 0
}