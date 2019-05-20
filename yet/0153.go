func find_even(nums []int, l num) {
	head := 0
	tail := l - 1

	for nums[head] > nums[tail] {
		head++
		tail--
	}

	return nums[tail+1]
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	}

	l := len(nums)
	if l%2 == 0 {
		return find_even(nums, l)
	} else {
		return find_odd(nums, l)
	}
}