func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp1 := make(map[rune]int)
	mp2 := make(map[rune]int)
	for _, v := range s {
		mp1[v]++
	}
	for _, v := range t {
		mp2[v]++
	}
	for _, v := range s {
		if mp1[v] != mp2[v] {
			return false
		}
	}
	return true
}
