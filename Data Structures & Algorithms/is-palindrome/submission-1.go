func isPalindrome(s string) bool {
	var cleaned strings.Builder
	for _, char := range strings.ToLower(s) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(char)
		}
	}
	cleanedStr := cleaned.String()
	start := 0
	end := len(cleanedStr) - 1
	for start < end {

		if cleanedStr[start] != cleanedStr[end] {
			return false
		}
		start++
		end--
	}
	return true
}
