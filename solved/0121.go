func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	min := prices[0]
	max := prices[0]

	cur_max := -1

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			max = prices[i]
		}
		if prices[i] > max {
			max = prices[i]
		}

		if max-min > cur_max {
			cur_max = max - min
		}
	}

	return cur_max
}