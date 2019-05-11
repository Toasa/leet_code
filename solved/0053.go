func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	memo := make([]int, len(nums))
	memo[0] = nums[0]
	max := nums[0]
	for k := 1; k < len(nums); k++ {
		if memo[k-1]+nums[k] >= nums[k] {
			memo[k] = memo[k-1] + nums[k]
		} else {
			memo[k] = nums[k]
		}
		if memo[k] > max {
			max = memo[k]
		}
	}
	return max
}