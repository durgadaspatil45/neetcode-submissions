func maxArea(heights []int) int {
	start := 0
	end := len(heights) - 1
	result := 0
	for start < end {
		// min here as whoever minimum that is our bucket height can't consider higher one it will spill our water
		area := (end - start) * min(heights[start], heights[end]) // area = width * height
		result = max(result, area)
		if heights[start] < heights[end] {
			start++
		} else {
			end--
		}
	}
	return result
}
