func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0

	for l < r {
		res = max(res, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}