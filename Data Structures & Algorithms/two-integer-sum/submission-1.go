func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	rem := make(map[int]int)
	for i, v := range nums {
		r := target - v
		if val, ok := rem[r]; ok {
			return []int{val, i}
		}
		rem[v] = i
	}
	return nil
}
