func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	for start < end {
		if numbers[start]+numbers[end] == target {
			break
		}
		if numbers[start]+numbers[end] > target {
			end--
		} else {
			start++
		}
	}
	return []int{start+1, end+1}
}
