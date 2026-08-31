type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	encStr := ""
	for _, val := range strs {
		n := len(string(val))
		num := strconv.Itoa(n)
		encStr += num + "#" + val
	}
	return encStr
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for string(encoded[j]) != "#" {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+length])
		i += length
	}
	return res
}
