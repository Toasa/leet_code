func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					swap(&nums[i], &nums[j])
					break
				}
			}
		}
	}
}

// [0,12,0,3,1]
// [12,0,0,3,1]
// [12,0,0,3,1]
// [12,3,1,0,0]