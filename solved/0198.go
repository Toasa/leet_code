func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	memo := make([]int, len(nums))
	memo[0] = nums[0]
	memo[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		memo[i] = Max(memo[i-1], nums[i]+memo[i-2])
	}

	return memo[len(memo)-1]
}