func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return Max(nums[0], nums[1])
	}

	// 末尾を省く
	memo1 := make([]int, len(nums)-1)
	// 先頭を省く
	memo2 := make([]int, len(nums)-1)

	memo1[0] = nums[0]
	memo1[1] = Max(nums[0], nums[1])

	memo2[0] = nums[1]
	memo2[1] = Max(nums[1], nums[2])

	for i := 2; i < len(nums)-1; i++ {
		memo1[i] = Max(memo1[i-1], memo1[i-2]+nums[i])
		memo2[i] = Max(memo2[i-1], memo2[i-2]+nums[i+1])
	}
	return Max(memo1[len(memo1)-1], memo2[len(memo2)-1])
}