import "math"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	memo := make([]int, len(nums))
	memo[0] = 1
	max := 1

	for i := 1; i < len(nums); i++ {
		tmp_max := math.MinInt32
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if tmp_max < memo[j]+1 {
					tmp_max = memo[j] + 1
				}
			}
		}

		if tmp_max == math.MinInt32 {
			memo[i] = 1
		} else {
			memo[i] = tmp_max
		}

		if tmp_max > max {
			max = tmp_max
		}
	}
	//fmt.Println(memo)

	return max
}