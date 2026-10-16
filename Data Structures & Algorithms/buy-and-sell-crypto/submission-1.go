func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxPro := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxPro = max(maxPro, profit)
		} else {
			l = r
		}
		r += 1
	}
	return maxPro
}
