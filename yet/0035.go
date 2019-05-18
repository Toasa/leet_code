var Nums []int
var Target int

// 中央にmidが来る
// 1 2 3 4 5
// (0 + 4) / 2 = 2

// 左に１つずれ、midが来る
// 1 2 3 4
// (0 + 3) / 2 = 1

// [1, 3, 5, 6]
// target = 6
// min = 0, max = 3, mid = 1
// min = 1, max = 3, mid = 2
// min = 2, max = 3, mid = 2

func biSearch(min, max int) int {
	if max-min == 1 && (Nums[min] < Target && Target < Nums[max]) {
		return max
	}

	mid := (min + max) / 2

	if Nums[mid] == Target {
		return mid
	} else if Nums[min] == Target {
		return min
	} else if Nums[max] == Target {
		return max
	}

	if Nums[mid] < Target {
		return biSearch(mid, max)
	} else {
		return biSearch(min, mid)
	}
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	Nums = nums
	Target = target
	return biSearch(0, len(nums)-1)
}