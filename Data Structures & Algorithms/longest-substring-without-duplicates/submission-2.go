func lengthOfLongestSubstring(s string) int {
	charM := make(map[byte]bool)
	l := 0
	res := 0
	for r:=0; r < len(s); r++ {
		for charM[s[r]] {
				delete(charM, s[l]) // delete the repeated char
				l++
		}
		charM[s[r]] = true
		res = max(res, r - l + 1 )
	}
	return res
}
