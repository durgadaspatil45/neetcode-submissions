func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charM := make(map[byte]bool)
	l := 0
	res := 0
	for r:=0; r < len(s); r++ {
		for charM[s[r]] {
				delete(charM, s[l])
				l++
		}
		charM[s[r]] = true
		res = max(res, r - l + 1 )
	}
	return res
}
