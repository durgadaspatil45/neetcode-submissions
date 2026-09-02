func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	longest := 0
	for i, v := range nums {
		mp[v] = i
	}
	for num := range mp {
		if _, ok := mp[num-1]; !ok {
			length := 1
			for {
				if _, exists := mp[num+length]; exists {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}