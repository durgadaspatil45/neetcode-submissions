func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int) //occurence --> values which has occurences
	freq := make([][]int, len(nums)+1)
	// freq is 1 index is occurences under 1 occurencse it has list element which has 1 occurences and so on
	for _, v := range nums {
		mp[v]++
	}
	for k, v := range mp {
		freq[v] = append(freq[v], k)
	}

	res := []int{}
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
