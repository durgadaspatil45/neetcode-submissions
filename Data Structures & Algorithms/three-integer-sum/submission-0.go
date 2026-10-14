func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums) //sort the input slice
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			tsum := nums[i] + nums[start] + nums[end]
			if tsum > 0{
				end--
			} else if tsum < 0 {
				start++
			} else {
				result = append(result, []int{nums[i], nums[start], nums[end]})
				start++
				for nums[start] == nums[start-1] && start < end{ // if new start is same as previous one jump to next number do not want repeate start
					start++
				}
			}
			
		}
	}
	return result
}