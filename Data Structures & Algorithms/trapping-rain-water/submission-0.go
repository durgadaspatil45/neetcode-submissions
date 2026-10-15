func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	leftMax := height[l]
	rightMax := height[r]
	res := 0
	for l < r {
		if leftMax < rightMax {
			l += 1
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r -= 1
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}
	return res
}
