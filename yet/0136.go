func singleNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; {
		if nums[i] != nums[i + 1] {
			return nums[i]
		}
		old := nums[i]
		for nums[i] == old {
			i++
		}
	}
	return nums[len(nums) - 1]
}