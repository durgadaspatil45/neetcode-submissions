func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string) ////Go allows arrays to be map keys if their elements are comparable. slices are not.
	for _, word := range strs{
		var count [26]int//Go allows arrays to be map keys if their elements are comparable. slices are not.
		for _, char := range word{
			count[char - 'a']++ //e-a that is 101 - 97 = 4 so at count array at index 4 the value become count[4]=+1
		}
		res[count] = append(res[count], word)
	}
	var result [][]string
	for _, group := range res{
		result=append(result, group)
	}
	return result
}
